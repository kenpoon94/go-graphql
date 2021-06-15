package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kenpoon94/go-graphql/database"
	"github.com/kenpoon94/go-graphql/graph/generated"
	"github.com/kenpoon94/go-graphql/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateAccount(ctx context.Context, input *model.NewAccount) (*model.Account, error) {
	return db.CreateAccount(ctx, input), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	return db.UpdateUser(ctx, input), nil
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, input *model.UpdateAccount) (*model.Account, error) {
	return db.UpdateAccount(ctx, input), nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return db.User(ctx, id), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.Users(), nil
}

func (r *queryResolver) Account(ctx context.Context, id string) (*model.Account, error) {
	return db.Account(ctx, id), nil
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*model.Account, error) {
	return db.Accounts(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
