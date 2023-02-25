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

func FindVehicle(Id string) (model.Vehicle, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.VehicleCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"vehicle_id": Id}

	var existingVehicle model.Vehicle
	findErr := collection.FindOne(ctx, filter).Decode(&existingVehicle)
	if findErr != nil {
		return model.Vehicle{}, fmt.Sprintf("No such vehicle"), fmt.Errorf(findErr.Error())
	}
	return existingVehicle, fmt.Sprintf("category found"), nil
}

func FindVehicleByPlateNumber(plateNumber string) (model.VehicleCategory, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.VehicleCollection
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
	collectionName := constants.VehicleCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	result, insertErr := collection.InsertOne(ctx, vehicle)
	if insertErr != nil {
		return nil, insertErr
	}
	return result, nil
}

func DeleteVehicle(Id string) (*mongo.DeleteResult, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.VehicleCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"vehicle_id": Id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Sprintf("Cannot decode"), err
	}
	return result, fmt.Sprintf("category found"), nil
}

func UpdateVehicle(vehicle model.Vehicle, id string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.VehicleCollection
	collection := database.GetCollection(client, databaseName, collectionName)

	filter := bson.M{"vehicle_id": id}
	upsert := true
	updateOption := options.UpdateOptions{Upsert: &upsert}
	updateVehicle := bson.D{{"$set", bson.D{{"make", vehicle.Make}, {"year", vehicle.Year}, {"model", vehicle.Model}, {"plate_number", vehicle.PlateNumber}, {"updated_at", vehicle.UpdatedAt}}}}
	result, UpdateErr := collection.UpdateOne(ctx, filter, updateVehicle, &updateOption)
	if UpdateErr != nil {
		return nil, UpdateErr
	}
	return result, nil
}
