package product

import (
	"context"
	"net/http"

	"github.com/moh-fajri/learn-jwt/repo"
	"github.com/moh-fajri/learn-jwt/util"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

type Request struct {
	SKU         string  `json:"sku" form:"sku" validate:"required"`
	ProductName string  `json:"product_name" form:"product_name" validate:"required"`
	Qty         float64 `json:"qty" form:"qty" validate:"required"`
	Price       float64 `json:"price" form:"price" validate:"required"`
	Unit        string  `json:"unit" form:"unit" validate:"required"`
	Status      int32   `json:"status" form:"status" validate:"required"`
}

// Create is
type Create struct{}

// NewCreate handler
func NewCreate() *Create {
	return &Create{}
}

func (h *Create) Handle(c echo.Context) (err error) {
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
	var product repo.Product
	if err = copier.Copy(&product, req); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	//create product
	if err = product.Create(ctx); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	// response
	resp := util.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
	}
	return resp.JSON(c)
}

func (h *Create) validate(r *Request, c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	return c.Validate(r)
}
