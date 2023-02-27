package reservation

import (
	"github.com/adeben33/vehicleParkingApi/internal/model"
	reservationService "github.com/adeben33/vehicleParkingApi/service/reservation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Controller struct {
	Validate *validator.Validate
}

func (base *Controller) CreateReservation(c *gin.Context) {
	var reservation model.Reservation
	err := c.BindJSON(&reservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.Validate.Struct(&reservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentRes, errstring, errString := reservationService.CreateReservation(reservation)

	if errString != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errstring, "Error": errString.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": paymentRes})
}

func (base *Controller) UpdatePaymment(c *gin.Context) {
}

func (base *Controller) GetPaymentById(c *gin.Context) {
	paymentId := c.Param("paymentId")

	paymentRes, errString, err := reservationService.GetPayment(paymentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": paymentRes})
}

func (base *Controller) GetPayments(c *gin.Context) {
	//	this will use querry to get vehicles registered by the data created
	sort := c.Query("sort")
	search := c.Query("search")

	page := c.Query("page")

	vehicleResponse, errString, err := reservationService.GetPayments(search, page, sort)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": vehicleResponse})

}

func (base *Controller) DeletePayment(c *gin.Context) {
	paymentId := c.Param("paymentId")

	categoryResponse, errString, err := reservationService.DeletePayment(paymentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})

}
