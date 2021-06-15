package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kenpoon94/go-graphql/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	host     = utils.GetEnvVariable("MONGODB_HOST")
	port     = utils.GetEnvVariable("MONGODB_PORT")
	database = utils.GetEnvVariable("MONGODB_DATABASE")
)

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

func Save(db *DB, col string, class interface{}) *mongo.InsertOneResult {
	collection := db.client.Database(database).Collection(col)
	res, err := collection.InsertOne(context.Background(), class)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func Find(db *DB, col string, fields bson.M) *mongo.SingleResult {
	collection := db.client.Database(database).Collection(col)
	res := collection.FindOne(context.Background(), fields)
	return res
}

func FindById(db *DB, col string, id string) *mongo.SingleResult {
	ObjectID, _ := primitive.ObjectIDFromHex(id)
	collection := db.client.Database(database).Collection(col)
	res := collection.FindOne(context.Background(), bson.M{"_id": ObjectID})
	return res
}

func All(db *DB, col string) mongo.Cursor {
	collection := db.client.Database(database).Collection(col)
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	return *cur
}

func UpdateById(db *DB, col string, id string, fields bson.M) {
	ObjectId, _ := primitive.ObjectIDFromHex(id)
	collection := db.client.Database(database).Collection(col)
	_, err := collection.UpdateByID(context.Background(), ObjectId, bson.M{"$set": fields})
	if err != nil {
		log.Fatal(err)
	}
}
