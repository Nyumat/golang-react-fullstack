package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnect connects to the database and returns a client pointer
func MongoConnect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// Create an instance of the client (will probably need later for creating new collections)
var DB *mongo.Client = MongoConnect()

// GetCollection returns a collection from the database given the name of the collection
func GetMongoCollection(client *mongo.Client, collection string) *mongo.Collection {
	return client.Database("go-todo").Collection(collection)
}

