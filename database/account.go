package database

import (
	"context"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kenpoon94/go-graphql/graph/model"
	"github.com/kenpoon94/go-graphql/utils"
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

func (db *DB) Account(ctx context.Context, id string) *model.Account {
	if !primitive.IsValidObjectID(id) {
		graphql.AddErrorf(ctx, "ID is not a valid primitive.ObjectID")
	}
	res := FindById(db, "accounts", id)
	account := model.Account{}
	res.Decode(&account)
	return &account
}

func (db *DB) CreateAccount(ctx context.Context, input *model.NewAccount) *model.Account {
	accountExist := Find(db, "accounts", bson.M{"email": input.Email})
	account := model.Account{}
	accountExist.Decode(&account)
	if account.Email == input.Email {
		graphql.AddErrorf(ctx, "Email is already used")
		return &model.Account{}
	}

	currentTime := time.Now().String()
	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Fatalf("Unexpected error when hashing password")
	}
	newAccount := bson.M{
		"email":     input.Email,
		"password":  hashPassword,
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
	return &model.Account{
		ID: accountId,
	}
}

func (db *DB) UpdateAccount(ctx context.Context, input *model.UpdateAccount) *model.Account {
	update := false
	updateAccount := bson.M{}

	if !primitive.IsValidObjectID(input.ID) {
		graphql.AddErrorf(ctx, "ID is not a valid primitive.ObjectID")
	}

	if input.Email != nil {
		updateAccount["email"] = input.Email
		update = true
	}
	if input.Password != nil {
		hashPassword, err := utils.HashPassword(*input.Password)
		if err != nil {
			log.Fatalf("Unexpected error when hashing password")
		}
		updateAccount["password"] = hashPassword
		update = true
	}

	if !update {
		graphql.AddErrorf(ctx, "No updateAccount present for updating")
		return nil
	} else {
		updateAccount["updatedon"] = time.Now().String()
	}

	UpdateById(db, "accounts", input.ID, updateAccount)
	res := FindById(db, "accounts", input.ID)
	account := model.Account{}
	res.Decode(&account)
	return &account
}
