package model

import (
	"fg_hub/backend/db/rdb/modules/game/service"
	coreModels "fg_hub/core/models/game"
)

type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Slug      string `json:"slug"`
}

func (game Game) Save() (*Game, error) {
	gameModel := coreModels.Game{Name: game.Name, Slug: game.Slug}
	_, err := service.Save(gameModel)
	if err != nil {
		return nil, err
	}
	return &game, nil
}
