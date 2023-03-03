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

func FindReservationByParkingSpace(parkingSpace uint16) (model.Reservation, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ReservationCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"parking_space": parkingSpace}

	var existingReservation model.Reservation
	findErr := collection.FindOne(ctx, filter).Decode(&existingReservation)
	if findErr != nil {
		return model.Reservation{}, fmt.Sprintf("No such reservation"), fmt.Errorf(findErr.Error())
	}

	return existingReservation, fmt.Sprintf("reservation found"), nil
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

func UpdateReservation(reservation model.Reservation, reservationId string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ReservationCollection
	collection := database.GetCollection(client, databaseName, collectionName)

	filter := bson.M{"reservation_id": reservationId}
	upsert := true
	updateOption := options.UpdateOptions{Upsert: &upsert}
	updateVehicle := bson.D{{"$set", bson.D{{"amount_paid", reservation.AmountPaid}, {"status", reservation.Status}, {"parking_space", reservation.ParkingSpace}, {"start_time", reservation.StartTime}, {"reservation_id", reservation.ReservationId}, {"end_time", reservation.EndTime}, {"payment_status", reservation.PaymentStatus}, {"vehicle_id", reservation.VehicleId}, {"update_at", reservation.UpdatedAt}}}}
	result, UpdateErr := collection.UpdateOne(ctx, filter, updateVehicle, &updateOption)
	if UpdateErr != nil {
		return nil, UpdateErr
	}
	return result, nil
}

func FindReservations(search, page, sort string) ([]model.ReservationRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.ReservationCollection
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
			findOption = findOption.SetSort(bson.D{{"reservation_id", 1}})
		} else if sort == "desc" {
			findOption = findOption.SetSort(bson.D{{"reservation_id", -1}})
		}
	}

	filter := bson.M{}
	if search != " " {
		filter = bson.M{
			"$or": []bson.M{
				{"user_id": search},
				{"reservation_id": search},
				{"vehicle_id": search},
				{"status": search},
				{"paymentI_id": search},
				{"start_time": search},
				{"end_time": search},
			},
		}
	}

	//	vehicle variable
	var reservations []model.Reservation
	var reservationsResponse []model.ReservationRes
	//database
	cursor, err := collection.Find(ctx, filter, findOption)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var reservation model.Reservation
		cursor.Decode(&reservation)
		reservations = append(reservations, reservation)
		reservationsRes := model.ReservationRes{
			ParkingSpace:  reservation.ParkingSpace,
			VehicleId:     reservation.VehicleId,
			Status:        reservation.Status,
			StartTime:     reservation.StartTime,
			EndTime:       reservation.EndTime,
			AmountPaid:    reservation.AmountPaid,
			PaymentId:     reservation.PaymentId,
			PaymentStatus: reservation.PaymentStatus,
			ReservationId: reservation.ReservationId,
		}
		reservationsResponse = append(reservationsResponse, reservationsRes)
	}

	return reservationsResponse, nil
}
