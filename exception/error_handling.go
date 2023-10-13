package exception

import (
	"EEP/e-wallets/helper"
	"EEP/e-wallets/model/dto/resp"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorHandling(err error, c echo.Context) {
	if ErrNotFound(err, c) {
		return
	}

	if ErrValidation(err, c) {
		return
	}

	ErrInternalServer(err, c)
}

func ErrNotFound(err error, c echo.Context) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		c.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

		apiResponse := resp.ApiResponse{
			Status:  http.StatusNotFound,
			Message: "not found",
			Data:    exception.Error(),
		}
		c.Response().WriteHeader(apiResponse.Status)
		err := helper.WriteToResponseBody(c, apiResponse)
		if err != nil {
			return false
		}
		return true
	} else {
		return false
	}
}

func ErrValidation(err error, c echo.Context) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		c.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

		apiResponse := resp.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Data:    exception.Error(),
		}
		c.Response().WriteHeader(apiResponse.Status)
		err := helper.WriteToResponseBody(c, apiResponse)
		if err != nil {
			return false
		}
		return true
	} else {
		return false
	}
}

func ErrInternalServer(err error, c echo.Context) {
	c.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	apiResponse := resp.ApiResponse{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
		Data:    err.Error(),
	}
	c.Response().WriteHeader(apiResponse.Status)
	err = helper.WriteToResponseBody(c, apiResponse)
	if err != nil {
		return
	}
}
