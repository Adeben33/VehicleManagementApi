package main


import (
	"fmt"
	"vehicleManagementApi/internal/config"
	"vehicleManagementApi/internal/database"
	"vehicleManagementApi/pkg/router"
	"log"
)

func init() {
	config.Setup()
	database.ConnectToDb()
}

func main() {
	getConfig := config.GetConfig()
	validatorRef := validator.New()
	r:= router.Setup(validatorRef)
	fmt.Printf("Server is starting at 127.0.0.1:%s", getConfig.Server.Port)
	log.Fatal(r.Run(getConfig.Server.Port)
}
