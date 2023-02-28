package reservation

import (
	"errors"
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/adeben33/vehicleParkingApi/pkg/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CreateReservation(reservation model.Reservation) (model.ReservationRes, string, error) {
	//save the payment
	_, _, err := mongodb.FindReservationByParkingSpace(reservation.ParkingSpace)
	if err == nil {
		return model.ReservationRes{}, fmt.Sprintf("The space is already reserved"), errors.New("The space is already reserved")
	}
	reservation.CreatedAt = time.Now().Local().Format(time.DateTime)
	reservation.UpdatedAt = time.Now().Local().Format(time.DateTime)
	reservation.Id = primitive.NewObjectID()
	reservation.ReservationId = reservation.Id.Hex()

	_, err = mongodb.CreateReservation(reservation)
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

func GetReservationById(reservationId string) (model.ReservationRes, string, error) {
	reservation, stmt, err := mongodb.GetReservationById(reservationId)
	if err != nil {
		return model.ReservationRes{}, stmt, err
	}

	reservationRes := model.ReservationRes{
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
	return reservationRes, fmt.Sprintf("Model generated"), nil
}

func DeleteReservation(reservationId string) (*mongo.DeleteResult, string, error) {
	deleteResult, stmt, err := mongodb.DeleteReservation(reservationId)
	if err != nil {
		return nil, stmt, err
	}
	return deleteResult, fmt.Sprintf("Payment deleted successfully"), nil
}

// Update payment
func UpdateReservation(reservation model.Reservation, reservationId string) (model.ReservationRes, error) {
	existingReservation, _, err := mongodb.GetReservationById(reservationId)
	if err != nil {
		return model.ReservationRes{}, err
	}

	if reservation.AmountPaid != 0 {
		existingReservation.AmountPaid = reservation.AmountPaid
	}
	if reservation.ParkingSpace != 0 {
		existingReservation.ParkingSpace = reservation.ParkingSpace
	}
	if reservation.Status != " " {
		existingReservation.Status = reservation.Status
	}

	if reservation.PaymentId != " " {
		existingReservation.PaymentId = reservation.PaymentId
	}
	if reservation.StartTime != " " {
		existingReservation.StartTime = reservation.StartTime
	}
	if reservation.EndTime != " " {
		existingReservation.EndTime = reservation.EndTime
	}
	if reservation.PaymentStatus != " " {
		existingReservation.PaymentStatus = reservation.PaymentStatus
	}
	if reservation.VehicleId != " " {
		existingReservation.VehicleId = reservation.VehicleId
	}
	existingReservation.UpdatedAt = time.Now().Local().Format(time.DateTime)

	_, err = mongodb.UpdateReservation(existingReservation, reservationId)

	if err != nil {
		return model.ReservationRes{}, err
	}
	response := model.ReservationRes{
		ParkingSpace:  existingReservation.ParkingSpace,
		VehicleId:     existingReservation.VehicleId,
		Status:        existingReservation.Status,
		StartTime:     existingReservation.StartTime,
		EndTime:       existingReservation.EndTime,
		AmountPaid:    existingReservation.AmountPaid,
		PaymentId:     existingReservation.PaymentId,
		PaymentStatus: existingReservation.PaymentStatus,
		ReservationId: existingReservation.ReservationId,
	}
	return response, nil
}

func GetReservation(search, page, sort string) ([]model.ReservationRes, string, error) {
	reservationRes, err := mongodb.FindReservations(search, page, sort)
	if err != nil {
		return []model.ReservationRes{}, fmt.Sprintf("paymentRes not generated"), nil
	}
	return reservationRes, fmt.Sprintf("paymentRes generated"), nil
}
