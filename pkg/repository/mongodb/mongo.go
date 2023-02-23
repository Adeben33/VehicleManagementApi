package mongodb

import (
	"context"
	"github.com/adeben33/vehicleParkingApi/internal/config"
	"github.com/adeben33/vehicleParkingApi/internal/constants"
	"github.com/adeben33/vehicleParkingApi/internal/database"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func FindUser(userEmail string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.UserCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"email": userEmail}

	var existingUser model.User
	err := collection.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		return model.User{}, err
	}
	return existingUser, nil
}

func SaveUser(user model.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.UserCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
