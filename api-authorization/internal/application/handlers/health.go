package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthStatus struct {
	Status string
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthStatus{
		Status: "Healthy",
	})
}
