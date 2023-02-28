package mongodb

import (
	"context"
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/config"
	"github.com/adeben33/vehicleParkingApi/internal/constants"
	"github.com/adeben33/vehicleParkingApi/internal/database"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

func CreateParkingSpace(space model.ParkingSpace) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ParkingSpaceCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	result, insertErr := collection.InsertOne(ctx, space)
	if insertErr != nil {
		return nil, insertErr
	}
	return result, nil
}

func GetParkingSpaceById(parkingSpaceId string) (model.ParkingSpace, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ParkingSpaceCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"parking_space_id": parkingSpaceId}

	var parkingSpace model.ParkingSpace
	findErr := collection.FindOne(ctx, filter).Decode(&parkingSpace)
	if findErr != nil {
		return model.ParkingSpace{}, fmt.Sprintf("No such vehicle"), fmt.Errorf(findErr.Error())
	}

	return parkingSpace, fmt.Sprintf("parkingSpace found"), nil
}

func GetParkingSpaceBySpaceNumber(spaceNumber uint16) (model.ParkingSpace, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ParkingSpaceCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"space_number": spaceNumber}

	var parkingSpace model.ParkingSpace
	findErr := collection.FindOne(ctx, filter).Decode(&parkingSpace)
	if findErr != nil {
		return model.ParkingSpace{}, fmt.Sprintf("No such parking space"), fmt.Errorf(findErr.Error())
	}

	return parkingSpace, fmt.Sprintf("parkingSpace found"), nil
}

func DeleteParkingSpace(spaceNumber uint16) (*mongo.DeleteResult, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ParkingSpaceCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"space_number": spaceNumber}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Sprintf("not deleted"), err
	}
	return result, fmt.Sprintf("Parking space deleted"), nil
}

func GetParkingSpaces(search, page, sort string) ([]model.ParkingSpaceRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ParkingSpaceCollection
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
	findOption = findOption.SetLimit(skippingLimit)

	if sort != " " {
		if sort == "asc" {
			findOption = findOption.SetSort(bson.D{{"space_number", 1}})
		} else if sort == "desc" {
			findOption = findOption.SetSort(bson.D{{"space_number", -1}})
		}
	}
	filter := bson.M{}
	if search != " " {

		filter = bson.M{
			"$or": []bson.M{
				{"Vehicle_id ": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
				},
				{"reserved_by ": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
				},
				{"parking_space_id": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
				},
			},
		}
	}

	//	vehicle variable
	var parkingSpaces []model.ParkingSpace
	var parkingSpaceResponse []model.ParkingSpaceRes
	//database
	cursor, err := collection.Find(ctx, filter, findOption)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var parkingSpace model.ParkingSpace
		var parkingSpaceRes model.ParkingSpaceRes
		cursor.Decode(&parkingSpace)
		parkingSpaces = append(parkingSpaces, parkingSpace)
		parkingSpaceRes.SpaceNumber = parkingSpace.SpaceNumber
		parkingSpaceRes.Charges = parkingSpace.Charges
		parkingSpaceRes.IsOccupied = parkingSpace.IsOccupied
		parkingSpaceRes.OccupiedBy = parkingSpace.OccupiedBy
		parkingSpaceRes.VehicleId = parkingSpace.VehicleId
		parkingSpaceRes.ReservedBy = parkingSpace.ReservedBy
		parkingSpaceResponse = append(parkingSpaceResponse, parkingSpaceRes)
	}

	return parkingSpaceResponse, nil

}

func UpdateParkingSpace(parkingSpace model.ParkingSpace, spacenumber uint16) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ParkingSpaceCollection
	collection := database.GetCollection(client, databaseName, collectionName)

	filter := bson.M{"space_number": spacenumber}
	upsert := true
	updateOption := options.UpdateOptions{Upsert: &upsert}
	updateVehicle := bson.D{{"$set", bson.D{{"space_number", parkingSpace.SpaceNumber}, {"currency", parkingSpace.Currency}, {"occupied_by", parkingSpace.OccupiedBy}, {"is_occupied", parkingSpace.IsOccupied}, {"charges", parkingSpace.Charges}, {"vehicle_id", parkingSpace.VehicleId}, {"reserved_by", parkingSpace.ReservedBy}, {"updated_at", parkingSpace.UpdatedAt}}}}
	result, UpdateErr := collection.UpdateOne(ctx, filter, updateVehicle, &updateOption)
	if UpdateErr != nil {
		return nil, UpdateErr
	}
	return result, nil
}
