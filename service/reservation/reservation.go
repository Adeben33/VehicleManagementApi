package reservation

import (
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/adeben33/vehicleParkingApi/pkg/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CreateReservation(reservation model.Reservation) (model.ReservationRes, string, error) {
	//save the payment
	reservation.CreatedAt = time.Now().Local().Format(time.DateTime)
	reservation.UpdatedAt = time.Now().Local().Format(time.DateTime)
	reservation.Id = primitive.NewObjectID()
	reservation.PaymentId = reservation.Id.Hex()

	_, err := mongodb.CreateReservation(reservation)
	if err != nil {
		return model.ReservationRes{}, fmt.Sprintf("Error saving into database"), fmt.Errorf(err.Error())
	}
	reservationResponse := model.ReservationRes{
		ParkingSpace:  reservation.ParkingSpace,
		VehicleId:     reservation.VehicleId,
		Status:        reservation.Status,
		StartTime:     reservation.StartTime,
		EndTime:       reservation.EndTime,
		AmountPaid:    reservation.AmountPaid,
		PaymentId:     reservation.PaymentId,
		PaymentStatus: reservation.PaymentStatus,
		ReservationId: reservation.ReservationId,
	}
	return reservationResponse, fmt.Sprintf("Category saved successfully into the db"), nil
}

func GetPayment(paymentId string) (model.PaymentRes, string, error) {
	payment, stmt, err := mongodb.GetPayment(paymentId)
	if err != nil {
		return model.PaymentRes{}, stmt, err
	}
	paymentRes := model.PaymentRes{
		UserId:        payment.UserId,
		ReservationId: payment.ReservationId,
		Amount:        payment.Amount,
		PaymentMethod: payment.PaymentMethod,
		Status:        payment.Status,
		PaymentId:     payment.PaymentId,
	}
	return paymentRes, fmt.Sprintf("Model generated"), nil
}

func DeletePayment(paymentId string) (*mongo.DeleteResult, string, error) {
	deleteResult, stmt, err := mongodb.DeletePayment(paymentId)
	if err != nil {
		return nil, stmt, err
	}
	return deleteResult, fmt.Sprintf("Payment deleted successfully"), nil
}

// Update payment
func GetPayments(search, page, sort string) ([]model.PaymentRes, string, error) {
	paymentRes, err := mongodb.FindPayments(search, page, sort)
	if err != nil {
		return []model.PaymentRes{}, fmt.Sprintf("paymentRes not generated"), nil
	}
	return paymentRes, fmt.Sprintf("paymentRes generated"), nil
}
