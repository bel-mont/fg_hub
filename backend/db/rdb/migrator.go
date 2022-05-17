package rdb

import (
	"context"
	"fg_hub/backend/db/rdb/migrations"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun/migrate"
)

func Migrate() {
	db := NewConn()
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
