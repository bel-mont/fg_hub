package rdb

// Switch to https://github.com/volatiletech/sqlboiler
// Migrations with https://github.com/rubenv/sql-migrate
// https://github.com/golang-migrate/migrate
// https://www.howtographql.com/graphql-go/1-getting-started/
// https://github.com/graph-gophers/graphql-go check out graphql setups
// https://www.apollographql.com/blog/graphql/golang/using-graphql-with-golang/
// https://www.howtographql.com/basics/3-big-picture/
import (
	"github.com/golang-migrate/migrate/v4"
	"log"

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

//var Conn *pgx.Conn
//var databaseUrl = "host=localhost user=postgres password=root dbname=fghub port=5432 sslmode=disable TimeZone=Asia/Tokyo"
var databaseUrl2 = "pgx://postgres:root@localhost:5432/fghub?sslmode=disable"

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
	//conn, err := pgx.Connect(context.Background(), databaseUrl2)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//defer conn.Close(context.Background())
	//Conn = conn
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
		databaseUrl2)
	if err != nil {
		log.Fatal("Migration error", err)
	}
	if err := m.Up(); err != nil {
		log.Println("Up error", err)
	}
}
