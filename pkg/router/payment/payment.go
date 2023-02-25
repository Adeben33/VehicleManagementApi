package payment

import (
	"fmt"
	paymentHandler "github.com/adeben33/vehicleParkingApi/pkg/handler/payment"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Payment(r *gin.Engine, validate *validator.Validate, APiVersion string) *gin.Engine {
	payment := paymentHandler.PaymentController{Validate: validate}

	paymentUrl := r.Group(fmt.Sprintf("/api/%v", APiVersion))
	{
		paymentUrl.POST("/payment", payment.CreatePayment)
		paymentUrl.GET("/payment/:paymentId", payment.GetPaymentById)
		paymentUrl.PUT("/payment/:paymentId", payment.UpdatePaymment)
		paymentUrl.DELETE("/payment", payment.DeletePayment)
		paymentUrl.GET("/payment", payment.GetPaymentById)

	}

	return r
}
