package main

import (
	"fmt"
	"github.com/adeben33/vehicleParkingApi/internal/config"
	"github.com/adeben33/vehicleParkingApi/internal/database"
	"github.com/adeben33/vehicleParkingApi/pkg/router"
	"github.com/go-playground/validator/v10"
	"log"
)

func init() {
	config.Setup()
	database.ConnectToDb()
}

func main() {

	config := config.GetConfig()
	validatorRef := validator.New()
	r := router.Setup(validatorRef)
	fmt.Printf("Server is starting at 127.0.0.1:%s", config.Server.Port)
	log.Fatal(r.Run(config.Server.Port))

}
