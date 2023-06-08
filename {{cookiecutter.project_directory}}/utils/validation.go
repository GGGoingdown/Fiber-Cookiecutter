package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorStructResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func (e *ErrorStructResponse) Error() string {
	return fmt.Sprintf("FailedField: %s, Tag: %s, Value: %s", e.FailedField, e.Tag, e.Value)
}

func StructValidator(stu interface{}) []*ErrorStructResponse {
	var errors []*ErrorStructResponse
	err := validate.Struct(stu)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorStructResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
