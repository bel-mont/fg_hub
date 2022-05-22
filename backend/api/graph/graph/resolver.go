package graph

import "fg_hub/backend/api/graph/graph/model"

// This file will not be regenerated automatically.
//go:generate go run github.com/99designs/gqlgen generate
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	series     []*model.Series
	games      []*model.Game
	characters []*model.Character
}
