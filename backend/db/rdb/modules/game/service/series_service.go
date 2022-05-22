package service

import (
	"context"
	"fg_hub/backend/db/rdb"
	"fg_hub/backend/db/rdb/modules/models"
	coreModels "fg_hub/core/models/game"
)

func SaveSeries(series coreModels.Series) (*coreModels.Series, error) {
	conn := rdb.NewConn()
	newSeries := &models.Series{Name: series.Name, Slug: series.Slug}
	_, err := conn.NewInsert().Model(newSeries).On("CONFLICT (id) DO UPDATE").Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return &series, nil
}

func Find() ([]coreModels.Series, error) {
	conn := rdb.NewConn()
	var series []coreModels.Series
	_, err := conn.NewSelect().Model(&series).Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return series, nil
}
