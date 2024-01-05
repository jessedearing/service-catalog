package main

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-migrate/migrate/v4"
	migratepgx5 "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jessedearing/service-catalog/graph"
	"github.com/jessedearing/service-catalog/internal/storage"
	"github.com/jessedearing/service-catalog/internal/vars"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()
	pgurl := os.Getenv("PGURL")
	dbconfig, err := pgxpool.ParseConfig(pgurl)
	if err != nil {
		panic(err)
	}

	dbconfig.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		pgxUUID.Register(c.TypeMap())
		return nil
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, dbconfig)
	if err != nil {
		panic(err)
	}

	if err = dbpool.Ping(ctx); err != nil {
		panic(err)
	}

	conn := stdlib.OpenDBFromPool(dbpool)
	drv, err := migratepgx5.WithInstance(conn, &migratepgx5.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "service_catalog", drv)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		slog.ErrorContext(ctx, err.Error())
	} else {
		slog.InfoContext(ctx, "No new db migrations to run")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)
	db := storage.New(dbpool)

	httpserver := http.Server{
		BaseContext: func(l net.Listener) context.Context {
			dbctx := context.WithValue(ctx, vars.DBContextKey, db)
			return dbctx
		},
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(httpserver.ListenAndServe())
}
