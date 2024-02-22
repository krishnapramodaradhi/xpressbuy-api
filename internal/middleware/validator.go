package middleware

import (
	"errors"
	"fmt"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"

	"github.com/krishnapramodaradhi/xpressbuy-api/internal/util/constants"
	"github.com/labstack/echo/v4"
)

var (
	min             = regexp.MustCompile(constants.MIN)
	max             = regexp.MustCompile(constants.MAX)
	ErrInvalidRegex = errors.New("invalid regex")
)

type Validator struct {
	tag   string
	name  string
	value any
	field reflect.Value
}

func New() echo.Validator {
	return &Validator{}
}

func (v *Validator) Validate(data interface{}) error {
	r := reflect.ValueOf(data).Elem()
	for i := 0; i < r.NumField(); i++ {
		v.tag = r.Type().Field(i).Tag.Get("validate")
		if v.tag == "" || v.tag == "-" {
			continue
		}
		v.name = r.Type().Field(i).Name
		v.value = r.Field(i).Interface()
		v.field = r.Field(i)

	}
	return nil
}

func (v *Validator) required() error {
	if v.field.IsZero() {
		return fmt.Errorf("%v is required", v.name)
	}
	return nil
}

func (v *Validator) min() error {
	match := min.FindStringSubmatch(v.tag)
	minVal, _ := strconv.Atoi(match[1])
	if v.field.Len() < minVal {
		return fmt.Errorf("%v should be minimum %d characters", v.name, minVal)
	}
	return nil
}

func (v *Validator) max() error {
	match := max.FindStringSubmatch(v.tag)
	maxVal, _ := strconv.Atoi(match[1])
	if v.field.Len() > maxVal {
		return fmt.Errorf("%v should be maximum %d characters", v.name, maxVal)
	}
	return nil
}

func (v *Validator) email() error {
	if _, err := mail.ParseAddress(v.value.(string)); err != nil {
		return fmt.Errorf("not a valid email: %v", v.value)
	}
	return nil
}

func (v *Validator) regex(regExp string) error {
	r, err := regexp.Compile(regExp)
	if err != nil {
		return ErrInvalidRegex
	}
	match := r.FindStringSubmatch(v.value.(string))
	if len(match) == 0 {
		return fmt.Errorf("%v doesn't match with the regex: %v", v.name, regExp)
	}
	return nil
}
