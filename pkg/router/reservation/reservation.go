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
			ReservationUrl.POST("/reservation", Reservation.CreateReservation)
			ReservationUrl.GET("/reservation/:reservationId", Reservation.GetReservationById)
			ReservationUrl.PUT("/reservation/:reservationId", Reservation.UpdateReservation)
			ReservationUrl.DELETE("/reservation/:reservationId", Reservation.DeleteReservation)
			ReservationUrl.GET("/reservation", Reservation.GetReservations)

		}
	}
	return r
}
