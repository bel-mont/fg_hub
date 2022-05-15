package migrations

import (
	"context"
	"fg_hub/backend/db/rdb/models"
	"fmt"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println(" Run migration init ")
		_, err := db.NewCreateTable().
			Model((*models.Series)(nil)).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not create table Series, ", err)
		}
		_, err = db.NewCreateTable().
			Model((*models.Game)(nil)).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not create table Game, ", err)
		}
		_, err = db.NewCreateTable().
			Model((*models.Character)(nil)).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not create table Character, ", err)
		}
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Println(" Rollback migration init ")
		_, err := db.NewDropTable().
			Model((*models.Character)(nil)).
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not drop table Character, ", err)
		}
		_, err = db.NewDropTable().
			Model((*models.Game)(nil)).
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not drop table Game, ", err)
		}
		_, err = db.NewDropTable().
			Model((*models.Series)(nil)).
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not drop table Series, ", err)
		}
		return nil
	})
}
