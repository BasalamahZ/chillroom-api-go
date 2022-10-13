package helpers

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func validationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", e.Field(), e.Param())
	case "email":
		return fmt.Sprintf("Invalid %s format", e.Field())
	case "eqfield":
		return fmt.Sprintf("%s must be match", e.Field())
	}
	return fmt.Sprintf("%s is not valid", e.Field())
}

func FormatError(err error) []string {
	var errorr []string
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range err.(validator.ValidationErrors) {
			errorr = append(errorr, validationErrorToText(e))
		}
	}
	return errorr
}
