package database

import (
	"context"
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var (
	mongoClient *mongo.Client
)

// Connection this will help access the mongodb connection anywhere in the app
func Connection() (db *mongo.Client) {
	return mongoClient
}
func GetCollection(client *mongo.Client, databaseName, collectionName string) *mongo.Collection {
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}

func ConnectToDb() *mongo.Client {
	//the database url
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	url := config.GetConfig().Mongodb.Url
	mongoConnection := options.Client().ApplyURI(url)

	mongoClient, err = mongo.Connect(ctx, mongoConnection)
	//	Pinging the connection
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	//The connection is successful
	fmt.Println("Mongodb connected successfully")
	return mongoClient
}
