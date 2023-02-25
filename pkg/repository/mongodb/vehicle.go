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

func FindVehicle(Id string) (model.VehicleCategory, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.CategoryCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"vehicle_id": Id}

	var existingVehicle model.VehicleCategory
	findErr := collection.FindOne(ctx, filter).Decode(&existingVehicle)
	if findErr != nil {
		return model.VehicleCategory{}, fmt.Sprintf("Cannot decode"), fmt.Errorf(findErr.Error())
	}
	return existingVehicle, fmt.Sprintf("category found"), nil
}

func FindVehicleByPlateNumber(plateNumber string) (model.VehicleCategory, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.CategoryCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"plate_number": plateNumber}

	var existingVehicle model.VehicleCategory
	findErr := collection.FindOne(ctx, filter).Decode(&existingVehicle)
	if findErr != nil {
		return model.VehicleCategory{}, fmt.Sprintf("Cannot decode"), fmt.Errorf(findErr.Error())
	}
	return existingVehicle, fmt.Sprintf("category found"), nil
}

func SaveVehicle(vehicle model.Vehicle) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.CategoryCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	result, insertErr := collection.InsertOne(ctx, vehicle)
	if insertErr != nil {
		return nil, insertErr
	}
	return result, nil
}

func DeleteDelete(Id string) (*mongo.DeleteResult, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.CategoryCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"vehicle_id": Id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Sprintf("Cannot decode"), err
	}
	return result, fmt.Sprintf("category found"), nil
}

func UpdateVehicle(category model.VehicleCategory, id string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.CategoryCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"vehicle_id": id}
	upsert := true
	updateOption := options.UpdateOptions{Upsert: &upsert}
	updateVehicle := bson.D{{"$set", bson.D{{"name", category.Name}, {"updated_at", category.UpdatedAT}, {"rate_per_day", category.RatePerDay}}}}
	result, UpdateErr := collection.UpdateOne(ctx, filter, updateVehicle, &updateOption)
	if UpdateErr != nil {
		return nil, UpdateErr
	}
	return result, nil
}
