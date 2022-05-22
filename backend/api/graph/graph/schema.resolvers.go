package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fg_hub/backend/api/graph/graph/generated"
	"fg_hub/backend/api/graph/graph/model"
	seriesService "fg_hub/backend/db/rdb/modules/game/service"
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

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Series(ctx context.Context) ([]*model.Series, error) {
	series, err := seriesService.Find()
	if err != nil {
		return nil, err
	}
	var convertedModels []*model.Series
	for _, v := range series {
		convertedModels = append(convertedModels, &model.Series{Name: v.Name, ID: v.ID, Slug: v.Slug, CreatedAt: v.CreatedAt})
	}
	return convertedModels, nil
}

func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

// Character returns generated.CharacterResolver implementation.
func (r *Resolver) Character() generated.CharacterResolver { return &characterResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type characterResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
