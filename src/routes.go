package src

import (
	"go-boilerplate-v2/src/controllers"

	"github.com/labstack/echo/v4"
	godi "github.com/putraawali/go-di"
)

func NewRoutes(app *echo.Echo, di godi.Container) {
	ctrl := controllers.NewController(di)

	app.POST("/register", ctrl.User.Register)
	app.POST("/login", ctrl.User.Login)
}
