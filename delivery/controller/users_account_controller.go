package controller

import (
	"EEP/e-wallets/helper"
	"EEP/e-wallets/model/dto/req"
	"EEP/e-wallets/model/dto/resp"
	"EEP/e-wallets/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserAccountController struct {
	usersAccountUC usecase.UsersAccountUseCase
	walletsUC      usecase.WalletsUseCase
	engine         *echo.Echo
}

func NewUsersAccountController(usersAccountUC usecase.UsersAccountUseCase, walletUC usecase.WalletsUseCase, engine *echo.Echo) *UserAccountController {
	return &UserAccountController{
		usersAccountUC: usersAccountUC,
		walletsUC:      walletUC,
		engine:         engine,
	}
}

func (ua *UserAccountController) AuthRoute() {
	rg := ua.engine.Group("/api/v1")

	rg.POST("auth/register", ua.registerHandler)

	//ua.engine.HTTPErrorHandler = exception.ErrorHandling
}

func (ua *UserAccountController) registerHandler(c echo.Context) error {
	var payload req.RegisterRequest
	err := helper.ReadFromRequestBody(c, &payload)
	if err != nil {
		return err
	}

	err = ua.usersAccountUC.Register(payload)
	if err != nil {
		return err
	}

	response := resp.ApiResponse{
		Status:  http.StatusCreated,
		Message: "successfully register",
		Data:    nil,
	}

	return helper.WriteToResponseBody(c, response)

}
