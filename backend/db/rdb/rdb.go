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
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func Test() {
	m, _ := migrate.New(
		"github://mattes:personal-access-token@mattes/migrate_test",
		"postgres://postgres:root@localhost:5432/fghub-db?sslmode=disable")
	// host=localhost user=postgres password=root dbname=fghub-db port=5432 sslmode=disable TimeZone=Asia/Tokyo
	m.Steps(2)
}
