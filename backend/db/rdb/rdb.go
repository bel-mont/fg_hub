package rdb

import (
	"context"
	"database/sql"
	"fg_hub/backend/db/rdb/migrations"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

// Used to run the DB migrations.
var pgxMigrateConnStr = "pgx://postgres:root@localhost:5432/fghub?sslmode=disable"

var connStr = "postgres://postgres:root@localhost:5432/fghub?sslmode=disable"

// GetConn Fetches a connection
// TODO: get rid of this, implement connections together with SQL Boiler
func GetConn() *bun.DB {
	pgDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))
	db := bun.NewDB(pgDb, pgdialect.New())
	return db
}

func Migrate() {
	db := GetConn()
	migrator := migrate.NewMigrator(db, migrations.Migrations)
	ctx := context.Background()
	err := migrator.Init(ctx)
	if err != nil {
		fmt.Println("Error on migration init: ", err)
	}
	migrationGroup, err := migrator.Migrate(ctx)
	if err != nil {
		fmt.Println("Error on migrate: ", err)
	}
	fmt.Println("Applied migration: ", migrationGroup.ID)
}
