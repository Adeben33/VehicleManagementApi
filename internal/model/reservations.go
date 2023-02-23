package model

type Reservation struct {
	Id            primitive.ObjectId `bson:"_id"`
	userId        string             `json:"userId" bson:"user_id"`
	ParkingLotId  string             `json:"parkingLotId" bson:"parking_lot_id"`
	VehicleId     string             `json:"vehicleId" bson:"vehicle_id"`
	StartTime     string             `json:"startTime" bson:"start_time"`
	EndTime       string             `json:"endTime" bson:"end_time"`
	AmountPaid    uint               `json:"amountPaid" bson:"amount_paid"`
	PaymentStatus string             `json:"paymentStatus" `
	CreatedAt     string             `json:"createdAt" bson:"created_at"`
	UpdatedAt     string             `json:"updatedAt" bson:"updated_at"`
	ReservationId string             `json:"reservationId" bson:"reservation_id"`
}
