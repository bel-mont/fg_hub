package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fg_hub/backend/api/graph/graph/generated"
	"fg_hub/backend/api/graph/graph/model"
	"fmt"
	"log"
)

func (r *characterResolver) Game(ctx context.Context, obj *model.Character) (*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSeries(ctx context.Context, input model.NewSeries) (*model.Series, error) {
	series := &model.Series{
		Slug: input.Slug,
		Name: input.Name,
	}
	_, err := series.Save()

	if err != nil {
		log.Println("Exec error", err)
	}
	return series, nil
}

func (r *mutationResolver) CreateGame(ctx context.Context, input model.NewGame) (*model.Game, error) {
	game := &model.Game{
		Slug:     input.Slug,
		Name:     input.Name,
		SeriesID: input.SeriesID,
	}
	_, err := game.Save()

	if err != nil {
		log.Println("Exec error", err)
	}
	return game, nil
}

func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.NewCharacter) (*model.Character, error) {
	character := &model.Character{
		Name: input.Name,
	}
	//conn, _ := rdb.GetConn()
	//_, err := conn.Exec(context.Background(), "INSERT INTO characters(name, game_id) VALUES($1,$2)", input.Name, input.GameID)
	//if err != nil {
	//	log.Println("Exec error", err)
	//}
	return character, nil
}

// Character returns generated.CharacterResolver implementation.
func (r *Resolver) Character() generated.CharacterResolver { return &characterResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type characterResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
