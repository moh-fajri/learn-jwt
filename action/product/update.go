package product

import (
	"context"
	"net/http"

	"github.com/moh-fajri/learn-jwt/repo"
	"github.com/moh-fajri/learn-jwt/util"

	"github.com/jinzhu/copier"

	"google.golang.org/grpc/status"

	"github.com/labstack/echo/v4"
)

// Update is
type Update struct{}

// NewUpdate handler
func NewUpdate() *Update {
	return &Update{}
}

func (h *Update) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	// get request
	req := new(Request)
	err = h.validate(req, c)
	if err != nil {
		return
	}
	// convert struct
	var product repo.Product
	if err = copier.Copy(&product, req); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	product.ID = uint(util.StringToInt(c.Param("id")))
	// update product
	if err = product.Update(ctx); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	// response
	resp := util.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
	}
	return resp.JSON(c)
}

func (h *Update) validate(r *Request, c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	return c.Validate(r)
}
