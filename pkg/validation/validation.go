package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	customvalidator "link-shortener/pkg/validation/custom-validator"
)

type Validator struct {
	*validator.Validate
}

var validatorInstance *Validator

func getInstance() *Validator {
	if validatorInstance == nil {
		validatorInstance = create()
	}

	return validatorInstance
}

func create() *Validator {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.RegisterValidation("yyyymmdd", customvalidator.Yyyymmdd); err != nil {
		panic(fmt.Sprint("RegisterValidation Validation error", err))
	}

	return &Validator{validate}
}

func IsValidStruct[T any](payload *T) error {
	validate := getInstance()
	err := validate.Struct(payload)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			//for _, validationError := range validationErrors {
			//	log.Error("validationError: ", validationError)
			//}
			return validationErrors
		}
		return err
	}
	return nil
}
