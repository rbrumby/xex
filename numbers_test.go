package xex

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	//Get the conversion function for int
	convFunc, err := GetFunction("int")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0].(int) != expected[0].(int) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to int
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(int); ok {
			if v != expected {
				t.Errorf("int conversion expected %d, got %d", expected, v)
			}
		} else {
			t.Error("value is not int")
		}
	})
}
func TestInt32(t *testing.T) {
	//Get the conversion function for int32
	convFunc, err := GetFunction("int32")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0] != expected[0].(int32) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to int32
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(int32); ok {
			if v != int32(expected) {
				t.Errorf("int32 conversion expected %d, got %d", expected, v)
			}
		} else {
			t.Error("value is not int32")
		}
	})
}

func TestInt64(t *testing.T) {
	//Get the conversion function for int64
	convFunc, err := GetFunction("int64")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0] != expected[0].(int64) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to int64
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(int64); ok {
			if v != int64(expected) {
				t.Errorf("int64 conversion expected %d, got %d", expected, v)
			}
		} else {
			t.Error("value is not int64")
		}
	})
}
func TestUInt(t *testing.T) {
	//Get the conversion function for uint
	convFunc, err := GetFunction("uint")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0] != expected[0].(uint) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to uint
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(uint); ok {
			if v != uint(expected) {
				t.Errorf("uint conversion expected %d, got %d", expected, v)
			}
		} else {
			t.Error("value is not uint")
		}
	})
}

func TestUInt8(t *testing.T) {
	//Get the conversion function for uint8
	convFunc, err := GetFunction("uint8")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0] != expected[0].(uint8) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to uint8
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(uint8); ok {
			if v != uint8(expected) {
				t.Errorf("uint8 conversion expected %d, got %d", expected, v)
			}
		} else {
			t.Error("value is not uint8")
		}
	})
}

func TestUInt16(t *testing.T) {
	//Get the conversion function for uint16
	convFunc, err := GetFunction("uint16")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0] != expected[0].(uint16) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to uint16
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(uint16); ok {
			if v != uint16(expected) {
				t.Errorf("uint16 conversion expected %d, got %d", expected, v)
			}
		} else {
			t.Error("value is not uint16")
		}
	})
}

func TestUInt32(t *testing.T) {
	//Get the conversion function for uint32
	convFunc, err := GetFunction("uint32")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0] != expected[0].(uint32) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to uint32
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(uint32); ok {
			if v != uint32(expected) {
				t.Errorf("uint32 conversion expected %d, got %d", expected, v)
			}
		} else {
			t.Error("value is not uint32")
		}
	})
}

func TestUInt64(t *testing.T) {
	//Get the conversion function for uint64
	convFunc, err := GetFunction("uint64")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0] != expected[0].(uint64) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to uint64
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(uint64); ok {
			if v != uint64(expected) {
				t.Errorf("uint64 conversion expected %d, got %d", expected, v)
			}
		} else {
			t.Error("value is not uint64")
		}
	})
}

func TestFloat32(t *testing.T) {
	//Get the conversion function for float32
	convFunc, err := GetFunction("float32")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0] != expected[0].(float32) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to float32
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(float32); ok {
			if v != float32(expected) {
				t.Errorf("float32 conversion expected %d, got %f", expected, v)
			}
		} else {
			t.Error("value is not float32")
		}
	})
}

func TestFloat64(t *testing.T) {
	//Get the conversion function for float64
	convFunc, err := GetFunction("float64")
	if err != nil {
		t.Error(err)
		return
	}
	expected, err := convFunc.Exec(24)
	if err != nil {
		t.Error(err)
		return
	}
	answer := testMathsFunctions(t, convFunc)
	if answer[0] != expected[0].(float64) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}

	//test conversions to float64
	testConversion(t, convFunc, func(value interface{}, expected int) {
		if v, ok := value.(float64); ok {
			if v != float64(expected) {
				t.Errorf("float64 conversion expected %d, got %f", expected, v)
			}
		} else {
			t.Error("value is not float64")
		}
	})
}

