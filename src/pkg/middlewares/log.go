package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat:   time.RFC3339Nano,
		DisableHTMLEscape: true,
	}

	logger.Level = logrus.TraceLevel
	logger.SetOutput(os.Stdout)

	return logger
}

func LogRequest(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			header := req.Header

			requestID := header.Get("request-id")
			if requestID == "" {
				id := uuid.New()
				requestID = id.String()
				c.Request().Header.Set("request-id", requestID)
			}

			reqBody, _ := io.ReadAll(req.Body)
			c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody))

			logger.WithFields(logrus.Fields{
				"type":       "request",
				"request_id": requestID,
				"host":       req.Host,
				"method":     req.Method,
				"url":        req.URL.String(),
				"header":     generateHeader(header),
			}).Trace(string(reqBody))

			return next(c)
		}
	}
}

func LogResponse(logger *logrus.Logger) echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			status := c.Response().Status

			req := c.Request()
			header := req.Header

			// Try getting request id from context
			reqID, _ := req.Context().Value("request-id").(string)

			// Try getting request id from header
			if reqID == "" {
				reqID = header.Get("request-id")
			}

			// Try getting request id from struct property
			// Note: this is most probably not valid due to race condition, but just so request id on log message has value, we'll getting from this struct property

			logger.WithFields(logrus.Fields{
				"type":       "response",
				"request_id": reqID,
				"host":       req.Host,
				"method":     req.Method,
				"url":        req.URL.String(),
				"header":     generateHeader(header),
				"status":     status,
			}).Trace(string(resBody))
		},
	})
}

func generateHeader(header http.Header) string {
	var headers []string
	for k, v := range header {
		headers = append(headers, k+":"+strings.Join(v, ","))
	}

	return strings.Join(headers, "|")
}
