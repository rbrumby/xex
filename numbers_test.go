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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(int) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(int32) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(int64) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(uint) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(uint8) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(uint16) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(uint32) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(uint64) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(float32) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
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
	answer := testGetAnswer(t, convFunc)
	if answer[0] != expected[0].(float64) {
		t.Error(fmt.Errorf("Expected %v, got %v", expected[0], answer[0]))
	}
}

func testGetAnswer(t *testing.T, convFunc *Function) (answer []interface{}) {
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
	init, err := convFunc.Exec(20)
	if err != nil {
		t.Error(err)
		return
	}
	plus, err := convFunc.Exec(30) //50
	if err != nil {
		t.Error(err)
		return
	}
	minus, err := convFunc.Exec(10) //40
	if err != nil {
		t.Error(err)
		return
	}
	times, err := convFunc.Exec(3) //120
	if err != nil {
		t.Error(err)
		return
	}
	over, err := convFunc.Exec(5) //24
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

// func TestAdd(t *testing.T) {
// 	add, err := GetFunction("add")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	//int
// 	sum, err := add.Exec(999994, 3)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if sum[0] != 999997 {
// 		t.Errorf("Expected 999997, got %d", sum[0])
// 	}
// 	//int32
// 	sum, err = add.Exec(int32(999994), int32(3))
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if sum[0] != int32(999997) {
// 		t.Errorf("Expected 999997, got %d", sum[0])
// 	}
// 	//int64
// 	sum, err = add.Exec(int64(999994), int64(3))
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if sum[0] != int64(999997) {
// 		t.Errorf("Expected 999997, got %d", sum[0])
// 	}

// 	sum, err = add.Exec(float64(999994), 3.5)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if sum[0] != 999997.5 {
// 		t.Errorf("Expected 999997.5, got %f", sum[0])
// 	}

// 	sum, err = add.Exec(uint32(666), uint32(333))
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if sum[0] != uint32(999) {
// 		t.Errorf("Expected 999, got %d", sum[0])
// 	}
// }

// func TestAddUnmatchedTypes(t *testing.T) {
// 	add, err := GetFunction("add")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	_, err = add.Exec(float64(999994), int(5))
// 	if err == nil {
// 		t.Error(errors.New("This should have failed"))
// 		return
// 	}
// 	// log.Println(sum[0], err)
// }

// func TestSubtract(t *testing.T) {
// 	sub, err := GetFunction("subtract")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	sum, err := sub.Exec(float64(999994), 3.5)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if sum[0] != 999990.5 {
// 		t.Errorf("Expected 999990.5, got %f", sum[0])
// 	}

// 	sum, err = sub.Exec(uint16(9994), uint16(9993))
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if sum[0] != uint16(1) {
// 		t.Errorf("Expected 1, got %f", sum[0])
// 	}
// }
