package database

import (
	"context"
	"log"

	"github.com/kenpoon94/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
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

func (db *DB) CreateAccount(fields bson.M) string{
	res := Save(db, "accounts", fields)
	return res.InsertedID.(primitive.ObjectID).Hex()
}

func (db *DB) UpdateAccount(id string, fields bson.M) *model.Account{
	UpdateById(db, "accounts", id, fields)
	res := FindById(db, "accounts", id)
	account := model.Account{}
	res.Decode(&account)
	return &account
}
