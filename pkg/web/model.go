package web

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Response[T any] struct {
	Code  int    `json:"code"`
	Data  T      `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// Person
// Demo struct for form binding
type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

// SUri
// Demo struct for URI binding
type SUri struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

// Booking  contains binded and validated data
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfieId=CheckIn,bookabledate" time_format:"2006-01-02"`
}

var BookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}
