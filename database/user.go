package database

import (
	"context"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kenpoon94/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) Users() []*model.User {
	cur := All(db, "users")
	var users []*model.User
	if err := cur.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}
	return users
}

func (db *DB) User(ctx context.Context, id string) *model.User {
	if !primitive.IsValidObjectID(id) {
		graphql.AddErrorf(ctx, "ID is not a valid primitive.ObjectID")
	}
	res := FindById(db, "users", id)
	user := model.User{}
	res.Decode(&user)
	return &user
}

func (db *DB) UpdateUser(ctx context.Context, input *model.UpdateUser) *model.User {
	update := false
	updateUser := bson.M{}
	if !primitive.IsValidObjectID(input.ID) {
		graphql.AddErrorf(ctx, "ID is not a valid primitive.ObjectID")
	}

	if input.Name != nil {
		updateUser["name"] = input.Name
		update = true
	}
	if input.Age != nil {
		updateUser["age"] = input.Age
		update = true
	}
	if input.Jobtitle != nil {
		updateUser["jobtitle"] = input.Jobtitle
		update = true
	}
	if input.City != nil {
		updateUser["city"] = input.City
		update = true
	}
	if input.Hobbies != nil {
		updateUser["hobbies"] = input.Hobbies
		update = true
	}

	if !update {
		graphql.AddErrorf(ctx, "No updateUser present for updating")
		return nil
	} else {
		updateUser["updatedon"] = time.Now().String()
	}

	UpdateById(db, "users", input.ID, updateUser)
	res := FindById(db, "users", input.ID)
	user := model.User{}
	res.Decode(&user)
	return &user
}
