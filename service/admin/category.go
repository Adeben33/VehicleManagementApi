package admin

import (
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/adeben33/vehicleParkingApi/pkg/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CreateCategory(category model.VehicleCategory) (model.VehicleCategoryRes, string, error) {
	//	check if the categoryexist
	_, stmt, err := mongodb.FindCategory(category.VehicleCategoryId)
	if err == nil {
		return model.VehicleCategoryRes{}, stmt, err
	}
	//Then save the category
	category.CreatedAt = time.Now().Local().Format(time.DateTime)
	category.CreatedAt = time.Now().Local().Format(time.DateTime)
	_, err = mongodb.SaveCategory(category)
	if err != nil {
		return model.VehicleCategoryRes{}, fmt.Sprintf("Error saving into database"), err
	}
	vehicleResponse := model.VehicleCategoryRes{
		Name:       category.Name,
		RatePerDay: category.RatePerDay,
	}
	return vehicleResponse, fmt.Sprintf("Category saved successfully into the db"), nil
}

func GetCategory(id string) (model.VehicleCategoryRes, string, error) {

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

func DeleteCategory(id string) (*mongo.DeleteResult, string, error) {
	//Taking to the database
	deleteResult, stmt, err := mongodb.DeleteCategory(id)
	if err != nil {
		return nil, stmt, err
	}
	return deleteResult, fmt.Sprintf("Category deleted successfully"), nil
}

func UpdateCategory(category model.VehicleCategory, id string) (model.VehicleCategoryRes, error) {
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
	existingCategory.CreatedAt = time.Now().Local().Format(time.DateTime)
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
