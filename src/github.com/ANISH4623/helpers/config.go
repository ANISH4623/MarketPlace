package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	SECRET_KEY  string
}

func LoadConfig(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
		return
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	AppConfig.PORT = port
	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_PORT = os.Getenv("DB_PORT")
	AppConfig.DB_USER = os.Getenv("DB_USER")
	AppConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.SECRET_KEY = os.Getenv("SECRET_KEY")

}
