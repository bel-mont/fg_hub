package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fg_hub/backend/api/hub/graph/generated"
	"fg_hub/backend/api/hub/graph/model"
	"fmt"
	"log"

	"github.com/bel-mont/fg_hub/backend/db/rdb"
)

func (r *gameResolver) Slug(ctx context.Context, obj *model.Game) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateGame(ctx context.Context, input model.NewGame) (*model.Game, error) {
	game := &model.Game{
		Slug: input.Slug,
		Name: input.Name,
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
	conn, _ := rdb.GetConn()
	_, err := conn.Exec(context.Background(), "INSERT INTO characters(name, game_id) VALUES($1,$2)", input.Name, input.GameID)
	if err != nil {
		log.Println("Exec error", err)
	}
	return character, nil
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type gameResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
