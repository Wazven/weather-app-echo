package weather

import (
	"net/http"

	"github.com/labstack/echo"
)

func SendJSONResponse(c echo.Context, response interface{}) error {
	return c.JSON(http.StatusOK, response)
}

func SendErrorResponse(c echo.Context, statusCode int, message string) error {
	errorResponse := map[string]interface{}{
		"error": message,
	}
	return c.JSON(statusCode, errorResponse)
}