package route

import (
	"github.com/moh-fajri/learn-jwt/route/middleware"

	"github.com/labstack/echo/v4"
)

var middlewareHandler = map[string]echo.MiddlewareFunc{
	"api_key": middleware.AuthApIKey,
	"auth":    middleware.Auth,
}
