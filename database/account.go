package database

import (
	"context"
	"log"
	"time"

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

func (db *DB) CreateAccount(input *model.NewAccount) string {
	currentTime := time.Now().String()
	newAccount := bson.M{
		"email":     input.Email,
		"password":  input.Password,
		"createdon": &currentTime,
		"updatedon": &currentTime,
	}

	res := Save(db, "accounts", newAccount)
	accountId := res.InsertedID.(primitive.ObjectID).Hex()

	newUser := &model.NewUser{
		Name:      input.Name,
		Jobtitle:  input.Jobtitle,
		City:      input.City,
		Age:       input.Age,
		Hobbies:   input.Hobbies,
		AccountID: &accountId,
		CreatedOn: &currentTime,
		UpdatedOn: &currentTime,
	}

	Save(db, "users", newUser)
	return res.InsertedID.(primitive.ObjectID).Hex()
}

func (db *DB) UpdateAccount(id string, fields bson.M) *model.Account {
	UpdateById(db, "accounts", id, fields)
	res := FindById(db, "accounts", id)
	account := model.Account{}
	res.Decode(&account)
	return &account
}
