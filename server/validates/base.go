package validates

import (
	// echo
	"net/http"
	"github.com/labstack/echo/v4"

	// バリデーション
	"github.com/go-playground/validator/v10"
)

type Validator struct {
    validator *validator.Validate
}

func SetValidator() echo.Validator {
    return &Validator{validator: validator.New()}
}

// バリデーションエラーはこれで設定(ステータスコードも)
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
	  // Optionally, you could return the error to give each route more control over the status code
	  return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return nil
  }
