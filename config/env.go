package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var env ENV

func GetEnvString(key string) string {
	return os.Getenv(key)
}

func GetEnv() *ENV {
	return &env
}

func InitDotEnv() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort := GetEnvString("SERVER_PORT")
	database := Database{URI: GetEnvString("MONGODB_URL"), Name: GetEnvString("MONGODB_DATABASE")}

	env = ENV{
		AppPort:  appPort,
		Database: database,
	}
}
