package db

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/pkg/alog"

	// Import this to avoid import cycle error
	// See https://entgo.io/docs/hooks#hooks-registration
	_ "github.com/Guaderxx/gowebtmpl/ent/runtime"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open new connection
func Open(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("open db failed: %s", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	AutoMigrate(client)
	return client, nil
}

func AutoMigrate(db *ent.Client) {
	if err := db.Schema.Create(context.Background()); err != nil {
		alog.Fatal("create schema resources failed", "error", err.Error())
	}
}
