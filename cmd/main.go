package main

import (
	"fmt"

	"github.com/wazven/weather-app-echo/config"
)


func main() {
	appConfig := config.LoadConfig()

	fmt.Printf("Weather API Key: %s\n", appConfig.WeatherAPIKey)
	fmt.Printf("Server Port: %d\n", appConfig.ServerPort)
}