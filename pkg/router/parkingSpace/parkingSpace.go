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
			ParkingSpaceUrl.GET("/payment/:paymentId", parkingSpace.GetParkingSpacebyId)
			ParkingSpaceUrl.PUT("/payment/:paymentId", parkingSpace.UpdateParkingSpace)
			ParkingSpaceUrl.DELETE("/payment", parkingSpace.DeleteParkingSpace)
			ParkingSpaceUrl.GET("/payment", parkingSpace.GetParkingSpace)

		}
	}

}
