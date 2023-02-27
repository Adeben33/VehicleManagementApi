package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	Mongodb MongodbConfiguration
	Server  ServerConfiguration
}

//Configuration variable to call when needed

var (
	Config *Configuration
)

func Setup() {
	var configuration *Configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("Unable to decode into struct, %v", err)
	}
	Config = configuration
	fmt.Print("configurations loading successfully")
}

func GetConfig() *Configuration {
	return Config
}
