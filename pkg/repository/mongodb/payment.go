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

func SavePayment(payment model.Payment) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.PaymentCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	result, insertErr := collection.InsertOne(ctx, payment)
	if insertErr != nil {
		return nil, insertErr
	}
	return result, nil
}

func GetPayment(paymentId string) (model.Payment, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.PaymentCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"payment_id": paymentId}

	var payment model.Payment
	findErr := collection.FindOne(ctx, filter).Decode(&payment)
	if findErr != nil {
		return model.Payment{}, fmt.Sprintf("No such vehicle"), fmt.Errorf(findErr.Error())
	}

	return payment, fmt.Sprintf("Vehicle found"), nil
}

func DeletePayment(paymentId string) (*mongo.DeleteResult, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.PaymentCollection
	collection := database.GetCollection(client, databaseName, collectionName)
	filter := bson.M{"payment_id": paymentId}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Sprintf("Cannot decode"), err
	}
	return result, fmt.Sprintf("category found"), nil
}

func FindPayments(search, page, sort string) ([]model.PaymentRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := database.Connection()
	databaseName := config.GetConfig().Mongodb.Database
	collectionName := constants.PaymentCollection
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
			findOption = findOption.SetSort(bson.D{{"payment_id", 1}})
		} else if sort == "desc" {
			findOption = findOption.SetSort(bson.D{{"payment_id", -1}})
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
				{"paymentMethod": bson.M{
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
				{"paymentId": bson.M{
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
	var payments []model.Payment
	var paymentResponse []model.PaymentRes
	//database
	cursor, err := collection.Find(ctx, filter, findOption)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var payment model.Payment
		var paymentRes model.PaymentRes
		cursor.Decode(&payment)
		payments = append(payments, payment)
		paymentRes.UserId = payment.UserId
		paymentRes.ReservationId = payment.ReservationId
		paymentRes.Amount = payment.Amount
		paymentRes.PaymentMethod = payment.PaymentMethod
		paymentRes.Status = payment.Status
		paymentRes.PaymentId = payment.PaymentId
		paymentResponse = append(paymentResponse, paymentRes)
	}

	return paymentResponse, nil
}
