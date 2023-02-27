package admin

import (
	"errors"
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/adeben33/vehicleParkingApi/pkg/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CreateVehicle(vehicle model.Vehicle) (model.VehicleRes, string, error) {
	//	check if the categoryexist
	_, stmt, err := mongodb.FindVehicleByPlateNumber(vehicle.PlateNumber)
	if err == nil {
		return model.VehicleRes{}, stmt, errors.New("Vehicle is already in the database")
	}

	//Then save the category
	vehicle.CreatedAt = time.Now().Local().Format(time.DateTime)
	vehicle.UpdatedAt = time.Now().Local().Format(time.DateTime)
	vehicle.Id = primitive.NewObjectID()
	vehicle.VehicleId = vehicle.Id.Hex()
	_, err = mongodb.SaveVehicle(vehicle)
	if err != nil {
		return model.VehicleRes{}, fmt.Sprintf("Error saving into database"), fmt.Errorf(err.Error())
	}
	vehicleResponse := model.VehicleRes{
		VehicleCategoryId: vehicle.VehicleCategoryId,
		Make:              vehicle.Make,
		Model:             vehicle.Model,
		Year:              vehicle.Year,
		PlateNumber:       vehicle.PlateNumber,
		VehicleId:         vehicle.VehicleId,
	}
	return vehicleResponse, fmt.Sprintf("Category saved successfully into the db"), nil
}

func GetVehicle(id string) (model.VehicleRes, string, error) {

	vehicleCategory, stmt, err := mongodb.FindVehicle(id)
	if err != nil {
		return model.VehicleRes{}, stmt, err
	}
	vehicleRes := model.VehicleRes{
		VehicleCategoryId: vehicleCategory.VehicleCategoryId,
		Make:              vehicleCategory.Make,
		Model:             vehicleCategory.Make,
		Year:              vehicleCategory.Year,
		PlateNumber:       vehicleCategory.PlateNumber,
		VehicleId:         vehicleCategory.VehicleId,
	}
	return vehicleRes, fmt.Sprintf("Model generated"), nil
}

func DeleteVehicle(id string) (*mongo.DeleteResult, string, error) {
	//Taking to the database
	deleteResult, stmt, err := mongodb.DeleteVehicle(id)
	if err != nil {
		return nil, stmt, err
	}
	return deleteResult, fmt.Sprintf("Category deleted successfully"), nil
}

func UpdateVehicle(vehicle model.Vehicle, id string) (model.VehicleRes, error) {
	existingVehicle, _, err := mongodb.FindVehicle(vehicle.VehicleId)
	if err != nil {
		return model.VehicleRes{}, err
	}
	if vehicle.Year != " " {
		existingVehicle.Year = vehicle.Year
	}
	if vehicle.Make != " " {
		existingVehicle.Make = vehicle.Make
	}
	if vehicle.Model != " " {
		existingVehicle.Model = vehicle.Model
	}
	if vehicle.Color != " " {
		existingVehicle.Color = vehicle.Color
	}
	existingVehicle.UpdatedAt = time.Now().Local().Format(time.DateTime)
	_, err = mongodb.UpdateVehicle(existingVehicle, vehicle.VehicleId)
	if err != nil {
		return model.VehicleRes{}, err
	}
	response := model.VehicleRes{
		VehicleCategoryId: existingVehicle.VehicleId,
		Make:              existingVehicle.Make,
		Model:             existingVehicle.Model,
		Year:              existingVehicle.Year,
		PlateNumber:       existingVehicle.PlateNumber,
		VehicleId:         existingVehicle.VehicleId,
	}
	return response, nil
}

func GetVehicles(timeLow, timeHigh, page, sort string) ([]model.VehicleRes, string, error) {
	vehicleRes, err := mongodb.FindVehicles(timeLow, timeHigh, page, sort)
	if err != nil {
		return []model.VehicleRes{}, fmt.Sprintf("vehicleRes not generated"), nil
	}
	return vehicleRes, fmt.Sprintf("vehicleRes generated"), nil
}

func GetVehicleByParkingSpaceNumber(spaceNumber string) (model.VehicleRes, string, error) {
	//	find the vehicleId from parkingspace collection
	vehicleId, stmt, err := mongodb.FindVehicleIdByParkingSpaceNumber(spaceNumber)
	if err != nil {
		return model.VehicleRes{}, stmt, err
	}
	//	find the vehicleId from the vehicle collection
	vehicle, stmt, err := mongodb.FindVehicle(vehicleId)
	if err != nil {
		return model.VehicleRes{}, stmt, err
	}
	vehicleRes := model.VehicleRes{
		VehicleCategoryId: vehicle.VehicleCategoryId,
		Make:              vehicle.Make,
		Model:             vehicle.Model,
		Year:              vehicle.Year,
		PlateNumber:       vehicle.PlateNumber,
		VehicleId:         vehicle.VehicleId,
	}
	return vehicleRes, stmt, nil
}

//func IncomingVehicle(incoming model.IncomingVehicle, days string) {
//	plateNumber := incoming.PlateNumber
//	spaceNumber := incoming.ParkingSpaceNumber
//	vehicle, stmt, err := mongodb.FindVehicleByPlateNumber(plateNumber)
//	//	get the vehicle id to get if the vehicle has be reserved to parked
//	reservation,stmt,err := mongodb.FindReservationByVehicleId(vehicle.VehicleId)
//	//if there is no reservation
//	if err != nil {
//	//	create a reservation
//	//	reservation variable
//		reserved := model.Reservation{
//			parkingSpace: spaceNumber,
//			VehicleId : vehicle.VehicleId,
//			Status: "Occupied",
//			StartTime: time.Now().Local().Format(time.DateTime),
//			EndTime:reservation
//			AmountPaid:
//			PaymentStatus
//			PaymentId
//			CreatedAt
//			UpdaatedAt
//		}
//		mongodb.CreateReservation(vehicle.VehicleId,spaceNumber,days)
//	}
//
//}
