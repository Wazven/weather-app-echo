package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	WeatherAPIKey string
	ServerPort    int
}

func LoadConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return AppConfig{
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
		ServerPort:    8080,
	}
}