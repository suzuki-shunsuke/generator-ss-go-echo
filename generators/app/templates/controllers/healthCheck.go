package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"<%= package %>/constants/logTypes"
)

func CheckHealth(c echo.Context) error {
	log.WithFields(log.Fields{
		"type": logTypes.Info,
	}).Info("Start CheckHealth")
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
