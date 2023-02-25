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
	"log"
	"strconv"
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

	return existingVehicle, fmt.Sprintf("Vehicle found"), nil
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

func FindVehicles(timeLow, timeHigh, page, sort string) ([]model.VehicleRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.VehicleCollection
	collection := database.GetCollection(client, databaseName, collectionName)

	filter := bson.D{{"$and", []bson.D{
		{{"created_at", bson.D{{"$gte", timeLow}}}},
		{{"created_at", bson.D{{"$lte", timeHigh}}}}}},
	}
	//it will be querry based on the created time
	perpage := int64(9)
	pageInt, _ := strconv.Atoi(page)
	skippingLimit := (int64(pageInt) - 1) * perpage
	findOption := options.Find()
	findOption = findOption.SetSkip(skippingLimit)
	findOption = findOption.SetLimit(skippingLimit)
	findOption = findOption.SetSort(bson.D{{"vehicle_id", 1}})

	//	vehicle variable
	var vehicles []model.Vehicle
	var vehiclesRes []model.VehicleRes
	//database
	cursor, err := collection.Find(ctx, filter, findOption)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var vehicle model.Vehicle
		var vehicleRes model.VehicleRes
		cursor.Decode(&vehicle)
		vehicles = append(vehicles, vehicle)
		vehicleRes.VehicleId = vehicle.VehicleId
		vehicleRes.VehicleCategoryId = vehicle.VehicleCategoryId
		vehicleRes.PlateNumber = vehicle.PlateNumber
		vehicleRes.Year = vehicle.Year
		vehicleRes.Make = vehicle.Make
		vehicleRes.Model = vehicle.Model
		vehiclesRes = append(vehiclesRes, vehicleRes)
	}

	return vehiclesRes, nil
}

func FindVehicleIdByParkingSpaceNumber(spaceNumber string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ParkingSpaceCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"space_number": spaceNumber}

	var parkingSpace model.ParkingSpace
	findErr := collection.FindOne(ctx, filter).Decode(parkingSpace)
	if findErr != nil {
		return " ", fmt.Sprintf("No such ParkingSpace"), fmt.Errorf(findErr.Error())
	}
	return parkingSpace.VehicleId, fmt.Sprintf("category found"), nil
}
