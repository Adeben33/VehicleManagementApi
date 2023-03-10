package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vehicle struct {
	Id                primitive.ObjectID `json:"_id" bson:"_id"`
	VehicleCategoryId string             `json:"vehicleCategoryId" bson:"vehicle_category_id"`
	Color             string             `json:"color" bson:"color"`
	Make              string             `json:"make" bson:"make"`
	Model             string             `json:"model" bson:"model"`
	Year              string             `json:"year" bson:"year"`
	PlateNumber       string             `json:"plateNumber" bson:"plate_number"`
	UserId            string             `json:"userId" bson:"user_id"`
	CreatedAt         string             `json:"createdAT" bson:"created_at"`
	UpdatedAt         string             `json:"updatedAt" bson:"updated_at"`
	VehicleId         string             `json:"vehicleId" bson:"vehicle_id"`
}

type VehicleRes struct {
	VehicleCategoryId string `json:"vehicleCategoryId" `
	Color             string `json:"color" `
	Make              string `json:"make" `
	Model             string `json:"model"`
	Year              string `json:"year" `
	PlateNumber       string `json:"plateNumber" `
	VehicleId         string `json:"vehicleId" `
}
type VehicleCategory struct {
	Id                primitive.ObjectID `bson:"_id"`
	Name              string             `json:"name" bson:"name"`
	RatePerDay        int                `json:"ratePerDay" bson:"rate_per_day"`
	CreatedAt         string             `json:"createdAt"bson:"created_at"`
	UpdatedAT         string             `json:"updatedAT" bson:"updated_at"`
	VehicleCategoryId string             `json:"vehicleId" bson:"vehicle_id"`
}

type VehicleCategoryRes struct {
	Name              string
	RatePerDay        int
	VehicleCategoryId string
}

type IncomingVehicle struct {
	PlateNumber        string
	ParkingSpaceNumber string
}
