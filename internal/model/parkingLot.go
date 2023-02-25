package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//type Parkinglot struct {
//	Id            primitive.ObjectID `bson:"_id"`
//	Name          string             `json:"name" bson:"name"`
//	Address       string             `json:"address" bson:"address"`
//	Capacity      uint16             `json:"capacity" bson:"capacity"`
//	PricePerHour  uint16             `json:"pricePerHour" bson:"price_per_hour"`
//	ParkingSpaces []string           `json:"parkingSpaces" bson:"parking_spaces"`
//	CreatedAt     string             `json:"createdAt" bson:"created_at"`
//	UpdatedAt     string             `json:"updatedAt" bson:"updated_at"`
//	ParkinglotId  string             `json:"parkinglotId" bson:"parkinglot_id"`
//}

type ParkingSpace struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	SpaceNumber    uint16             `json:"spaceNumber" bson:"space_number"`
	Charges        uint16             `json:"charges" bson:"charges"`
	IsOccupied     bool               `json:"isOccupied" bson:"is_occupied"`
	OccupiedBy     string             `json:"occupiedBy" bson:"occupied_by"` //this is the user string
	VehicleId      string             `json:"vehicleId" bson:"vehicle_id"`
	ReservedBy     string             `json:"reservedBy" bson:"reserved_by"`
	ParkingSpaceId string             `json:"parkingSpaceId" bson:"parking_space_id"`
}
