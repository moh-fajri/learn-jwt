package route

import (
	"github.com/moh-fajri/learn-jwt/action/auth"
	"github.com/moh-fajri/learn-jwt/action/product"

	"github.com/labstack/echo/v4"
)

// Handler endpoint to use it later
type Handler interface {
	Handle(c echo.Context) (err error)
}

var endpoint = map[string]Handler{
	//auth
	"login":    auth.NewLogin(),
	"register": auth.NewRegister(),
	//product
	"create": product.NewCreate(),
	"list":   product.NewList(),
	"delete": product.NewDelete(),
	"update": product.NewUpdate(),
}
