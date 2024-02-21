package middleware

import (
	"fmt"
	"reflect"

	"github.com/labstack/echo/v4"
)

type Validator struct{}

func New() echo.Validator {
	return &Validator{}
}

func (v *Validator) Validate(i interface{}) error {
	r := reflect.ValueOf(i).Elem()
	for j := 0; j < r.NumField(); j++ {
		tag := r.Type().Field(j).Tag.Get("validate")
		name := r.Type().Field(j).Name
		length := r.Field(j).Interface()
		fmt.Println(name, tag, length)
	}
	return nil
}
