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
	Id := c.Param("categoryId")

	categoryResponse, errString, err := adminService.GetVehicle(Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})

}

func (base *Controller) UpdateVehicle(c *gin.Context) {
	categoryId := c.Param("categoryId")
	var category model.VehicleCategory
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.Validate.Struct(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	CategoryRes, err := adminService.UpdateVehicle(category, categoryId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": CategoryRes})
}
func (base *Controller) DeleteVehicle(c *gin.Context) {
	Id := c.Param("categoryId")

	categoryResponse, errString, err := adminService.DeleteVehicle(Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})
}
