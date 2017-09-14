package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func CheckHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
