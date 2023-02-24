package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vehicle struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	VehicleType string             `json:"vehicleType" bson:"vehicle_type"`
	Color       string             `json:"color" bson:"color"`
	Make        string             `json:"make" bson:"make"`
	Model       string             `json:"model" bson:"model"`
	Year        string             `json:"year" bson:"year"`
	UserId      string             `json:"userId" bson:"user_id"`
	CreatedAT   string             `json:"createdAT" bson:"created_at"`
	UpdatedAt   string             `json:"updatedAt" bson:"updated_at"`
}

type VehicleCategory struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	RatePerDay uint64             `json:"ratePerDay" bson:"rate_per_day"`
	CreatedAt  string             `json:"createdAt"bson:"created_at"`
	UpdatedAT  string             `json:"updatedAT" bson:"updated_at"`
}
