package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var whiteListPaths = []string{
	"/api/auth/login",
	"/api/auth/register",
}

func WebSecurityConfig(e *echo.Echo) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return jwt.MapClaims{
				"expires_at": jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			}
		},
		SigningKey: []byte("secret"),
		Skipper:    skipAuth,
	}
	e.Use(echojwt.WithConfig(config))
}

func skipAuth(e echo.Context) bool {
	path := e.Path()
	for _, p := range whiteListPaths {
		if path == p {
			return true
		}
	}
	return false
}
