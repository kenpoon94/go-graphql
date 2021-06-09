package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kenpoon94/go-graphql/graph/model"
	"github.com/kenpoon94/go-graphql/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var host = utils.GetEnvVariable("MONGODB_HOST")
var port = utils.GetEnvVariable("MONGODB_PORT")
var database = utils.GetEnvVariable("MONGODB_DATABASE")

type DB struct {
	client *mongo.Client
}


func Connect() *DB {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":" + port)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())


	// Check the connection
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongodDB!")

	return &DB{
		client: client,

	}

}

func (db* DB) Save(input *model.NewUser) *model.User{
	collection := db.client.Database(database).Collection("users")
	res, err := collection.InsertOne(context.Background(), input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.User{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}
}


func (db* DB) FindById(ID string) *model.User{
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database(database).Collection("users")
	res := collection.FindOne(context.Background(), bson.M{"_id": ObjectID})
	user := model.User{}
	res.Decode(&user)
	return &user
}

func (db* DB) All() []*model.User{
	collection := db.client.Database(database).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var users []*model.User
	for cur.Next(ctx){
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}