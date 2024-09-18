package db

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/Guaderxx/gowebtmpl/ent"

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
	return ent.NewClient(ent.Driver(drv)), nil
}
