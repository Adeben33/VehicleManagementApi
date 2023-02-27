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
	"time"
)

func CreateReservation(reservation model.Reservation) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ReservationCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	result, insertErr := collection.InsertOne(ctx, reservation)
	if insertErr != nil {
		return nil, insertErr
	}
	return result, nil
}

func GetReservationById(reservationId string) (model.Reservation, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ReservationCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"reservation_id": reservationId}

	var existingReservation model.Reservation
	findErr := collection.FindOne(ctx, filter).Decode(&existingReservation)
	if findErr != nil {
		return model.Reservation{}, fmt.Sprintf("No such vehicle"), fmt.Errorf(findErr.Error())
	}

	return existingReservation, fmt.Sprintf("Vehicle found"), nil
}

func DeleteReservation(reservationId string) (*mongo.DeleteResult, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ReservationCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"reservation_id": reservationId}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Sprintf("Cannot decode"), err
	}
	return result, fmt.Sprintf("category found"), nil
}
