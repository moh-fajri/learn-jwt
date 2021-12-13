package product

import (
	"context"
	"net/http"

	"github.com/moh-fajri/learn-jwt/repo"
	"github.com/moh-fajri/learn-jwt/util"

	"google.golang.org/grpc/status"

	"github.com/labstack/echo/v4"
)

// List is
type List struct{}

// NewList handler
func NewList() *List {
	return &List{}
}

func (h *List) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	count := util.StringToInt(c.QueryParam("count"))
	if count == 0 {
		count = 15
	}
	page := util.StringToInt(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	sort := c.QueryParam("sort")
	if sort == "" {
		sort = "created_at desc"
	}
	// get products
	var product repo.Product
	product.SKU = c.QueryParam("sku")
	pagination := util.NewPagination(int32(page), int32(count))
	products, pg, err := product.ListWithPagination(ctx, pagination, sort)
	if err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}

	// response
	resp := &util.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: map[string]interface{}{
			"products": products,
		},
	}
	resp = resp.WithPagination(c, pg)
	return resp.JSON(c)
}
