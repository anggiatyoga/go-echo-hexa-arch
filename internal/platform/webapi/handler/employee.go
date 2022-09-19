package handler

import (
	"fmt"
	"net/http"

	"github.com/anggiatyoga/hris-api/internal/domain/employee"
	"github.com/anggiatyoga/hris-api/internal/platform"
	"github.com/anggiatyoga/hris-api/internal/platform/webapi"
	"github.com/anggiatyoga/hris-api/internal/platform/webapi/request"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateHandler(modules platform.AppModule) echo.HandlerFunc {
	return func(c echo.Context) error {
		var r request.EmployeeParam
		err := c.Bind(&r)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, webapi.Response{
				Status:  webapi.Error,
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		err = modules.EmployeeModules.Create(r.ToEntity())
		if err != nil {
			fmt.Printf("Create Handler error: %s", err.Error())
			return c.JSON(http.StatusInternalServerError, webapi.Response{
				Status:  webapi.Error,
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, webapi.Response{
			Status:  webapi.Success,
			Code:    http.StatusOK,
			Message: "Success create",
		})
	}
}

func ReadHandler(modules platform.AppModule) echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := modules.EmployeeModules.Read()
		if err != nil {
			fmt.Printf("Read Handler error: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, webapi.Response{
				Status:  webapi.Error,
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if len(result) == 0 {
			result = make([]employee.Employee, 0)
		}

		return c.JSON(http.StatusOK, webapi.Response{
			Status:  webapi.Success,
			Code:    http.StatusOK,
			Data:    result,
			Message: "Ini Read handler",
		})
	}
}

func UpdateHandler(modules platform.AppModule) echo.HandlerFunc {
	return func(c echo.Context) error {
		var r request.EmployeeParam
		err := c.Bind(&r)
		fmt.Printf("id: %v", r.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, webapi.Response{
				Status:  webapi.Error,
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		v := validator.New()

		if err := v.Struct(r); err != nil || r.ID == 0 {
			fmt.Printf("Update Handler error: %s\n", err.Error())
			return c.JSON(http.StatusMethodNotAllowed, webapi.Response{
				Status:  webapi.Fail,
				Code:    http.StatusMethodNotAllowed,
				Message: "please fill requirement",
			})
		}

		// if r.ID == 0 {
		// 	fmt.Printf("Update Handler error: %s\n", err.Error())
		// 	return c.JSON(http.StatusMethodNotAllowed, webapi.Response{
		// 		Status:  webapi.Fail,
		// 		Code:    http.StatusMethodNotAllowed,
		// 		Message: "please fill requirement",
		// 	})
		// }

		err = modules.EmployeeModules.Update(r.ToEntity())
		if err != nil {
			fmt.Printf("Update Handler error: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, webapi.Response{
				Status:  webapi.Error,
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, webapi.Response{
			Status:  webapi.Success,
			Code:    http.StatusOK,
			Message: fmt.Sprintf("Success update id: %v\n", r.ID),
		})
	}
}

func DeleteHandler(modules platform.AppModule) echo.HandlerFunc {
	return func(c echo.Context) error {
		var r request.EmployeeParam
		err := c.Bind(&r)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, webapi.Response{
				Status:  webapi.Error,
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		v := validator.New()

		if err := v.Struct(r); err != nil {
			fmt.Printf("Delete Handler error: %s\n", err.Error())
			return c.JSON(http.StatusMethodNotAllowed, webapi.Response{
				Status:  webapi.Fail,
				Code:    http.StatusMethodNotAllowed,
				Message: "please fill requirement",
			})
		}

		err = modules.EmployeeModules.Delete(r.ToEntity())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, webapi.Response{
				Status:  webapi.Error,
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, webapi.Response{
			Status:  webapi.Success,
			Code:    http.StatusOK,
			Message: fmt.Sprintf("Success delete id: ", r.ID),
		})
	}
}
