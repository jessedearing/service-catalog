package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jessedearing/service-catalog/graph/model"
)

// TODO: PageSize is hard coded now, but should be made variable later <2024-01-05, Jesse Dearing>
// PageSize is the number of services per page.
const PageSize = 5

type ServiceStorage interface {
	All(ctx context.Context, page int) ([]model.Service, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Service, error)
	FindByName(ctx context.Context, name string) ([]*model.Service, error)
	SearchAll(ctx context.Context, query string) ([]*model.Service, error)
}

type DB struct {
	DB *pgxpool.Pool
	ServiceStorage
}

// New creates a new PostgresDB object
func New(connpool *pgxpool.Pool) *DB {
	return &DB{
		DB: connpool,
	}
}

// All will return the number of serices for the page number specified by the
// `page` parameter
func (d *DB) All(ctx context.Context, page int) ([]*model.Service, error) {
	rows, err := d.DB.Query(ctx, allPagedQuery, PageSize, (page-1)*PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return readServicesFromRows(rows)
}

func (d *DB) FindByID(ctx context.Context, id uuid.UUID) (*model.Service, error) {
	rows, err := d.DB.Query(ctx, singleServiceQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	svcs, err := readServicesFromRows(rows)
	if err != nil {
		return nil, err
	}
	return svcs[0], nil
}

func (d *DB) FindByName(ctx context.Context, name string) ([]*model.Service, error) {
	rows, err := d.DB.Query(ctx, searchByNameQuery, name, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	svcs, err := readServicesFromRows(rows)
	if err != nil {
		return nil, err
	}

	return svcs, nil
}

func (d *DB) SearchAll(ctx context.Context, query string) ([]*model.Service, error) {
	rows, err := d.DB.Query(ctx, searchAllQuery, query, query, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	svcs, err := readServicesFromRows(rows)
	if err != nil {
		return nil, err
	}

	return svcs, nil
}

// readServicesFromRows will read the services and the versions and return the
// graphql model populated from the db
func readServicesFromRows(rows pgx.Rows) ([]*model.Service, error) {
	svcs := map[string]*model.Service{}
	for rows.Next() {
		var id, vid uuid.UUID
		var name, description, version string
		err := rows.Scan(&id, &name, &description, &vid, &version)
		if err != nil {
			return []*model.Service{}, err
		}
		svc := svcs[id.String()]
		if svc != nil {
			svc.Versions = append(svc.Versions, &model.Version{ID: vid, Version: version})
		} else {
			svc = &model.Service{
				ID:          id,
				Name:        name,
				Description: description,
				Versions:    []*model.Version{&model.Version{ID: vid, Version: version}},
			}
		}

		svcs[id.String()] = svc
	}

	retsvcs := []*model.Service{}

	for _, svc := range svcs {
		retsvcs = append(retsvcs, svc)
	}

	return retsvcs, nil
}
