package mongodb

import (
	"context"
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/config"
	"github.com/adeben33/vehicleParkingApi/internal/constants"
	"github.com/adeben33/vehicleParkingApi/internal/database"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FindCategory(Id string) (model.VehicleCategory, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.UserCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"vehicleId": Id}
	var existingVehicle model.VehicleCategory
	err := collection.FindOne(ctx, filter).Decode(&existingVehicle)
	if err != nil {
		return model.VehicleCategory{}, fmt.Sprintf("Cannot decode"), err
	}
	return existingVehicle, fmt.Sprintf("category found"), nil
}

func SaveCategory(category model.VehicleCategory) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.UserCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	result, insertErr := collection.InsertOne(ctx, category)
	if insertErr != nil {
		return nil, insertErr
	}
	return result, nil
}

func DeleteCategory(Id string) (*mongo.DeleteResult, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.UserCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"vehicleId": Id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Sprintf("Cannot decode"), err
	}
	return result, fmt.Sprintf("category found"), nil
}

func UpdateCategory(category model.VehicleCategory, id string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.UserCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"category_id": id}
	upsert := true
	updateOption := options.UpdateOptions{Upsert: &upsert}
	result, UpdateErr := collection.UpdateOne(ctx, filter, category, &updateOption)
	if UpdateErr != nil {
		return nil, UpdateErr
	}
	return result, nil
}