//testMathsFunction returns 24 as a interface{} with underlying concrete type per the convFunc
func testMathsFunctions(t *testing.T, convFunc *Function) (answer []interface{}) {
	//Get all of the basic maths Functions
	add, err := GetFunction("add")
	if err != nil {
		t.Error(err)
		return
	}
	subtract, err := GetFunction("subtract")
	if err != nil {
		t.Error(err)
		return
	}
	multiply, err := GetFunction("multiply")
	if err != nil {
		t.Error(err)
		return
	}
	divide, err := GetFunction("divide")
	if err != nil {
		t.Error(err)
		return
	}
	//Use conv to convert vars to the required type
	init, err := convFunc.Exec(20) //initial 20
	if err != nil {
		t.Error(err)
		return
	}
	plus, err := convFunc.Exec(30) //plus 30 = 50
	if err != nil {
		t.Error(err)
		return
	}
	minus, err := convFunc.Exec(10) //minus 10 = 40
	if err != nil {
		t.Error(err)
		return
	}
	times, err := convFunc.Exec(3) //multiplied by 3 = 120
	if err != nil {
		t.Error(err)
		return
	}
	over, err := convFunc.Exec(5) //divded by 5 = 24
	if err != nil {
		t.Error(err)
		return
	}
	//Do the maths
	answer, err = add.Exec(init[0], plus[0])
	if err != nil {
		t.Error(err)
		return
	}
	answer, err = subtract.Exec(answer[0], minus[0])
	if err != nil {
		t.Error(err)
		return
	}
	answer, err = multiply.Exec(answer[0], times[0])
	if err != nil {
		t.Error(err)
		return
	}
	answer, err = divide.Exec(answer[0], over[0])
	if err != nil {
		t.Error(err)
		return
	}
	return
}

func testConversion(t *testing.T, convFunc *Function, validate func(value interface{}, expected int)) {
	i, err := convFunc.Exec(int(99))
	if err != nil {
		t.Errorf("converting int to %q: %s", convFunc.Name(), err)
		return
	}
	validate(i[0], 99)

	i32, err := convFunc.Exec(int32(99))
	if err != nil {
		t.Errorf("converting int32 to %q: %s", convFunc.Name(), err)
		return
	}
	validate(i32[0], 99)
	i64, err := convFunc.Exec(int64(99))
	if err != nil {
		t.Errorf("converting int64 to %q: %s", convFunc.Name(), err)
		return
	}
	validate(i64[0], 99)
	ui, err := convFunc.Exec(uint(99))
	if err != nil {
		t.Errorf("converting uint to %q: %s", convFunc.Name(), err)
		return
	}
	validate(ui[0], 99)
	ui8, err := convFunc.Exec(uint8(99))
	if err != nil {
		t.Errorf("converting uint8 to %q: %s", convFunc.Name(), err)
		return
	}
	validate(ui8[0], 99)
	ui16, err := convFunc.Exec(uint16(99))
	if err != nil {
		t.Errorf("converting uint16 to %q: %s", convFunc.Name(), err)
		return
	}
	validate(ui16[0], 99)
	ui32, err := convFunc.Exec(uint32(99))
	if err != nil {
		t.Errorf("converting uint32 to %q: %s", convFunc.Name(), err)
		return
	}
	validate(ui32[0], 99)
	ui64, err := convFunc.Exec(uint64(99))
	if err != nil {
		t.Errorf("converting uint64 to %q: %s", convFunc.Name(), err)
		return
	}
	validate(ui64[0], 99)
	f32, err := convFunc.Exec(float32(99))
	if err != nil {
		t.Errorf("converting float32 to %q: %s", convFunc.Name(), err)
		return
	}
	validate(f32[0], 99)
	f64, err := convFunc.Exec(float64(99))
	if err != nil {
		t.Errorf("converting float64 to %q: %s", convFunc.Name(), err)
		return
	}
	validate(f64[0], 99)
}
