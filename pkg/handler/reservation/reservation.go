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

func (base *Controller) UpdateReservation(c *gin.Context) {
	reservationId := c.Param("reservationId")
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
	CategoryRes, err := reservationService.UpdateReservation(reservation, reservationId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error(), "number": 1})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": CategoryRes})
}

func (base *Controller) GetReservationById(c *gin.Context) {
	reservationId := c.Param("reservationId")

	paymentRes, errString, err := reservationService.GetReservationById(reservationId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": paymentRes})
}

func (base *Controller) GetReservations(c *gin.Context) {
	//	this will use querry to get vehicles registered by the data created
	sort := c.Query("sort")
	search := c.Query("search")

	page := c.Query("page")

	vehicleResponse, errString, err := reservationService.GetReservation(search, page, sort)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": vehicleResponse})

}

func (base *Controller) DeleteReservation(c *gin.Context) {
	reservationId := c.Param("reservationId")

	result, errString, err := reservationService.DeleteReservation(reservationId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Details": result})

}
