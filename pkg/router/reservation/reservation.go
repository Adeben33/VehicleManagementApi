package reservation

import (
	"fmt"
	reservationHandler "github.com/adeben33/vehicleParkingApi/pkg/handler/reservation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Reservation(r *gin.Engine, validate *validator.Validate, ApiVersion string) *gin.Engine {
	Reservation := reservationHandler.Controller{Validate: validate}
	ReservationUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		{
			ReservationUrl.POST("/parkingSpace", Reservation.CreateReservation)
			ReservationUrl.GET("/parkingSpace/:parkingSpaceId", Reservation.GetParkingSpacebyId)
			ReservationUrl.PUT("/parkingSpace/:parkingSpaceId", Reservation.UpdateParkingSpace)
			ReservationUrl.DELETE("/parkingSpace/:parkingSpaceId", Reservation.DeleteParkingSpace)
			ReservationUrl.GET("/parkingSpace", Reservation.GetParkingSpaces)

		}
	}
	return r
}
