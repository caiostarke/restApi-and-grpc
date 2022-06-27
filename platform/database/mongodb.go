package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURL = "mongodb://localhost:27017"
)

func MongoDBConnection() (*mongo.Client, error) {
	// Connect to MongoDB
	fmt.Println("Connecting to mongoDB")

	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	})

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("error connecting: ", err)
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return client, nil
}
