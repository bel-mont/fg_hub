package model

import (
	"fg_hub/backend/db/rdb/modules/game/service"
	coreModels "fg_hub/core/models/game"
)

type Series struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Slug      string `json:"slug"`
}

func (series Series) Save() (*Series, error) {
	seriesModel := coreModels.Series{Name: series.Name, Slug: series.Slug}
	_, err := service.SaveSeries(seriesModel)
	if err != nil {
		return nil, err
	}
	return &series, nil
}
