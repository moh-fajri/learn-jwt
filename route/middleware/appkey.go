package middleware

import (
	"net/http"
	"os"

	"github.com/moh-fajri/learn-jwt/util"

	"github.com/labstack/echo/v4"
)

func AuthApIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		appKey := c.Request().Header.Get("api-key")
		if appKey != os.Getenv("API_KEY") {
			resp := util.Response{
				Code:    http.StatusUnauthorized,
				Message: "API Key Not Valid",
			}
			return resp.JSON(c)
		}
		return next(c)
	}
}
