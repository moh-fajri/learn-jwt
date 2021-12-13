package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/moh-fajri/learn-jwt/auth"
	"github.com/moh-fajri/learn-jwt/repo"
	"github.com/moh-fajri/learn-jwt/util"

	"google.golang.org/grpc/status"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// Login is
type Login struct{}

// NewLogin handler
func NewLogin() *Login {
	return &Login{}
}

func (l *Login) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	// get request
	req := new(LoginRequest)
	err = l.validate(req, c)
	if err != nil {
		return
	}
	// check user with email
	var user repo.User
	err = user.GetWithEmail(ctx, req.Email)
	if err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	// check user with password
	err = user.CheckPassword(req.Password)
	if err != nil {
		return status.Errorf(util.Unauthorized, err.Error())
	}
	// generate token jwt
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       os.Getenv("JWT_SECRET"),
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		return status.Errorf(util.Unauthorized, err.Error())
	}
	// response
	resp := util.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data: map[string]interface{}{
			"token": signedToken,
			"user": map[string]interface{}{
				"id":         user.ID,
				"email":      user.Email,
				"created_at": user.CreatedAt,
			},
		},
	}
	return resp.JSON(c)
}

func (l *Login) validate(r *LoginRequest, c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return status.Errorf(util.InvalidArgument, err.Error())
	}
	return c.Validate(r)
}
