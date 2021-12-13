package util

import (
	"fmt"

	"github.com/labstack/echo/v4"
	grpcCode "google.golang.org/grpc/codes"
)

// Response struct
type Response struct {
	Code       grpcCode.Code          `json:"code"`
	Message    string                 `json:"message,omitempty"`
	Data       interface{}            `json:"data,omitempty"`
	Pagination *Pagination            `json:"pagination,omitempty"`
	Errors     []string               `json:"errors,omitempty"`
	Header     map[string]interface{} `json:"-"`
}

// WithPagination set response with pagination
func (r *Response) WithPagination(c echo.Context, pagination *Pagination) *Response {
	r.Pagination = pagination
	return r
}

// JSON render response as JSON
func (r *Response) JSON(c echo.Context) error {
	for k, v := range r.Header {
		fmt.Println(c.Response().Header().Get(k), v, k)
		c.Response().Header().Set(k, fmt.Sprintf("%v,%v", c.Response().Header().Get(k), v))
	}
	return c.JSON(HTTPStatusFromCode(r.Code), r)
}
