package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kenpoon94/go-graphql/database"
	"github.com/kenpoon94/go-graphql/graph/generated"
	"github.com/kenpoon94/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = database.Connect()

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	currentTime := time.Now().String()
	newUser := &model.NewUser{
		Name:      input.Name,
		Jobtitle:  input.Jobtitle,
		City:      input.City,
		Age:       input.Age,
		Hobbies:   input.Hobbies,
		CreatedOn: &currentTime,
		UpdatedOn: &currentTime,
	}
	return db.CreateUser(newUser), nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, input *model.NewAccount) (*model.Account, error) {
	currentTime := time.Now().String()
	newAccount := &model.NewAccount{
		Email:     input.Email,
		Password:  input.Password,
		CreatedOn: &currentTime,
		UpdatedOn: &currentTime,
	}
	return db.CreateAccount(newAccount), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {

	update := false
	var fields = bson.M{}
	if !primitive.IsValidObjectID(input.ID) {
		graphql.AddErrorf(ctx, "ID is not a valid primitive.ObjectID")
	}

	if input.Name != nil {
		fields["name"] = input.Name
		update = true
	}
	if input.Age != nil {
		fields["age"] = input.Age
		update = true
	}
	if input.Jobtitle != nil {
		fields["jobtitle"] = input.Jobtitle
		update = true
	}
	if input.City != nil {
		fields["city"] = input.City
		update = true
	}
	if input.Hobbies != nil {
		fields["hobbies"] = input.Hobbies
		update = true
	}

	if !update {
		graphql.AddErrorf(ctx, "No fields present for updating")
		return nil, nil
	} else {
		fields["updatedon"] = time.Now().String()
	}
	
	return db.UpdateUser(input.ID, fields), nil
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, input *model.UpdateAccount) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	if !primitive.IsValidObjectID(id) {
		graphql.AddErrorf(ctx, "ID is not a valid primitive.ObjectID")
	}
	return db.User(id), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.Users(), nil
}

func (r *queryResolver) Account(ctx context.Context, id string) (*model.Account, error) {
	return db.Account(id), nil
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*model.Account, error) {
	return db.Accounts(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
