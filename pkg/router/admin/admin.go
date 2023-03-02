package admin

import (
	"fmt"
	adminHandler "github.com/adeben33/vehicleParkingApi/pkg/handler/admin"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Admin(r *gin.Engine, validate *validator.Validate, ApiVersion string) *gin.Engine {
	admin := adminHandler.Controller{Validate: validate}

	adminUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))

	//Vehicle administration only admin can add vehicle

	{
		adminUrl.POST("/vehicle", admin.CreateVehicle)
		adminUrl.GET("/vehicle/:vehicleId", admin.GetVehicle)
		adminUrl.PUT("/vehicle/:vehicleId", admin.UpdateVehicle)
		adminUrl.DELETE("/vehicle/:vehicleId", admin.DeleteVehicle)
		adminUrl.GET("/vehicle", admin.GetVehicles)
	}

	//	Category
	{
		adminUrl.POST("/category", admin.CreateCategory)
		adminUrl.GET("/category/:categoryId", admin.GetCategory)
		adminUrl.PUT("/category/:categoryId", admin.UpdateCategory)
		adminUrl.DELETE("/category/:categoryId", admin.DeleteCategory)
		adminUrl.GET("/category", admin.GetCategorys)
	}
	//DashBoard
	{
		adminUrl.GET("/Dashboard/getvehicles", admin.GetVehicles)
		adminUrl.GET("/Dashboard/getvehicle/:parkingSpaceNumber", admin.GetVehicleByParkingNumber)
	}
	//ManageVehicle
	{
		adminUrl.GET("/manageVehicle/entryReport")
		//adminUrl.POST("/manageVehicle/incoming/:vehiclePlateNumber", admin.IncomingVehicle)
		adminUrl.GET("/manageVehicle/outcoming/:vehiclePlateNumber")
	}

	return r
}
