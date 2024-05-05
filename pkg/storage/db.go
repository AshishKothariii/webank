package storage

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func  GetSession() *mongo.Client {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_uri"))
		// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	return client;
}
