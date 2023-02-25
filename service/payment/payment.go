package payment

import (
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/model"
	"github.com/adeben33/vehicleParkingApi/pkg/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CreatePayment(payment model.Payment) (model.PaymentRes, string, error) {
	//save the payment
	payment.CreatedAt = time.Now().Local().Format(time.DateTime)
	payment.UpdatedAt = time.Now().Local().Format(time.DateTime)
	payment.Id = primitive.NewObjectID()
	payment.PaymentId = payment.Id.Hex()
	_, err := mongodb.SavePayment(payment)
	if err != nil {
		return model.PaymentRes{}, fmt.Sprintf("Error saving into database"), fmt.Errorf(err.Error())
	}
	paymentResponse := model.PaymentRes{
		UserId:        payment.UserId,
		ReservationId: payment.ReservationId,
		Amount:        payment.Amount,
		PaymentMethod: payment.PaymentMethod,
		Status:        payment.Status,
		PaymentId:     payment.PaymentId,
	}
	return paymentResponse, fmt.Sprintf("Category saved successfully into the db"), nil
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

func GetPayments(search, page, sort string) ([]model.PaymentRes, string, error) {
	paymentRes, err := mongodb.FindPayments(search, page, sort)
	if err != nil {
		return []model.PaymentRes{}, fmt.Sprintf("paymentRes not generated"), nil
	}
	return paymentRes, fmt.Sprintf("paymentRes generated"), nil
}
