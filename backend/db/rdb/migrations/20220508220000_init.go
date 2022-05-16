package migrations

import (
	"context"
	charModels "fg_hub/backend/db/rdb/modules/character/models"
	gameModels "fg_hub/backend/db/rdb/modules/game/models"
	"fmt"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println(" Run migration init ")
		_, err := db.NewCreateTable().
			Model((*gameModels.Series)(nil)).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not create table Series, ", err)
		}
		_, err = db.NewCreateTable().
			Model((*gameModels.Game)(nil)).
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
			Model((*gameModels.Game)(nil)).
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not drop table Game, ", err)
		}
		_, err = db.NewDropTable().
			Model((*gameModels.Series)(nil)).
			Exec(ctx)
		if err != nil {
			fmt.Println("Could not drop table Series, ", err)
		}
		return nil
	})
}
