package xex

import (
	"fmt"
	"math"
	"reflect"
)

func registerCoreBuiltins() {
	RegisterFunction(
		NewFunction(
			"equals",
			`compares 2 inputs returning a bool`,
			func(val1, val2 interface{}) bool {
				return val1 == val2
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"count",
			`returns the number of elements in the passed in slice / array or map`,
			func(val1 interface{}) (int, error) {
				switch reflect.TypeOf(val1).Kind() {
				case reflect.Array, reflect.Slice, reflect.Map:
					return reflect.ValueOf(val1).Len(), nil
				}
				return 0, fmt.Errorf("cannot count type %s", reflect.TypeOf(val1).String())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"switch",
			`Switches on the first value.
			The following values are equivalent to "case : result" pairs.
			If a final value is provided (an even number of arguments is passed), the final value is used as the default.
			If value1 equals value2, value3 is returned. Else if value1 equals value4, value5 is returned. And so on.
			If there is no default and no values matched, switch returns nil.`,
			func(values ...interface{}) interface{} {
				var dflt interface{}
				if math.Mod(float64(len(values)), 2) == 0 {
					//an even number of values was passed - the last value is the default
					dflt = values[len(values)-1]
					values = values[:len(values)-1]
				}
				for i := 1; i < len(values); i = i + 2 {
					if values[i] == values[0] {
						return values[i+1]
					}
				}
				return dflt
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"not",
			`Accepts a boolean & returns its inverse`,
			func(value bool) bool {
				return !value
			},
		),
	)

}
