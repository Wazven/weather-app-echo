package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/wazven/weather-app-echo/api/weather"
	"github.com/wazven/weather-app-echo/config"
)

var appConfig config.AppConfig

func main() {
	appConfig = config.LoadConfig()

	e := echo.New()

	e.GET("/weather", func(c echo.Context) error {
		return weather.WeatherForecastHandler(c, appConfig)
	})

	fmt.Printf("Server is running on :%d...\n", appConfig.ServerPort)
	e.Start(fmt.Sprintf(":%d", appConfig.ServerPort))
}