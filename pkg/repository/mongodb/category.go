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

func FindCategory(Id string) (model.VehicleCategory, string, error) {
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

func FindCategoryByName(name string) (model.VehicleCategory, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.CategoryCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"name": name}

	var existingVehicle model.VehicleCategory
	findErr := collection.FindOne(ctx, filter).Decode(&existingVehicle)
	if findErr != nil {
		return model.VehicleCategory{}, fmt.Sprintf("Cannot decode"), fmt.Errorf(findErr.Error())
	}
	return existingVehicle, fmt.Sprintf("category found"), nil
}
func SaveCategory(category model.VehicleCategory) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.CategoryCollection
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
	collectionName := constants.CategoryCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"vehicle_id": Id}
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

func FindCategorys(page, sort string) ([]model.VehicleCategoryRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.VehicleCollection
	collection := database.GetCollection(client, databaseName, collectionName)

	//Checking if the querry are empty

	//it will be querry based on the created time
	perpage := int64(9)
	pageInt, _ := strconv.Atoi(page)
	if pageInt <= 0 {
		pageInt = 1
	}
	skippingLimit := (int64(pageInt) - 1) * perpage
	findOption := options.Find()
	findOption = findOption.SetSkip(skippingLimit)
	findOption = findOption.SetLimit(perpage)

	if sort != " " {
		if sort == "asc" {
			findOption = findOption.SetSort(bson.D{{"space_number", 1}})
		} else if sort == "desc" {
			findOption = findOption.SetSort(bson.D{{"space_number", -1}})
		}
	}
	filter := bson.M{}

	//	vehicle variable
	var vehicles []model.VehicleCategory
	var vehicleResponses []model.VehicleCategoryRes
	//database
	cursor, err := collection.Find(ctx, filter, findOption)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var vehicleCat model.VehicleCategory
		cursor.Decode(&vehicleCat)
		vehicles = append(vehicles, vehicleCat)
		vehicleResponse := model.VehicleCategoryRes{
			Name:              vehicleCat.Name,
			RatePerDay:        vehicleCat.RatePerDay,
			VehicleCategoryId: vehicleCat.VehicleCategoryId,
		}
		vehicleResponses = append(vehicleResponses, vehicleResponse)
	}

	return vehicleResponses, nil
}
