package validators

import (
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

func ValidateGolangTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Golang")
}
