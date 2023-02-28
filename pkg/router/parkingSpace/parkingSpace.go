package parkingSpace

import (
	"fmt"
	parkingSpaceHandler "github.com/adeben33/vehicleParkingApi/pkg/handler/parkingSpace"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Parking(r *gin.Engine, validate *validator.Validate, ApiVersion string) *gin.Engine {
	parkingSpace := parkingSpaceHandler.Controller{Validate: validate}
	ParkingSpaceUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		{
			ParkingSpaceUrl.POST("/parkingspace", parkingSpace.CreateParkingSpace)
			ParkingSpaceUrl.GET("/parkingspace/:spacenumber", parkingSpace.GetParkingSpaceBySpaceNumber)
			ParkingSpaceUrl.PUT("/parkingspace/:spacenumber", parkingSpace.UpdateParkingSpace)
			ParkingSpaceUrl.DELETE("/parkingspace/:spacenumber", parkingSpace.DeleteParkingSpace)
			ParkingSpaceUrl.GET("/parkingspaces", parkingSpace.GetParkingSpaces)

		}
	}
	return r
}
