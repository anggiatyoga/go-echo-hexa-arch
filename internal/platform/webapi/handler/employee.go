package handler

import (
	"net/http"

	"github.com/anggiatyoga/hris-api/internal/platform"
	"github.com/anggiatyoga/hris-api/internal/platform/webapi"
	"github.com/labstack/echo/v4"
)

const (
	coba = 3
)

func CreateHandler(modules platform.AppModule) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, webapi.Response{
			Status:  webapi.Success,
			Code:    coba,
			Message: "Ini create handler",
		})
	}
}

func ReadHandler(modules platform.AppModule) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, webapi.Response{
			Status:  webapi.Success,
			Code:    coba,
			Message: "Ini Read handler",
		})
	}
}

func UpdateHandler(modules platform.AppModule) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, webapi.Response{
			Status:  webapi.Success,
			Code:    coba,
			Message: "Ini Read handler",
		})
	}
}

func DeleteHandler(modules platform.AppModule) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, webapi.Response{
			Status:  webapi.Success,
			Code:    coba,
			Message: "Ini Read handler",
		})
	}
}
