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
	var series []models.Series
	err := conn.NewSelect().Model(&series).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	var convertedModels []coreModels.Series
	for _, v := range series {
		convertedModels = append(convertedModels, coreModels.Series{Name: v.Name, ID: v.ID.String(), Slug: v.Slug, CreatedAt: v.CreatedAt.String()})
	}
	return convertedModels, nil
}
