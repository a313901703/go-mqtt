package controller

import (
	"errors"
	"fmt"
	"mqtt/api"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BaseController struct {
}

var validate *validator.Validate

func (cv BaseController) Validate(i interface{}) error {
	validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(i); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return errors.New("validation error")
		}
		errStr := "params error"
		for _, err := range err.(validator.ValidationErrors) {
			errStr += ",error field: " + err.Field()
			break
		}
		// Optionally, you could return the error to give each route more control over the status code
		return errors.New(errStr)
	}
	return nil
}

func RespSuccess(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, &api.Resp{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func RespError(c echo.Context, code int, message string) error {
	return c.JSON(http.StatusOK, &api.RespErr{
		Code:    code,
		Message: message,
	})
}
