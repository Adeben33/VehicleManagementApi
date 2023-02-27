package router

import (
	"github.com/adeben33/vehicleParkingApi/pkg/router/admin"
	"github.com/adeben33/vehicleParkingApi/pkg/router/parkingSpace"
	"github.com/adeben33/vehicleParkingApi/pkg/router/payment"
	"github.com/adeben33/vehicleParkingApi/pkg/router/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Setup(validate *validator.Validate) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	ApiVersion := "v1"

	user.User(r, validate, ApiVersion)
	admin.Admin(r, validate, ApiVersion)
	payment.Payment(r, validate, ApiVersion)
	parkingSpace.Parking(r, validate, ApiVersion)

	return r
}
