package webapi

import "github.com/labstack/echo/v4"

func NewAPI() (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true

	return e, nil
}
