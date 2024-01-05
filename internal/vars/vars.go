package vars

import (
	"context"
	"errors"

	"github.com/jessedearing/service-catalog/internal/storage"
)

var DBContextKey = "service-catalog-db"
var ErrNoDBInContext = errors.New("DB object not in context")

func GetDBFromContext(ctx context.Context) (*storage.DB, error) {
	db, ok := ctx.Value(DBContextKey).(*storage.DB)
	if !ok {
		return nil, ErrNoDBInContext
	}

	return db, nil
}
