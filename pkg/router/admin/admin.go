package admin

import (
	"fmt"
	adminHandler "github.com/adeben33/vehicleParkingApi/pkg/handler/admin"
	"github.com/adeben33/vehicleParkingApi/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Admin(r *gin.Engine, validate *validator.Validate, ApiVersion string) *gin.Engine {
	admin := adminHandler.Controller{Validate: validate}

	adminUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion)).Use(middleware.Authorization())

	//Vehicle administration only admin can add vehicle
	{
		adminUrl.POST("/vehicle", adminHandler.CreateVehicle)
		adminUrl.GET("/vehicle", adminHandler.GetVehicle)
		adminUrl.PUT("/vehicle", adminHandler.UpdateVehicle)
		adminUrl.DELETE("/vehicle", adminHandler.DeleteVehicle)
	}

	//	Category
	{
		adminUrl.POST("/Category", adminHandler.CreateCategory)
		adminUrl.GET("/Category", adminHandler.GetCategory)
		adminUrl.PUT("/category", adminHandler.UpdateCategory)
		adminUrl.DELETE("/category", adminHandler.DeleteCatogry)
	}

	return r
}
