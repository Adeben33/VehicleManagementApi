package payment

import (
	"github.com/adeben33/vehicleParkingApi/internal/model"
	paymentService "github.com/adeben33/vehicleParkingApi/service/payment"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type PaymentController struct {
	Validate *validator.Validate
}

func (base *PaymentController) CreatePayment(c *gin.Context) {
	var payment model.Payment
	err := c.BindJSON(&payment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.Validate.Struct(&payment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentRes, errstring, errString := paymentService.CreatePayment(payment)

	if errString != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errstring, "Error": errString.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": paymentRes})
}

func (base *PaymentController) UpdatePayment(c *gin.Context) {
	paymentId := c.Param("paymentId")
	var payment model.Payment

	err := c.BindJSON(&payment)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = base.Validate.Struct(&payment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentRes, err := paymentService.UpdatePayment(paymentId, payment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": paymentRes})
}

func (base *PaymentController) GetPaymentById(c *gin.Context) {
	paymentId := c.Param("paymentId")

	paymentRes, errString, err := paymentService.GetPayment(paymentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": paymentRes})
}

func (base *PaymentController) GetPayments(c *gin.Context) {
	//	this will use querry to get vehicles registered by the data created
	sort := c.Query("sort")
	search := c.Query("search")

	page := c.Query("page")

	vehicleResponse, errString, err := paymentService.GetPayments(search, page, sort)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": vehicleResponse})

}

func (base *PaymentController) DeletePayment(c *gin.Context) {
	paymentId := c.Param("paymentId")

	categoryResponse, errString, err := paymentService.DeletePayment(paymentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})

}
