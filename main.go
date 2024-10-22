package main

import (
	"go-boilerplate-v2/src"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load(".env")

	app := echo.New()

	module := src.Module{}

	module.New(app)

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	if port == 0 {
		port = 8000
	}

	app.Logger.Fatal(app.Start("localhost:" + strconv.Itoa(port)))
}
