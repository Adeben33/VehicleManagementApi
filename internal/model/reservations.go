package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reservation struct {
	Id            primitive.ObjectID `bson:"_id"`
	userId        string             `json:"userId" bson:"user_id"`
	ParkingSpace  uint16             `json:"parkingSpace" bson:"parking_space"`
	VehicleId     string             `json:"vehicleId" bson:"vehicle_id"`
	Status        string             `json:"status" bson:"status"` //parked,completed,vacant
	StartTime     string             `json:"startTime" bson:"start_time"`
	EndTime       string             `json:"endTime" bson:"end_time"`
	AmountPaid    uint               `json:"amountPaid" bson:"amount_paid"`
	PaymentId     string             `json:"paymentId" bson:"payment_id"`
	PaymentStatus string             `json:"paymentStatus" `
	CreatedAt     string             `json:"createdAt" bson:"created_at"`
	UpdatedAt     string             `json:"updatedAt" bson:"updated_at"`
	ReservationId string             `json:"reservationId" bson:"reservation_id"`
}

type ReservationRes struct {
	userId        string `json:"userId" bson:"user_id"`
	ParkingSpace  uint16 `json:"parkingSpace" bson:"parking_space"`
	VehicleId     string `json:"vehicleId" bson:"vehicle_id"`
	Status        string `json:"status" bson:"status"` //parked,completed,vacant
	StartTime     string `json:"startTime" bson:"start_time"`
	EndTime       string `json:"endTime" bson:"end_time"`
	AmountPaid    uint   `json:"amountPaid" bson:"amount_paid"`
	PaymentId     string `json:"paymentId" bson:"payment_id"`
	PaymentStatus string `json:"paymentStatus" `
	ReservationId string `json:"reservationId" bson:"reservation_id"`
}
