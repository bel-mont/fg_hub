package rdb

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// TODO: Move to a config file / env variables
var connStr = "postgres://postgres:root@fghub-db/fghub?sslmode=disable"

// NewConn Fetches a connection
// TODO: Check connection close
func NewConn() *bun.DB {
	pgDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))
	db := bun.NewDB(pgDb, pgdialect.New())
	return db
}
