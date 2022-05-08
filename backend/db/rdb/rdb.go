package rdb

import (
	"database/sql"
	"fg_hub/backend/db/rdb/migrations"
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
	migrate.NewMigrator(db, migrations.Migrations)
	//
	////filePath := "file://fg_hub/backend/db/rdb/migrations"
	//// for now, relative because it is breaking when called from another module
	//filePath := "file://../../db/rdb/migrations"
	//m, err := migrate.New(
	//	filePath,
	//	pgxMigrateConnStr)
	//if err != nil {
	//	log.Fatal("Migration error", err)
	//}
	//if err := m.Up(); err != nil {
	//	log.Println("Up error", err)
	//}
}
