package util

import "gopkg.in/go-playground/validator.v9"

func Validate(s interface{}) (error) {
	validate := validator.New()
	return validate.Struct(s)
}
