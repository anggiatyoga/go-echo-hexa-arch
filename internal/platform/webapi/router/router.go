package router

import (
	"errors"
	"fmt"

	"github.com/anggiatyoga/hris-api/internal/platform"
	"github.com/anggiatyoga/hris-api/internal/platform/webapi"
	"github.com/anggiatyoga/hris-api/internal/platform/webapi/handler"
	"github.com/labstack/echo/v4"
)

func Run(modules platform.AppModule) (*echo.Echo, error) {
	e, err := webapi.NewAPI()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error creating new api: ", err.Error()))
	}

	employee(e, modules)

	return e, nil
}

func employee(e *echo.Echo, modules platform.AppModule) {
	employeeGroup := e.Group("/employee")

	employeeGroup.POST("/create", handler.CreateHandler(modules))
	employeeGroup.GET("/read", handler.ReadHandler(modules))
	employeeGroup.POST("/update", handler.UpdateHandler(modules))
	employeeGroup.POST("/delete", handler.DeleteHandler(modules))
}
