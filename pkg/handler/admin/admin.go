package admin

import "C"
import (
	"github.com/adeben33/vehicleParkingApi/internal/model"
	adminService "github.com/adeben33/vehicleParkingApi/service/admin"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Controller struct {
	Validate *validator.Validate
}

func (base *Controller) CreateCategory(c *gin.Context) {
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
	categoryRes, errstring, err := adminService.CreateCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errstring, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryRes})
}

func (base *Controller) GetCategory(c *gin.Context) {
	Id := c.Param("categoryId")

	categoryResponse, errString, err := adminService.GetCategory(Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})

}

func (base *Controller) UpdateCategory(c *gin.Context) {
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
	CategoryRes, err := adminService.UpdateCategory(category, categoryId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": CategoryRes})
}
func (base *Controller) DeleteCategory(c *gin.Context) {
	Id := c.Param("categoryId")

	categoryResponse, errString, err := adminService.DeleteCategory(Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": errString, "Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User Data": categoryResponse})
}
