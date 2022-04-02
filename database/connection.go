package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	MongoDbURL := "mongodb://localhost:27017"
	fmt.Print(MongoDbURL)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDbURL))
	if err != nil {
		log.Fatal(err)
	}
	//! Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	connectErr := client.Connect(ctx)
	if connectErr != nil {
		log.Fatal(connectErr)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("restaurant").Collection(collectionName)
	return collection
}
