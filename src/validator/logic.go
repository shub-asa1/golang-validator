package validator

import (
	"reflect"
)

type field interface{}

//Func sdsdsd
type Func func(f field) bool

var bakedValidator = map[string]Func{
	"email": isEmail,
}

func isEmail(f field) bool {
	return emailRegex.MatchString(reflect.ValueOf(f).String())
}
