package service

import (
	"context"
	"fg_hub/backend/db/rdb"
	"fg_hub/backend/db/rdb/modules/models"
	coreModels "fg_hub/core/models/game"
	"github.com/google/uuid"
)

func SaveGame(game coreModels.Game) (*coreModels.Game, error) {
	conn := rdb.NewConn()
	id, err := uuid.Parse(game.SeriesID)
	if err != nil {
		return nil, err
	}
	newGame := &models.Game{Name: game.Name, Slug: game.Slug, SeriesID: id}
	_, err = conn.NewInsert().Model(newGame).On("CONFLICT (id) DO UPDATE").Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return &game, nil
}
