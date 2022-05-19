package service

import (
	"context"
	"fg_hub/backend/db/rdb"
	"fg_hub/backend/db/rdb/modules/models"
	coreModels "fg_hub/core/models/game"
)

func SaveSeries(game coreModels.Series) (*coreModels.Series, error) {
	conn := rdb.NewConn()
	newSeries := &models.Series{Name: game.Name, Slug: game.Slug}
	_, err := conn.NewInsert().Model(newSeries).On("CONFLICT (id) DO UPDATE").Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return &game, nil
}
