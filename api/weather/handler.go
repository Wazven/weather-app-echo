package weather

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wazven/weather-app-echo/config"
	"github.com/wazven/weather-app-echo/internal/weather"
)

var appConfig config.AppConfig
func WeatherForecastHandler(c echo.Context, appConfig config.AppConfig) error {
	location := c.QueryParam("location")
	if location == "" {
		return c.String(http.StatusBadRequest, "Location is required")
	}
	
	weatherData, err := weather.FetchWeatherData(appConfig.WeatherAPIKey, location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching weather data")
	}

	return c.JSON(http.StatusOK, weatherData)
}