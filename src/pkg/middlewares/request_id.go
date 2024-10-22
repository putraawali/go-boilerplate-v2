package middlewares

import (
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

func ValidateRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get("request-id")
			if requestID == "" {
				id := uuid.New()
				requestID = id.String()
				c.Request().Header.Set("request-id", requestID)
			}

			c.Set("request-id", requestID)

			return next(c)
		}
	}
}
