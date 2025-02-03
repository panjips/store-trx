package utils

import "github.com/go-playground/validator/v10"


type ValidationErrors struct {
	Errors map[string][]string `json:"errors"`
}

var validate = validator.New()

func ValidateRequest(data interface{}) *ValidationErrors{
	errs := validate.Struct(data)
	if errs == nil {
		return  nil
	}

	validationErrors := ValidationErrors{Errors: make(map[string][]string)}

	for _, err := range errs.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()
		validationErrors.Errors[field] = append(validationErrors.Errors[field], tag)
	}

	return &validationErrors
}