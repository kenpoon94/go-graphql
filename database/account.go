package database

import (
	"context"
	"log"

	"github.com/kenpoon94/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) Accounts() []*model.Account {
	cur := All(db, "accounts")
	var accounts []*model.Account
	if err := cur.All(context.Background(), &accounts); err != nil {
		log.Fatal(err)
	}
	return accounts
}

func (db *DB) Account(id string) *model.Account {
	res := FindById(db, "accounts", id)
	account := model.Account{}
	res.Decode(&account)
	return &account
}

func (db *DB) CreateAccount(input *model.NewAccount) *model.Account {
	res := Save(db, "accounts", input )
	return &model.Account{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}
}