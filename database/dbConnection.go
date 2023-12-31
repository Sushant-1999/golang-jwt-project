package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

/*
mongo.Connect() accepts a Context and a options.ClientOptions object, which is used to set the connection string and other driver settings. You can visit the options package documentation to see what configuration options are available.

Context is like a timeout or deadline that indicates when an operation should stop running and return. It helps to prevent performance degradation on production systems when specific operations are running slow. In this code, you’re passing context.TODO() to indicate that you’re not sure what context to use right now, but you plan to add one in the future.
*/

func ConnectMongoDB() *mongo.Client {
	//Load the configurations first from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env File")
	}
	MongoDb := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(MongoDb)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDb")
	return client
}

var Client *mongo.Client = ConnectMongoDB()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
