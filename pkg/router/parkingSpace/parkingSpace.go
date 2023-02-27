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
			ParkingSpaceUrl.POST("/parkingSpace", parkingSpace.CreateParkingSpace)
			ParkingSpaceUrl.GET("/parkingSpace/:parkingSpaceId", parkingSpace.GetParkingSpacebyId)
			ParkingSpaceUrl.PUT("/parkingSpace/:parkingSpaceId", parkingSpace.UpdateParkingSpace)
			ParkingSpaceUrl.DELETE("/parkingSpace/:parkingSpaceId", parkingSpace.DeleteParkingSpace)
			ParkingSpaceUrl.GET("/parkingSpace", parkingSpace.GetParkingSpaces)

		}
	}
	return r
}
