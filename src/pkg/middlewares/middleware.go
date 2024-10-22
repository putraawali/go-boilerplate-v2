package middlewares

import (
	"go-boilerplate-v2/src/pkg/jwt"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sarulabs/di"
)

func UseMiddlwares(app *echo.Echo, di di.Container) {
	app.Use(middleware.Recover())
	app.Use(ValidateRequestID())

	logger := NewLogger()
	app.Use(LogRequest(logger))
	app.Use(LogResponse(logger))

	app.Use(echojwt.WithConfig(jwt.ConfigJwt()))
}
