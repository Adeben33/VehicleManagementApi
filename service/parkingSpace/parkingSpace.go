package parkingSpace

import (
	"errors"
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/adeben33/vehicleParkingApi/pkg/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

func CreateParkingSpace(space model.ParkingSpace) (model.ParkingSpace, string, error) {
	//save the payment
	_, _, err := mongodb.GetParkingSpaceBySpaceNumber(space.SpaceNumber)
	if err == nil {
		return model.ParkingSpace{}, fmt.Sprintf("parking space already exist"), errors.New("parking Space already registered")
	}
	space.CreatedAt = time.Now().Local().Format(time.DateTime)
	space.UpdatedAt = time.Now().Local().Format(time.DateTime)
	space.Id = primitive.NewObjectID()
	space.ParkingSpaceId = space.Id.Hex()
	_, err = mongodb.CreateParkingSpace(space)
	if err != nil {
		return model.ParkingSpace{}, fmt.Sprintf("Error saving into database"), fmt.Errorf(err.Error())
	}

	return space, fmt.Sprintf("Category saved successfully into the db"), nil
}

func GetParkingSpaceById(spaceNumber string) (model.ParkingSpaceRes, string, error) {
	spaceNumberInt, _ := strconv.Atoi(spaceNumber)

	payment, stmt, err := mongodb.GetParkingSpaceBySpaceNumber(uint16(spaceNumberInt))
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

func DeleteParkingSpace(spaceNumber string) (*mongo.DeleteResult, string, error) {
	spaceNumberInt, _ := strconv.Atoi(spaceNumber)
	deleteResult, stmt, err := mongodb.DeleteParkingSpace(uint16(spaceNumberInt))
	if err != nil {
		return nil, stmt, err
	}
	return deleteResult, fmt.Sprintf("parking space deleted successfully"), nil
}

func GetParkingSpaces(search, page, sort string) ([]model.ParkingSpaceRes, string, error) {
	parkingSpaceRes, err := mongodb.GetParkingSpaces(search, page, sort)
	if err != nil {
		return []model.ParkingSpaceRes{}, fmt.Sprintf("paymentRes not generated"), nil
	}
	return parkingSpaceRes, fmt.Sprintf("paymentRes generated"), nil
}

func UpdateParkingSpace(parkingSpace model.ParkingSpace, spacenumber string) (model.ParkingSpaceRes, error) {
	spacenumberInt, _ := strconv.Atoi(spacenumber)

	existingParkingSpace, _, err := mongodb.GetParkingSpaceBySpaceNumber(uint16(spacenumberInt))
	if err != nil {
		return model.ParkingSpaceRes{}, err
	}
	if parkingSpace.Charges != 0 {
		existingParkingSpace.Charges = parkingSpace.Charges
	}
	if parkingSpace.SpaceNumber != 0 {
		existingParkingSpace.SpaceNumber = parkingSpace.SpaceNumber
	}
	if parkingSpace.ReservedBy != " " {
		existingParkingSpace.ReservedBy = parkingSpace.ReservedBy
	}
	if parkingSpace.OccupiedBy != " " {
		existingParkingSpace.OccupiedBy = parkingSpace.OccupiedBy
	}
	if parkingSpace.VehicleId != " " {
		existingParkingSpace.VehicleId = parkingSpace.VehicleId
	}
	if parkingSpace.ReservedBy != " " {
		existingParkingSpace.ReservedBy = parkingSpace.ReservedBy
	}

	if parkingSpace.IsOccupied != false {
		existingParkingSpace.IsOccupied = parkingSpace.IsOccupied
	}
	if parkingSpace.Currency != " " {
		existingParkingSpace.Currency = parkingSpace.Currency
	}
	existingParkingSpace.UpdatedAt = time.Now().Local().Format(time.DateTime)
	_, err = mongodb.UpdateParkingSpace(existingParkingSpace, uint16(spacenumberInt))
	if err != nil {
		return model.ParkingSpaceRes{}, err
	}

	response := model.ParkingSpaceRes{
		SpaceNumber:    existingParkingSpace.SpaceNumber,
		Charges:        existingParkingSpace.Charges,
		IsOccupied:     existingParkingSpace.IsOccupied,
		OccupiedBy:     existingParkingSpace.OccupiedBy,
		VehicleId:      existingParkingSpace.VehicleId,
		ReservedBy:     existingParkingSpace.ReservedBy,
		ParkingSpaceId: existingParkingSpace.ParkingSpaceId,
		Currency:       existingParkingSpace.Currency,
	}
	return response, nil
}
