package product

import (
	"context"
	"net/http"

	"github.com/moh-fajri/learn-jwt/repo"
	"github.com/moh-fajri/learn-jwt/util"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

// Delete is
type Delete struct{}

// NewDelete handler
func NewDelete() *Delete {
	return &Delete{}
}

func (h *Delete) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var product repo.Product
	product.ID = uint(util.StringToInt(c.Param("id")))
	if err = product.Delete(ctx); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	// response
	resp := util.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
	}
	return resp.JSON(c)
}
