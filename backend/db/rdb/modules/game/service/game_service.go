package service

import (
	"context"
	"fg_hub/backend/db/rdb"
	"fg_hub/backend/db/rdb/modules/models"
	coreModels "fg_hub/core/models/game"
)

func Save(game coreModels.Game) (*coreModels.Game, error) {
	conn := rdb.NewConn()
	newGame := models.Game{Name: game.Name, Slug: game.Slug}
	_, err := conn.NewInsert().Model(newGame).On("CONFLICT (id) DO UPDATE").Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return &game, nil
}
