package middlewares

import (
	"go-boilerplate-v2/src/pkg/jwt"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	godi "github.com/putraawali/go-di"
)

func UseMiddlwares(app *echo.Echo, di godi.Container) {
	app.Use(middleware.Recover())
	app.Use(ValidateRequestID())

	logger := NewLogger()
	app.Use(LogRequest(logger))
	app.Use(LogResponse(logger))

	app.Use(echojwt.WithConfig(jwt.ConfigJwt()))
}
