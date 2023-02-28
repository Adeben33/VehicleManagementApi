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
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
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
	sort := c.Query("sort")
	search := c.Query("search")

	page := c.Query("page")

	vehicleResponse, errString, err := adminService.GetVehicles(search, page, sort)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": vehicleResponse})
}

func (base *Controller) GetVehiclesWithTime(c *gin.Context) {
	//	this will use querry to get vehicles registered by the data created
	sort := c.Query("sort")
	timeLow := c.Query("timeLow")
	timeHigh := c.Query("timeHigh")
	page := c.Query("page")

	vehicleResponse, errString, err := adminService.GetVehiclesWithTime(timeLow, timeHigh, page, sort)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": vehicleResponse})

}

func (base *Controller) GetVehicleByParkingNumber(c *gin.Context) {
	spaceNumber := c.Param("parkingSpaceNumber")
	categoryResponse, errString, err := adminService.GetVehicleByParkingSpaceNumber(spaceNumber)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})

}

//func (base *Controller) IncomingVehicle(c *gin.Context) {
//	//
//	var vehicle model.IncomingVehicle
//	days := c.Query("days")
//	categoryResponse, errString, err := adminService.IncomingVehicle(vehicle, days)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})
//
//}
