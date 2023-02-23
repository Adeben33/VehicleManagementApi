package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	FirstName  string             `json:"firstName" bson:"first_name"`
	LastName   string             `json:"lastName" bson:"last_name"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"password" bson:"password"`
	Role       string             `json:"role" bson:"role"`
	VehiclesId string             `json:"vehiclesId" bson:"vehicles_id"`
	CreatedAt  string             `json:"created_at "bson:"created_at"`
	UpdatedAt  string             `json:"updatedAt" bson:"updated_at"`
	UserId     string             ` json:"userId" bson:"user_id"`
}

type UserRes struct {
	FirstName  string `json:"firstName" bson:"first_name"`
	LastName   string `json:"lastName" bson:"last_name"`
	Email      string `json:"email" bson:"email"`
	VehiclesId string `json:"vehiclesId" bson:"vehicles_id"`
}
