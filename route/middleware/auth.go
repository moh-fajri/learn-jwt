package middleware

import (
	"os"
	"strings"

	"github.com/moh-fajri/learn-jwt/auth"
	"github.com/moh-fajri/learn-jwt/util"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return status.Errorf(util.Forbidden, "No Authorization header provided")
		}
		extractedToken := strings.Split(token, "Bearer ")
		if len(extractedToken) == 2 {
			token = strings.TrimSpace(extractedToken[1])
		} else {
			return status.Errorf(util.InvalidArgument, "Incorrect Format of Authorization Token")
		}

		jwtWrapper := auth.JwtWrapper{
			SecretKey: os.Getenv("JWT_SECRET"),
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(token)
		if err != nil {
			return status.Errorf(util.Unauthorized, "Token Not Found or Expired")
		}

		c.Set("email", claims.Email)

		return next(c)
	}
}
