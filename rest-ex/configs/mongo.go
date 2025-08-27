package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database = ConnectDB()

func ConnectDB() *mongo.Database {
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client.Database("go-test-db")
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := DB.Collection(collectionName)
	return collection
}
