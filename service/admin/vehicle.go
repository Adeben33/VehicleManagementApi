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
	existingVehicle, _, err := mongodb.FindVehicle(vehicle.PlateNumber)
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
	_, err = mongodb.UpdateVehicle(existingVehicle, id)
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
