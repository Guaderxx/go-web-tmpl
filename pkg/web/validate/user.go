package validate

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	// password validator
	passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{4,12}$`)
)

var Password validator.Func = func(fl validator.FieldLevel) bool {
	if pwd, ok := fl.Field().Interface().(string); ok {
		return passwordRegex.MatchString(pwd)
	}
	return false
}
