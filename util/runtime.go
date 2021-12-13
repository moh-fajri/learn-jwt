package util

import (
	"net/http"

	"github.com/go-playground/locales/id"

	ut "github.com/go-playground/universal-translator"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/validator.v9"
	enTrans "gopkg.in/go-playground/validator.v9/translations/en"
)

var translator ut.Translator

func CustomErrorHandler(c *echo.Echo) {
	setValidator(c)
	c.HTTPErrorHandler = func(err error, c echo.Context) {
		// Validation Error
		if errs, ok := err.(validator.ValidationErrors); ok {
			var message []string

			translated := errs.Translate(translator)
			for _, v := range translated {
				message = append(message, v)
			}
			// call localization RPC
			resp := Response{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Errors:  message,
			}
			_ = resp.JSON(c)
			return
		}

		// gRPC Error
		if st, ok := status.FromError(err); ok {
			resp := Response{
				Code:    st.Code(),
				Message: http.StatusText(int(st.Code())),
				Errors:  []string{st.Message()},
			}

			_ = resp.JSON(c)
			return
		}
		resp := Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Errors:  []string{err.Error()},
		}
		_ = resp.JSON(c)
		return
	}
}

func setValidator(e *echo.Echo) {
	i := id.New()
	uni := ut.New(i, i)

	translator, _ = uni.GetTranslator("en")
	validate := validator.New()
	_ = enTrans.RegisterDefaultTranslations(validate, translator)
	e.Validator = &CustomValidator{Validator: validate}
}

// CustomValidator validation that handle validation
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
