package xex

import (
	"fmt"
	"math"
	"reflect"
)

//Set up built-in number functions
func registerNumberBuiltins() {
	RegisterFunction(
		NewFunction(
			"add",
			FunctionDocumentation{
				Text: `adds two numbers returning a single numerical result`,
				Parameters: []FunctionDocParam{
					{"num1", "The first number to add."},
					{"num2", "The second number to add."},
				},
			},
			func(num1, num2 interface{}) (interface{}, error) {
				switch n1 := num1.(type) {
				//First make sure we got a number in num1
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					//Then make sure num2 is the same type
					if reflect.TypeOf(num1) != reflect.TypeOf(num2) {
						return 0, fmt.Errorf("add cannot use different types (%s & %s) - convert them first", reflect.TypeOf(num1).Name(), reflect.TypeOf(num2).Name())
					}
					//Then do all the conversions & additions (now we know both types are the same)
					switch n2 := num2.(type) {
					case int:
						return n1.(int) + n2, nil
					case int8:
						return n1.(int8) + n2, nil
					case int16:
						return n1.(int16) + n2, nil
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
			FunctionDocumentation{
				Text: `subtracts two numbers returning a single numerical result`,
				Parameters: []FunctionDocParam{
					{"minuend", "The initial number to subtract from."},
					{"subtrahend", "The value to subreact from minuend."},
				},
			},
			func(minuend, subtrahend interface{}) (interface{}, error) {
				switch n1 := minuend.(type) {
				//First make sure we got a number in num1
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					//Then make sure num2 is the same type
					if reflect.TypeOf(minuend) != reflect.TypeOf(subtrahend) {
						return 0, fmt.Errorf("subtract cannot use different types (%s & %s) - convert them first", reflect.TypeOf(minuend).Name(), reflect.TypeOf(subtrahend).Name())
					}
					//Then do all the conversions & additions (now we know both types are the same)
					switch n2 := subtrahend.(type) {
					case int:
						return n1.(int) - n2, nil
					case int8:
						return n1.(int8) - n2, nil
					case int16:
						return n1.(int16) - n2, nil
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
				return 0, fmt.Errorf("subtract can only subtract numeric types, not %s and %s", reflect.TypeOf(minuend).Name(), reflect.TypeOf(subtrahend).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"multiply",
			FunctionDocumentation{
				Text: `multiplies two numbers returning a single numerical result`,
				Parameters: []FunctionDocParam{
					{"multiplicand", "The number to be multiplied."},
					{"multiplier", "The number to multiply by."},
				},
			},
			func(multiplicand, multiplier interface{}) (interface{}, error) {
				switch n1 := multiplicand.(type) {
				//First make sure we got a number in num1
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					//Then make sure num2 is the same type
					if reflect.TypeOf(multiplicand) != reflect.TypeOf(multiplier) {
						return 0, fmt.Errorf("multiply cannot use different types (%s & %s) - convert them first", reflect.TypeOf(multiplicand).Name(), reflect.TypeOf(multiplier).Name())
					}
					//Then do all the conversions & additions (now we know both types are the same)
					switch n2 := multiplier.(type) {
					case int:
						return n1.(int) * n2, nil
					case int8:
						return n1.(int8) * n2, nil
					case int16:
						return n1.(int16) * n2, nil
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
				return 0, fmt.Errorf("multiply can only add numeric types, not %s and %s", reflect.TypeOf(multiplicand).Name(), reflect.TypeOf(multiplier).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"divide",
			FunctionDocumentation{
				Text: `divides two numbers returning a single numerical result`,
				Parameters: []FunctionDocParam{
					{"dividend", "The number to be divided."},
					{"divisor", "The number to divide by."},
				},
			},
			func(dividend, divisor interface{}) (interface{}, error) {
				switch n1 := dividend.(type) {
				//First make sure we got a number in num1
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					//Then make sure num2 is the same type
					if reflect.TypeOf(dividend) != reflect.TypeOf(divisor) {
						return 0, fmt.Errorf("divide cannot use different types (%s & %s) - convert them first", reflect.TypeOf(dividend).Name(), reflect.TypeOf(divisor).Name())
					}
					//Then do all the conversions & additions (now we know both types are the same)
					switch n2 := divisor.(type) {
					case int:
						return n1.(int) / n2, nil
					case int8:
						return n1.(int8) / n2, nil
					case int16:
						return n1.(int16) / n2, nil
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
				return 0, fmt.Errorf("divide can only divide numeric types, not %s and %s", reflect.TypeOf(dividend).Name(), reflect.TypeOf(divisor).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"pow",
			FunctionDocumentation{
				Text: `pow returns x to the power of y (x**y).`,
				Parameters: []FunctionDocParam{
					{"x", "The base number."},
					{"y", "The exponent (number of times x is multiplied by itself)."},
				},
			},
			math.Pow,
		),
	)

	RegisterFunction(
		NewFunction(
			"mod",
			FunctionDocumentation{
				Text: `mod returns the remainder of dividend divided by divisor.`,
				Parameters: []FunctionDocParam{
					{"dividend", "The number to be divided."},
					{"divisor", "The number to divide by."},
				},
			},
			math.Mod,
		),
	)

	RegisterFunction(
		NewFunction(
			"int",
			FunctionDocumentation{
				Text: `int converts the passed in value to an int or returns a error if conversion isn't possible`,
			},
			func(number interface{}) (int, error) {
				switch num := number.(type) {
				case int:
					return num, nil
				case int8:
					return int(num), nil
				case int16:
					return int(num), nil
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
				return 0, fmt.Errorf("cannot convert %s to int", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"int8",
			FunctionDocumentation{
				Text: `int8 converts the passed in value to an int8 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (int8, error) {
				switch num := number.(type) {
				case int:
					return int8(num), nil
				case int8:
					return num, nil
				case int16:
					return int8(num), nil
				case int32:
					return int8(num), nil
				case int64:
					return int8(num), nil
				case uint:
					return int8(num), nil
				case uint8:
					return int8(num), nil
				case uint16:
					return int8(num), nil
				case uint32:
					return int8(num), nil
				case uint64:
					return int8(num), nil
				case float32:
					return int8(num), nil
				case float64:
					return int8(num), nil
				}
				return 0, fmt.Errorf("cannot convert %s to int8", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"int16",
			FunctionDocumentation{
				Text: `int16 converts the passed in value to an int16 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (int16, error) {
				switch num := number.(type) {
				case int:
					return int16(num), nil
				case int8:
					return int16(num), nil
				case int16:
					return num, nil
				case int32:
					return int16(num), nil
				case int64:
					return int16(num), nil
				case uint:
					return int16(num), nil
				case uint8:
					return int16(num), nil
				case uint16:
					return int16(num), nil
				case uint32:
					return int16(num), nil
				case uint64:
					return int16(num), nil
				case float32:
					return int16(num), nil
				case float64:
					return int16(num), nil
				}
				return 0, fmt.Errorf("cannot convert %s to int16", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"int32",
			FunctionDocumentation{
				Text: `int32 converts the passed in value to an int32 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (int32, error) {
				switch num := number.(type) {
				case int:
					return int32(num), nil
				case int8:
					return int32(num), nil
				case int16:
					return int32(num), nil
				case int32:
					return num, nil
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
				return 0, fmt.Errorf("cannot convert %s to int32", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"int64",
			FunctionDocumentation{
				Text: `int64 converts the passed in value to an int64 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (int64, error) {
				switch num := number.(type) {
				case int:
					return int64(num), nil
				case int8:
					return int64(num), nil
				case int16:
					return int64(num), nil
				case int32:
					return int64(num), nil
				case int64:
					return num, nil
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
				return 0, fmt.Errorf("cannot convert %s to int64", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint",
			FunctionDocumentation{
				Text: `uint converts the passed in value to an uint or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (uint, error) {
				switch num := number.(type) {
				case int:
					return uint(num), nil
				case int8:
					return uint(num), nil
				case int16:
					return uint(num), nil
				case int32:
					return uint(num), nil
				case int64:
					return uint(num), nil
				case uint:
					return num, nil
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
				return 0, fmt.Errorf("cannot convert %s to uint", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint8",
			FunctionDocumentation{
				Text: `uint8 converts the passed in value to an uint8 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (uint8, error) {
				switch num := number.(type) {
				case int:
					return uint8(num), nil
				case int8:
					return uint8(num), nil
				case int16:
					return uint8(num), nil
				case int32:
					return uint8(num), nil
				case int64:
					return uint8(num), nil
				case uint:
					return uint8(num), nil
				case uint8:
					return num, nil
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
				return 0, fmt.Errorf("cannot convert %s to uint8", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint16",
			FunctionDocumentation{
				Text: `uint16 converts the passed in value to an uint16 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (uint16, error) {
				switch num := number.(type) {
				case int:
					return uint16(num), nil
				case int8:
					return uint16(num), nil
				case int16:
					return uint16(num), nil
				case int32:
					return uint16(num), nil
				case int64:
					return uint16(num), nil
				case uint:
					return uint16(num), nil
				case uint8:
					return uint16(num), nil
				case uint16:
					return num, nil
				case uint32:
					return uint16(num), nil
				case uint64:
					return uint16(num), nil
				case float32:
					return uint16(num), nil
				case float64:
					return uint16(num), nil
				}
				return 0, fmt.Errorf("cannot convert %s to uint16", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint32",
			FunctionDocumentation{
				Text: `uint32 converts the passed in value to an uint32 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (uint32, error) {
				switch num := number.(type) {
				case int:
					return uint32(num), nil
				case int8:
					return uint32(num), nil
				case int16:
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
				case uint32:
					return num, nil
				case uint64:
					return uint32(num), nil
				case float32:
					return uint32(num), nil
				case float64:
					return uint32(num), nil
				}
				return 0, fmt.Errorf("cannot convert %s to uint32", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"uint64",
			FunctionDocumentation{
				Text: `uint64 converts the passed in value to an uint64 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (uint64, error) {
				switch num := number.(type) {
				case int:
					return uint64(num), nil
				case int8:
					return uint64(num), nil
				case int16:
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
				case uint64:
					return num, nil
				case float32:
					return uint64(num), nil
				case float64:
					return uint64(num), nil
				}
				return 0, fmt.Errorf("cannot convert %s to uint64", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"float32",
			FunctionDocumentation{
				Text: `float32 converts the passed in value to an float32 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (float32, error) {
				switch num := number.(type) {
				case int:
					return float32(num), nil
				case int8:
					return float32(num), nil
				case int16:
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
				case float32:
					return num, nil
				case float64:
					return float32(num), nil
				}
				return 0, fmt.Errorf("cannot convert %s to float32", reflect.TypeOf(number).Name())
			},
		),
	)

	RegisterFunction(
		NewFunction(
			"float64",
			FunctionDocumentation{
				Text: `float64 converts the passed in value to an float64 or returns a error if conversion isn't possible`,
				Parameters: []FunctionDocParam{
					{"number", "The number to convert."},
				},
			},
			func(number interface{}) (float64, error) {
				switch num := number.(type) {
				case int:
					return float64(num), nil
				case int8:
					return float64(num), nil
				case int16:
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
				case float64:
					return num, nil
				}
				return 0, fmt.Errorf("cannot convert %s to float64", reflect.TypeOf(number).Name())
			},
		),
	)
}
