package config

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
    Value       string
}


func ValidateStruct(entity interface{}) []string {
    var errors []string
    validate := validator.New()
    err := validate.Struct(entity)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            errors = append(errors,err.Error())
        }
    }
    return errors
}