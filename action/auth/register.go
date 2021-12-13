package auth

import (
	"context"
	"net/http"

	"github.com/moh-fajri/learn-jwt/repo"
	"github.com/moh-fajri/learn-jwt/util"

	"google.golang.org/grpc/status"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// Register is
type Register struct{}

// NewRegister handler
func NewRegister() *Register {
	return &Register{}
}

func (r *Register) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	// get request
	req := new(RegisterRequest)
	err = r.validate(req, c)
	if err != nil {
		return
	}
	var user repo.User
	if err = user.HashPassword(req.Password); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}

	user.Email = req.Email
	if err = user.Create(ctx); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	resp := util.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
	}
	return resp.JSON(c)
}

func (r *Register) validate(req *RegisterRequest, c echo.Context) error {
	if err := c.Bind(req); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	return c.Validate(req)
}
