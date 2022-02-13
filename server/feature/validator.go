package feature

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func initValidator() {
	validate = validator.New()
	validate.RegisterValidation("custom-validation", customValidation)
}

func customValidation(field validator.FieldLevel) bool {
	matched, err := regexp.MatchString("^[a-zA-Z0-9]+$", field.Field().String())
	if err != nil {
		return false
	}
	return matched
}
