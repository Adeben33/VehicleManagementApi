package parkingSpace

import (
	"github.com/adeben33/vehicleParkingApi/internal/model"
	parkingService "github.com/adeben33/vehicleParkingApi/service/parkingSpace"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Controller struct {
	Validate *validator.Validate
}

func (base *Controller) CreateParkingSpace(c *gin.Context) {
	var parkingSpace model.ParkingSpace
	err := c.BindJSON(&parkingSpace)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.Validate.Struct(&parkingSpace)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentRes, errstring, errString := parkingService.CreateParkingSpace(parkingSpace)

	if errString != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errstring, "Error": errString.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": paymentRes})
}

func (base *Controller) UpdateParkingSpace(c *gin.Context) {
	spacenumber := c.Param("spacenumber")
	var parkingSpace model.ParkingSpace
	err := c.BindJSON(&parkingSpace)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.Validate.Struct(&parkingSpace)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parkingRes, err := parkingService.UpdateParkingSpace(parkingSpace, spacenumber)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": parkingRes})
}

func (base *Controller) GetParkingSpaceBySpaceNumber(c *gin.Context) {
	spaceNumber := c.Param("spacenumber")

	parkingSpaceRes, errString, err := parkingService.GetParkingSpaceById(spaceNumber)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": parkingSpaceRes})
}

func (base *Controller) GetParkingSpaces(c *gin.Context) {
	//	this will use querry to get vehicles registered by the data created
	sort := c.Query("sort")
	search := c.Query("search")

	page := c.Query("page")

	vehicleResponse, errString, err := parkingService.GetParkingSpaces(search, page, sort)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": vehicleResponse})

}

func (base *Controller) DeleteParkingSpace(c *gin.Context) {
	spacenumber := c.Param("spacenumber")

	_, errString, err := parkingService.DeleteParkingSpace(spacenumber)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": errString})

}
