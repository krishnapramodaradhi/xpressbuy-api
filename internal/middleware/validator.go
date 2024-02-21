package middleware

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

type Validator struct{}

func New() echo.Validator {
	return &Validator{}
}

func (v *Validator) Validate(data interface{}) error {
	r := reflect.ValueOf(data).Elem()
	for i := 0; i < r.NumField(); i++ {
		tag := r.Type().Field(i).Tag.Get("validate")
		if tag == "" || tag == "-" {
			continue
		}
		// name := r.Type().Field(i).Name
		// value := r.Field(i).Interface()
		// TODO: Create a  struct to store the rules and validate
		/* for _, rule := range strings.Split(tag, ",") {
			switch rule {
			case "required":
				if r.Field(i).IsZero() {
					return fmt.Errorf("%v is required", name)
				}
			}
		} */
	}
	return nil
}
