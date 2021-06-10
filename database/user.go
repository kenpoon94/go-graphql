package database

import (
	"context"
	"log"

	"github.com/kenpoon94/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db* DB) Users()[]*model.User {
	cur := All(db, "users")	
	var users []*model.User
	if err := cur.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}
	return users
}

func (db* DB) User(id string) *model.User{
	res := FindById(db, "users", id)
	user := model.User{}
	res.Decode(&user)
	return &user
}

func (db* DB) CreateUser(input *model.NewUser) *model.User{
	res := Save(db, input, "users")
	return &model.User{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}
}