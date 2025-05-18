package customvalidator

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func Yyyymmdd(fl validator.FieldLevel) bool  {
	_, err := time.Parse(time.DateOnly, fl.Field().String())
	return err == nil
}
