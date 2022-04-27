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

func Test() {
	// m, _ := migrate.New(
	// 	"github://mattes:personal-access-token@mattes/migrate_test",
	// 	"postgres://postgres:root@localhost:5432/fghub-db?sslmode=disable")
	// // host=localhost user=postgres password=root dbname=fghub-db port=5432 sslmode=disable TimeZone=Asia/Tokyo
	// m.Steps(2)
}

var Conn *pgx.Conn

//var databaseUrl = "host=localhost user=postgres password=root dbname=fghub port=5432 sslmode=disable TimeZone=Asia/Tokyo"
var pgxConnStr = "postgres://postgres:root@localhost:5432/fghub?sslmode=disable"
var pgxMigrateStr = "pgx://postgres:root@localhost:5432/fghub?sslmode=disable"

func InitDB() {
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	// db, err := sql.Open("postgres", "postgres:root@localhost/fghub-db")
	// if err != nil {
	// 	log.Panic(err)
	// }

	// if err = db.Ping(); err != nil {
	// 	log.Panic(err)
	// }
	// Db = db
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), pgxConnStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	Conn = conn
	//
	//var name string
	//var weight int64
	//err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(name, weight)
}

func GetConn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), pgxConnStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func Migrate() {

	//if err := Conn.Ping(context.Background()); err != nil {
	//	log.Fatal(err)
	//}
	//driver, _ := pgx.WithInstance(Db, &postgres.Config{})
	//m, _ := migrate.NewWithDatabaseInstance(
	//	"file://fg_hub/backend/db/rdb/migrations",
	//	"postgres",
	//	driver,
	//)
	//if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	//	log.Fatal(err)
	//}C:\Users\cidhi\go\src\fg_hub\backend\db\rdb\migrations
	//filePath := "file://fg_hub/backend/db/rdb/migrations"
	// for now, relative because it is breaking when called from another module
	filePath := "file://../../db/rdb/migrations"
	m, err := migrate.New(
		filePath,
		pgxMigrateStr)
	if err != nil {
		log.Fatal("Migration error", err)
	}
	if err := m.Up(); err != nil {
		log.Println("Up error", err)
	}
}
