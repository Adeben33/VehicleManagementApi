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

func GetVehicle(id string) (model.VehicleCategoryRes, string, error) {

	vehicleCategory, stmt, err := mongodb.FindCategory(id)
	if err != nil {
		return model.VehicleCategoryRes{}, stmt, err
	}
	vehicleRes := model.VehicleCategoryRes{
		Name:              vehicleCategory.Name,
		RatePerDay:        vehicleCategory.RatePerDay,
		VehicleCategoryId: vehicleCategory.VehicleCategoryId,
	}
	return vehicleRes, fmt.Sprintf("Model generated"), nil
}

func DeleteVehicle(id string) (*mongo.DeleteResult, string, error) {
	//Taking to the database
	deleteResult, stmt, err := mongodb.DeleteCategory(id)
	if err != nil {
		return nil, stmt, err
	}
	return deleteResult, fmt.Sprintf("Category deleted successfully"), nil
}

func UpdateVehicle(category model.VehicleCategory, id string) (model.VehicleCategoryRes, error) {
	existingCategory, _, err := mongodb.FindCategory(id)
	if err != nil {
		return model.VehicleCategoryRes{}, err
	}
	if category.Name != " " {
		existingCategory.Name = category.Name
	}
	if category.RatePerDay != 0 {
		existingCategory.RatePerDay = category.RatePerDay
	}
	existingCategory.UpdatedAT = time.Now().Local().Format(time.DateTime)
	_, err = mongodb.UpdateCategory(existingCategory, id)
	if err != nil {
		return model.VehicleCategoryRes{}, err
	}
	response := model.VehicleCategoryRes{
		Name:              existingCategory.Name,
		RatePerDay:        existingCategory.RatePerDay,
		VehicleCategoryId: existingCategory.VehicleCategoryId,
	}
	return response, nil
}
