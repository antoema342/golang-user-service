package helpers

import "github.com/go-playground/validator/v10"

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Please provide a valid email address"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "number":
		return "Should be number " + fe.Param()
	case "min":
		return "Should be at least " + fe.Param() + " characters long"
	}

	return "Unknown error"
}

func GetErrorGormMsg(err string) string {
	switch err {
	case "unique_violation":
		return "already exists"
	}

	return "Unknown error"
}
