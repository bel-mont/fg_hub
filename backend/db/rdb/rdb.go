package rdb

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Conn *pgx.Conn

//var databaseUrl = "host=localhost user=postgres password=root dbname=fghub port=5432 sslmode=disable TimeZone=Asia/Tokyo"

// Used for DB connections
var pgxConnStr = "postgres://postgres:root@localhost:5432/fghub?sslmode=disable"

// Used to run the DB migrations.
var pgxMigrateConnStr = "pgx://postgres:root@localhost:5432/fghub?sslmode=disable"

func InitDB() {
	conn, err := pgx.Connect(context.Background(), pgxConnStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	Conn = conn
}

// GetConn Fetches a connection
// TODO: get rid of this, implement connections together with SQL Boiler
func GetConn() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), pgxConnStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}

func Migrate() {
	//filePath := "file://fg_hub/backend/db/rdb/migrations"
	// for now, relative because it is breaking when called from another module
	filePath := "file://../../db/rdb/migrations"
	m, err := migrate.New(
		filePath,
		pgxMigrateConnStr)
	if err != nil {
		log.Fatal("Migration error", err)
	}
	if err := m.Up(); err != nil {
		log.Println("Up error", err)
	}
}
