package jwt

import (
	"errors"
	"go-boilerplate-v2/src/pkg/response"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var secretKey = os.Getenv("JWT_SECRET")

type JwtCustomClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"phone"`
	jwt.RegisteredClaims
}

func GenerateToken(id int64, email string) (accessToken string) {
	claims := &JwtCustomClaims{
		UserID: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, _ = parseToken.SignedString([]byte(secretKey))

	return
}

func ConfigJwt() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
		Skipper: func(c echo.Context) bool {
			url := c.Request().URL.Path
			if url == "/register" || url == "/login" {
				return true
			}

			return false
		},
		ContextKey: "user-data",
		ErrorHandler: func(c echo.Context, _ error) error {
			resp := response.NewResponse()
			err := resp.NewError().
				SetContext(c.Request().Context()).
				SetDetail("Unauthorized").
				SetMessage(errors.New("please login first")).
				SetStatusCode(http.StatusUnauthorized)

			return c.JSON(resp.Send(0, nil, err))
		},
	}
}

func GetUserData(c echo.Context) *JwtCustomClaims {
	userClaim := c.Get("user-data").(*jwt.Token)
	user := userClaim.Claims.(*JwtCustomClaims)

	return user
}
