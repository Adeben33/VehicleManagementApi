package model

type Payment struct {
	Id            primitive.ObjectId `bson:"_id"`
	UserId        string             `json:"userId" bson:"user_id"`
	ReservationId string             `json:"reservationId" bson:"reservation_id"`
	Amount        string             `json:"amount" bson:"amount"`
	PaymentMethod string             `json:"paymentMethod" bson:"payment_method"`
	Status        string             `json:"status" bson:"status"`
	CreatedAt     string             `json:"createdAt" bson:"created_at"`
	UpdatedAt     string             `json:"updatedAt" bson:"updated_at"`
}
