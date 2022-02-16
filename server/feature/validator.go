package feature

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Validate is the main validator of the feature.
var Validate *validator.Validate = newValidator()

func newValidator() *validator.Validate {
	validate := validator.New()
	err := validate.RegisterValidation("custom-validation", customValidation)
	if err != nil {
		panic(err)
	}
	return validate
}

func customValidation(field validator.FieldLevel) bool {
	matched, err := regexp.MatchString("^[a-zA-Z0-9,:!?.\\s]+$", field.Field().String())
	if err != nil {
		return false
	}
	return matched
}
