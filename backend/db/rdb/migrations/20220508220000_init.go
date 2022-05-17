package migrations

import (
	"context"
	charModels "fg_hub/backend/db/rdb/modules/models"
	"fmt"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println(" Run migration init ")
		_, err := db.NewCreateTable().
			Model((*charModels.Series)(nil)).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not create table Series, ", err)
		}
		_, err = db.NewCreateTable().
			Model((*charModels.Game)(nil)).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not create table Game, ", err)
		}
		_, err = db.NewCreateTable().
			Model((*charModels.Character)(nil)).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not create table Character, ", err)
		}
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Println(" Rollback migration init ")
		_, err := db.NewDropTable().
			Model((*charModels.Character)(nil)).
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not drop table Character, ", err)
		}
		_, err = db.NewDropTable().
			Model((*charModels.Game)(nil)).
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not drop table Game, ", err)
		}
		_, err = db.NewDropTable().
			Model((*charModels.Series)(nil)).
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not drop table Series, ", err)
		}
		return nil
	})
}
