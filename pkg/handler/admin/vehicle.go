package admin

import (
	"github.com/adeben33/vehicleParkingApi/internal/model"
	adminService "github.com/adeben33/vehicleParkingApi/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (base *Controller) CreateVehicle(c *gin.Context) {
	var vehicle model.Vehicle
	err := c.BindJSON(&vehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.Validate.Struct(&vehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryRes, errstring, errString := adminService.CreateVehicle(vehicle)

	if errString != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errstring, "Error": errString.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryRes})
}

func (base *Controller) GetVehicle(c *gin.Context) {
	Id := c.Param("vehicleId")

	categoryResponse, errString, err := adminService.GetVehicle(Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})

}

func (base *Controller) UpdateVehicle(c *gin.Context) {
	vehicleId := c.Param("vehicleId")
	var vehicle model.Vehicle
	err := c.BindJSON(&vehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.Validate.Struct(&vehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	CategoryRes, err := adminService.UpdateVehicle(vehicle, vehicleId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error(), "number": 1})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": CategoryRes})
}
func (base *Controller) DeleteVehicle(c *gin.Context) {
	Id := c.Param("vehicleId")

	categoryResponse, errString, err := adminService.DeleteVehicle(Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})
}

func (base *Controller) GetVehicles(c *gin.Context) {
	//	this will use querry to get vehicles registered by the data created
}
