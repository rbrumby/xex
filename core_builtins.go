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
			"decode",
			`Tests the first value against the even-positioned values in the value pairs which follow: 
			 If value1 equals value2, value3 is returned. Else if value1 equals value4, value5 is returned.
			 And so on.
			 If an even number of values are passed, the last value is not part of a pair and is
			 therefore used as the default if non of the previous values match.
			 If there is no default and no values matched, returns nil.`,
			func(values []interface{}) interface{} {
				var dflt interface{}
				if math.Mod(float64(len(values)), 2) == 0 {
					//an even number of values was passed - the last value is the default
					dflt = values[len(values)-1]
					values = values[:len(values)-1]
				}
				for i := 1; i < len(values); i = i + 2 {
					if values[0] == values[i] {
						return values[i+1]
					}
				}
				return dflt
			},
		),
	)

}
