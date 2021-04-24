package xex

import (
	"fmt"
	"math"
	"reflect"
)

//Set up built-in string functions
func registerNumberBuiltins() {
	RegisterFunction(
		NewFunction(
			"add",
			`adds two numbers returning a single numerical result`,
			func(num1, num2 interface{}) (interface{}, error) {
				switch n1 := num1.(type) {
				//First make sure we got a number in num1
				case int, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					//Then make sure num2 is the same type
					if reflect.TypeOf(num1) != reflect.TypeOf(num2) {
						return 0, fmt.Errorf("add cannot use different types (%s & %s) - convert them first", reflect.TypeOf(num1).Name(), reflect.TypeOf(num2).Name())
					}
					//Then do all the conversions & additions (now we know both types are the same)
					switch n2 := num2.(type) {
					case int:
						return n1.(int) + n2, nil
					case int32:
						return n1.(int32) + n2, nil
					case int64:
						return n1.(int64) + n2, nil
					case uint:
						return n1.(uint) + n2, nil
					case uint8:
						return n1.(uint8) + n2, nil
					case uint16:
						return n1.(uint16) + n2, nil
					case uint32:
						return n1.(uint32) + n2, nil
					case uint64:
						return n1.(uint64) + n2, nil
					case float32:
						return n1.(float32) + n2, nil
					case float64:
						return n1.(float64) + n2, nil
					}
				}
				return 0, fmt.Errorf("add can only add numeric types, not %s and %s", reflect.TypeOf(num1).Name(), reflect.TypeOf(num2).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"subtract",
			`subtracts two numbers returning a single numerical result`,
			func(num1, num2 interface{}) (interface{}, error) {
				switch n1 := num1.(type) {
				//First make sure we got a number in num1
				case int, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					//Then make sure num2 is the same type
					if reflect.TypeOf(num1) != reflect.TypeOf(num2) {
						return 0, fmt.Errorf("subtract cannot use different types (%s & %s) - convert them first", reflect.TypeOf(num1).Name(), reflect.TypeOf(num2).Name())
					}
					//Then do all the conversions & additions (now we know both types are the same)
					switch n2 := num2.(type) {
					case int:
						return n1.(int) - n2, nil
					case int32:
						return n1.(int32) - n2, nil
					case int64:
						return n1.(int64) - n2, nil
					case uint:
						return n1.(uint) - n2, nil
					case uint8:
						return n1.(uint8) - n2, nil
					case uint16:
						return n1.(uint16) - n2, nil
					case uint32:
						return n1.(uint32) - n2, nil
					case uint64:
						return n1.(uint64) - n2, nil
					case float32:
						return n1.(float32) - n2, nil
					case float64:
						return n1.(float64) - n2, nil
					}
				}
				return 0, fmt.Errorf("subtract can only subtract numeric types, not %s and %s", reflect.TypeOf(num1).Name(), reflect.TypeOf(num2).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"multiply",
			`multiplies two numbers returning a single numerical result`,
			func(num1, num2 interface{}) (interface{}, error) {
				switch n1 := num1.(type) {
				//First make sure we got a number in num1
				case int, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					//Then make sure num2 is the same type
					if reflect.TypeOf(num1) != reflect.TypeOf(num2) {
						return 0, fmt.Errorf("multiply cannot use different types (%s & %s) - convert them first", reflect.TypeOf(num1).Name(), reflect.TypeOf(num2).Name())
					}
					//Then do all the conversions & additions (now we know both types are the same)
					switch n2 := num2.(type) {
					case int:
						return n1.(int) * n2, nil
					case int32:
						return n1.(int32) * n2, nil
					case int64:
						return n1.(int64) * n2, nil
					case uint:
						return n1.(uint) * n2, nil
					case uint8:
						return n1.(uint8) * n2, nil
					case uint16:
						return n1.(uint16) * n2, nil
					case uint32:
						return n1.(uint32) * n2, nil
					case uint64:
						return n1.(uint64) * n2, nil
					case float32:
						return n1.(float32) * n2, nil
					case float64:
						return n1.(float64) * n2, nil
					}
				}
				return 0, fmt.Errorf("multiply can only add numeric types, not %s and %s", reflect.TypeOf(num1).Name(), reflect.TypeOf(num2).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"divide",
			`divides two numbers returning a single numerical result`,
			func(num1, num2 interface{}) (interface{}, error) {
				switch n1 := num1.(type) {
				//First make sure we got a number in num1
				case int, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					//Then make sure num2 is the same type
					if reflect.TypeOf(num1) != reflect.TypeOf(num2) {
						return 0, fmt.Errorf("divide cannot use different types (%s & %s) - convert them first", reflect.TypeOf(num1).Name(), reflect.TypeOf(num2).Name())
					}
					//Then do all the conversions & additions (now we know both types are the same)
					switch n2 := num2.(type) {
					case int:
						return n1.(int) / n2, nil
					case int32:
						return n1.(int32) / n2, nil
					case int64:
						return n1.(int64) / n2, nil
					case uint:
						return n1.(uint) / n2, nil
					case uint8:
						return n1.(uint8) / n2, nil
					case uint16:
						return n1.(uint16) / n2, nil
					case uint32:
						return n1.(uint32) / n2, nil
					case uint64:
						return n1.(uint64) / n2, nil
					case float32:
						return n1.(float32) / n2, nil
					case float64:
						return n1.(float64) / n2, nil
					}
				}
				return 0, fmt.Errorf("divide can only divide numeric types, not %s and %s", reflect.TypeOf(num1).Name(), reflect.TypeOf(num2).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"pow",
			`pow returns x to the power of y (x**y)`,
			math.Pow,
		),
	)

	RegisterFunction(
		NewFunction(
			"mod",
			`mod returns the remainder of x/y`,
			math.Mod,
		),
	)

	RegisterFunction(
		NewFunction(
			"int",
			`int converts the passed in value to an int or returns a error if conversion isn't possible`,
			func(number interface{}) (int, error) {
				switch num := number.(type) {
				case int:
					return num, nil
				case int32:
					return int(num), nil
				case int64:
					return int(num), nil
				case uint:
					return int(num), nil
				case uint8:
					return int(num), nil
				case uint16:
					return int(num), nil
				case uint32:
					return int(num), nil
				case uint64:
					return int(num), nil
				case float32:
					return int(num), nil
				case float64:
					return int(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to int", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"int32",
			`int32 converts the passed in value to an int32 or returns a error if conversion isn't possible`,
			func(number interface{}) (int32, error) {
				switch num := number.(type) {
				case int32:
					return num, nil
				case int:
					return int32(num), nil
				case int64:
					return int32(num), nil
				case uint:
					return int32(num), nil
				case uint8:
					return int32(num), nil
				case uint16:
					return int32(num), nil
				case uint32:
					return int32(num), nil
				case uint64:
					return int32(num), nil
				case float32:
					return int32(num), nil
				case float64:
					return int32(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to int32", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"int64",
			`int converts the passed in value to an int or returns a error if conversion isn't possible`,
			func(number interface{}) (int64, error) {
				switch num := number.(type) {
				case int64:
					return num, nil
				case int:
					return int64(num), nil
				case int32:
					return int64(num), nil
				case uint:
					return int64(num), nil
				case uint8:
					return int64(num), nil
				case uint16:
					return int64(num), nil
				case uint32:
					return int64(num), nil
				case uint64:
					return int64(num), nil
				case float32:
					return int64(num), nil
				case float64:
					return int64(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to int64", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint",
			`uint converts the passed in value to a uint or returns a error if conversion isn't possible`,
			func(number interface{}) (uint, error) {
				switch num := number.(type) {
				case uint:
					return num, nil
				case int:
					return uint(num), nil
				case int32:
					return uint(num), nil
				case int64:
					return uint(num), nil
				case uint8:
					return uint(num), nil
				case uint16:
					return uint(num), nil
				case uint32:
					return uint(num), nil
				case uint64:
					return uint(num), nil
				case float32:
					return uint(num), nil
				case float64:
					return uint(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to uint", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint8",
			`uint8 converts the passed in value to a uint8 or returns a error if conversion isn't possible`,
			func(number interface{}) (uint8, error) {
				switch num := number.(type) {
				case uint8:
					return num, nil
				case int:
					return uint8(num), nil
				case int32:
					return uint8(num), nil
				case int64:
					return uint8(num), nil
				case uint:
					return uint8(num), nil
				case uint16:
					return uint8(num), nil
				case uint32:
					return uint8(num), nil
				case uint64:
					return uint8(num), nil
				case float32:
					return uint8(num), nil
				case float64:
					return uint8(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to uint8", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint16",
			`uint16 converts the passed in value to a uint16 or returns a error if conversion isn't possible`,
			func(number interface{}) (uint16, error) {
				switch num := number.(type) {
				case uint16:
					return num, nil
				case int:
					return uint16(num), nil
				case int32:
					return uint16(num), nil
				case int64:
					return uint16(num), nil
				case uint8:
					return uint16(num), nil
				case uint:
					return uint16(num), nil
				case uint32:
					return uint16(num), nil
				case uint64:
					return uint16(num), nil
				case float32:
					return uint16(num), nil
				case float64:
					return uint16(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to uint16", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint32",
			`uint32 converts the passed in value to a uint32 or returns a error if conversion isn't possible`,
			func(number interface{}) (uint32, error) {
				switch num := number.(type) {
				case uint32:
					return num, nil
				case int:
					return uint32(num), nil
				case int32:
					return uint32(num), nil
				case int64:
					return uint32(num), nil
				case uint:
					return uint32(num), nil
				case uint8:
					return uint32(num), nil
				case uint16:
					return uint32(num), nil
				case uint64:
					return uint32(num), nil
				case float32:
					return uint32(num), nil
				case float64:
					return uint32(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to uint32", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint64",
			`uint64 converts the passed in value to a uint64 or returns a error if conversion isn't possible`,

			func(number interface{}) (uint64, error) {
				switch num := number.(type) {
				case uint64:
					return num, nil
				case int:
					return uint64(num), nil
				case int32:
					return uint64(num), nil
				case int64:
					return uint64(num), nil
				case uint:
					return uint64(num), nil
				case uint8:
					return uint64(num), nil
				case uint16:
					return uint64(num), nil
				case uint32:
					return uint64(num), nil
				case float32:
					return uint64(num), nil
				case float64:
					return uint64(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to uint64", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"float32",
			`float converts the passed in value to a float32 or returns a error if conversion isn't possible`,
			func(number interface{}) (float32, error) {
				switch num := number.(type) {
				case float32:
					return num, nil
				case int:
					return float32(num), nil
				case int32:
					return float32(num), nil
				case int64:
					return float32(num), nil
				case uint:
					return float32(num), nil
				case uint8:
					return float32(num), nil
				case uint16:
					return float32(num), nil
				case uint32:
					return float32(num), nil
				case uint64:
					return float32(num), nil
				case float64:
					return float32(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to float32", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"float64",
			`float64 converts the passed in value to a float64 or returns a error if conversion isn't possible`,
			func(number interface{}) (float64, error) {
				switch num := number.(type) {
				case float64:
					return num, nil
				case int:
					return float64(num), nil
				case int32:
					return float64(num), nil
				case int64:
					return float64(num), nil
				case uint:
					return float64(num), nil
				case uint8:
					return float64(num), nil
				case uint16:
					return float64(num), nil
				case uint32:
					return float64(num), nil
				case uint64:
					return float64(num), nil
				case float32:
					return float64(num), nil
				}
				return 0, fmt.Errorf("Cannot convert %s to float64", reflect.TypeOf(number).Name())
			},
		),
	)
}
