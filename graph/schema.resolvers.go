package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kenpoon94/go-graphql/database"
	"github.com/kenpoon94/go-graphql/graph/generated"
	"github.com/kenpoon94/go-graphql/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	return db.CreateUser(input), nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, input *model.NewAccount) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return db.User(id), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.Users(), nil
}

func (r *queryResolver) Account(ctx context.Context, id string) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
