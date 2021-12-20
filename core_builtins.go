package xex

import (
	"math"
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
			"switch",
			`Switches on the first value.
			The following values are equivalent to "case : result" pairs.
			If a final value is provided (an even number of arguments is passed in total), the final value is used as the default.
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

	RegisterFunction(
		NewFunction(
			"nil",
			`Returns what is passed - used to implement parenthesis grouping`,
			func(value interface{}) interface{} {
				return value
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"add_or_Concat",
			`Chooses to call add or concat depending on args`,
			func(val1 interface{}, val2 interface{}) (interface{}, error) {
				switch val1.(type) {
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					switch val2.(type) {
					case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
						add, err := GetFunction("add")
						if err != nil {
							return nil, err
						}
						res, err := add.Exec(val1, val2)
						if err != nil {
							return nil, err
						}
						return res[0], err
					}
				}
				concat, err := GetFunction("concat")
				if err != nil {
					return nil, err
				}
				res, err := concat.Exec(val1, val2)
				if err != nil {
					return nil, err
				}
				return res[0], err
			},
		),
	)

	//TODO: and, or, Greater, less than, etc

}
