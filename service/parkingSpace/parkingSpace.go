package parkingSpace

import (
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/adeben33/vehicleParkingApi/pkg/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CreateParkingSpace(space model.ParkingSpace) (model.ParkingSpace, string, error) {
	//save the payment
	space.CreatedAt = time.Now().Local().Format(time.DateTime)
	space.UpdatedAt = time.Now().Local().Format(time.DateTime)
	space.Id = primitive.NewObjectID()
	space.ParkingSpaceId = space.Id.Hex()
	_, err := mongodb.CreateParkingSpace(space)
	if err != nil {
		return model.ParkingSpace{}, fmt.Sprintf("Error saving into database"), fmt.Errorf(err.Error())
	}

	return space, fmt.Sprintf("Category saved successfully into the db"), nil
}

func GetParkingSpaceById(paymentId string) (model.ParkingSpaceRes, string, error) {
	payment, stmt, err := mongodb.GetParkingSpaceById(paymentId)
	if err != nil {
		return model.ParkingSpaceRes{}, stmt, err
	}
	parkingSpaceRes := model.ParkingSpaceRes{
		SpaceNumber:    payment.SpaceNumber,
		Charges:        payment.Charges,
		IsOccupied:     payment.IsOccupied,
		VehicleId:      payment.VehicleId,
		ReservedBy:     payment.ReservedBy,
		ParkingSpaceId: payment.ParkingSpaceId,
	}
	return parkingSpaceRes, fmt.Sprintf("Model generated"), nil
}

func DeleteParkingSpace(paymentId string) (*mongo.DeleteResult, string, error) {
	deleteResult, stmt, err := mongodb.DeleteParkingSpace(paymentId)
	if err != nil {
		return nil, stmt, err
	}
	return deleteResult, fmt.Sprintf("Payment deleted successfully"), nil
}

func GetParkingSpaces(search, page, sort string) ([]model.ParkingSpaceRes, string, error) {
	parkingSpaceRes, err := mongodb.FindParkingSpace(search, page, sort)
	if err != nil {
		return []model.ParkingSpaceRes{}, fmt.Sprintf("paymentRes not generated"), nil
	}
	return parkingSpaceRes, fmt.Sprintf("paymentRes generated"), nil
}
