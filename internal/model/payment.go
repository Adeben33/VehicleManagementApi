package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	Id            primitive.ObjectID `bson:"_id"`
	UserId        string             `json:"userId" bson:"user_id"`
	ReservationId string             `json:"reservationId" bson:"reservation_id"`
	Amount        uint16             `json:"amount" bson:"amount"`
	PaymentMethod string             `json:"paymentMethod" bson:"payment_method"`
	Status        string             `json:"status" bson:"status"`
	CreatedAt     string             `json:"createdAt" bson:"created_at"`
	UpdatedAt     string             `json:"updatedAt" bson:"updated_at"`
	PaymentId     string             `json:"paymentId" bson:"payment_id"`
}

type PaymentRes struct {
	UserId        string `json:"userId" bson:"user_id"`
	ReservationId string `json:"reservationId" bson:"reservation_id"`
	Amount        uint16 `json:"amount" bson:"amount"`
	PaymentMethod string `json:"paymentMethod" bson:"payment_method"`
	Status        string `json:"status" bson:"status"`
	PaymentId     string `json:"paymentId" bson:"payment_id"`
}

type Charge struct {
	Id           string `json:"id"`
	Amount       int    `json:"amount"`
	ReceiptEmail string `json:"Email"`
	ProductName  string `json:"productName"`
}
