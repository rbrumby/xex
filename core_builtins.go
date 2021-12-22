package xex

import (
	"math"
)

func registerCoreBuiltins() {

	RegisterFunction(
		NewFunction(
			"equals",
			FunctionDocumentation{
				Text: `compares 2 inputs returning a bool`,
				Parameters: map[string]string{
					"val1": "The first value to compare",
					"val2": "The second value to compare",
				},
			},
			func(val1, val2 interface{}) bool {
				return val1 == val2
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"switch",
			FunctionDocumentation{
				Text: `Switches on the first value.
				The following values are equivalent to "case : result" pairs.
				If a final value is provided (an even number of arguments is passed in total), the final value is used as the default.
				If value1 equals value2, value3 is returned. Else if value1 equals value4, value5 is returned. And so on.
				If there is no default and no values matched, switch returns nil.`,
				Parameters: map[string]string{
					"values": "variadic - the value to test then alternate if/else pairs and finally an optional else value",
				},
			},
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
			FunctionDocumentation{
				Text: `Accepts a boolean & returns its inverse`,
				Parameters: map[string]string{
					"value": "The value to invert.",
				},
			},
			func(value bool) bool {
				return !value
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"nil",
			FunctionDocumentation{
				Text: `Returns what is passed - used to implement parenthesis grouping`,
				Parameters: map[string]string{
					"value": "The value which will be returned as this function does nothing!",
				},
			},
			func(value interface{}) interface{} {
				return value
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"add_or_Concat",
			FunctionDocumentation{
				Text: `Chooses to call add or concat depending if args are numeric or not.`,
				Parameters: map[string]string{
					"val1": "The first value to add / concat.",
					"val2": "The second value to add / concat.",
				},
			},
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
