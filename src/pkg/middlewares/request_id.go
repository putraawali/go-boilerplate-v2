package middlewares

import (
	"context"
	"go-boilerplate-v2/src/constants"

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

			ctx := context.WithValue(c.Request().Context(), constants.RequestID, requestID)

			r := c.Request().WithContext(ctx)
			c.SetRequest(r)

			return next(c)
		}
	}
}
