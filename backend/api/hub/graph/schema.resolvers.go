package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fg_hub/backend/api/hub/graph/generated"
	"fg_hub/backend/api/hub/graph/model"
	"log"

	"github.com/bel-mont/fg_hub/backend/db/rdb"
)

func (r *mutationResolver) CreateGame(ctx context.Context, input model.NewGame) (*model.Game, error) {
	game := &model.Game{
		ID:   input.ID,
		Name: input.Name,
	}
	conn := rdb.GetConn()
	_, err := conn.Exec(context.Background(), "INSERT INTO games(id, name) VALUES($1,$2)", input.ID, input.Name)
	if err != nil {
		log.Println("Exec error", err)
	}
	return game, nil
}

func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.NewCharacter) (*model.Character, error) {
	character := &model.Character{
		Name: input.Name,
	}
	conn := rdb.GetConn()
	_, err := conn.Exec(context.Background(), "INSERT INTO characters(name, game_id) VALUES($1,$2)", input.Name, input.GameID)
	if err != nil {
		log.Println("Exec error", err)
	}
	return character, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
