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

func UpdateReservation(reservation model.Reservation, reservationId string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.VehicleCollection
	collection := database.GetCollection(client, databaseName, collectionName)

	filter := bson.M{"reservation_id": reservationId}
	upsert := true
	updateOption := options.UpdateOptions{Upsert: &upsert}
	updateVehicle := bson.D{{"$set", bson.D{{"amount_paid", reservation.AmountPaid}, {"status", reservation.Status}, {"parking_space", reservation.ParkingSpace}, {"status", reservation.PaymentId}, {"start_time", reservation.StartTime}, {"end_time", reservation.EndTime}, {"payment_status", reservation.PaymentStatus}, {"vehicle_id", reservation.VehicleId}, {"update_at", reservation.UpdatedAt}}}}
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
	skippingLimit := (int64(pageInt) - 1) * perpage
	findOption := options.Find()
	findOption = findOption.SetSkip(skippingLimit)
	findOption = findOption.SetLimit(skippingLimit)

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
				{"user_id": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
				},
				{"reservation_id": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
				},
				{"vehicle_id": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
				},
				{"status": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
				},
				{"paymentI_id": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
				},
				{"start_time": bson.M{
					"$regex": primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				},
				},
				{"end_time": bson.M{
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
