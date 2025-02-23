package controllers

import (
	"go-boilerplate-v2/src/constants"
	"go-boilerplate-v2/src/dtos"
	"go-boilerplate-v2/src/pkg/response"
	"go-boilerplate-v2/src/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
	godi "github.com/putraawali/go-di"
)

type UserController interface {
	Register(echo.Context) error
	Login(echo.Context) error
}

type userController struct {
	uc       *usecases.Usecases
	response *response.Response
}

func NewUserController(di godi.Container) UserController {
	return &userController{
		uc:       di.Get(constants.USECASE).(*usecases.Usecases),
		response: di.Get(constants.RESPONSE).(*response.Response),
	}
}

func (u *userController) Register(c echo.Context) error {
	param := dtos.RegisterParam{}

	err := c.Bind(&param)
	if err != nil {
		err = u.response.NewError().
			SetContext(c.Request().Context()).
			SetDetail(err.Error()).
			SetMessage(err).
			SetStatusCode(http.StatusBadRequest)

		return c.JSON(u.response.Send(0, nil, err))
	}

	err = param.Validate()
	if err != nil {
		err = u.response.NewError().
			SetContext(c.Request().Context()).
			SetDetail(err.Error()).
			SetMessage(err).
			SetStatusCode(http.StatusBadRequest)

		return c.JSON(u.response.Send(0, nil, err))
	}

	err = u.uc.User.Register(c.Request().Context(), param)
	if err != nil {
		return c.JSON(u.response.Send(0, nil, err))
	}

	return c.JSON(u.response.Send(http.StatusCreated, nil, nil))
}

func (u *userController) Login(c echo.Context) error {
	param := dtos.LoginParam{}

	err := c.Bind(&param)
	if err != nil {
		err = u.response.NewError().
			SetContext(c.Request().Context()).
			SetDetail(err.Error()).
			SetMessage(err).
			SetStatusCode(http.StatusBadRequest)

		return c.JSON(u.response.Send(0, nil, err))
	}

	err = param.Validate()
	if err != nil {
		err = u.response.NewError().
			SetContext(c.Request().Context()).
			SetDetail(err.Error()).
			SetMessage(err).
			SetStatusCode(http.StatusBadRequest)

		return c.JSON(u.response.Send(0, nil, err))
	}

	data, err := u.uc.User.Login(c.Request().Context(), param)
	if err != nil {
		return c.JSON(u.response.Send(0, nil, err))
	}

	return c.JSON(u.response.Send(http.StatusOK, data, nil))
}
