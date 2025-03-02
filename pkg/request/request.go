package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func GetBody[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	body, err := jsonDecode[T](r.Body)
	if err != nil {
		return *new(T), err
	}

	if err := isValid[T](body); err != nil {
		return *new(T), err
	}

	return body, nil
}

func jsonDecode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}

	return payload, nil
}

func isValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		// Type assert the error to validator.ValidationErrors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, validationError := range validationErrors {
				fmt.Println("validationError: ", validationError)
			}
			return validationErrors
		}
		return err
	}
	return nil
}
