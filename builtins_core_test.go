package xex

import (
	"testing"
)

type testSubStruct struct {
	value string
}

func TestEquals(t *testing.T) {
	fn, err := GetFunction("equals")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := fn.Exec("Hello", 5)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] == true {
		t.Error("values shouldn't match")
		return
	}
}

func TestNotEquals(t *testing.T) {
	fn, err := GetFunction("notEquals")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := fn.Exec("Hello", 5)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] == false {
		t.Error("values should be inequal (notEquals should return true)")
		return
	}
}

func TestSwitch(t *testing.T) {
	fn, err := GetFunction("switch")
	if err != nil {
		t.Error(err)
		return
	}

	values := []interface{}{"testval", "x", "y", "testval", testSubStruct{value: "answer"}}
	res, err := fn.Exec(values...)
	if err != nil {
		t.Error(err)
		return
	}
	if val, ok := res[0].(testSubStruct); !ok {
		t.Error("val is not a testSubStruct")
		return
	} else if val.value != "answer" {
		t.Errorf(`expected "answer", got %q`, res[0].(testSubStruct).value)
		return
	}

	values = []interface{}{
		"testval",
		"x", "y",
		"z", testSubStruct{value: "answer"},
		"default",
	}
	res, err = fn.Exec(values...)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != "default" {
		t.Errorf(`expected "default", got %q`, res[0])
		return
	}
}

func TestNot(t *testing.T) {
	fn, err := GetFunction("not")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := fn.Exec(false)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true, got %t`, res[0])
		return
	}

}

func TestAnd(t *testing.T) {
	fn, err := GetFunction("and")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := fn.Exec(false, true)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != false {
		t.Errorf(`expected false, got %t`, res[0])
		return
	}

}

func TestOr(t *testing.T) {
	fn, err := GetFunction("or")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := fn.Exec(false, true)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true, got %t`, res[0])
		return
	}

}

func TestGreaterThan(t *testing.T) {
	fn, err := GetFunction("greaterThan")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := fn.Exec(7.1, float64(5)) //float & int
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true, got %t`, res[0])
		return
	}

	res, err = fn.Exec("def", "abc") //strings
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true, got %t`, res[0])
		return
	}

	_, err = fn.Exec(7, "abc") //string & number
	if err == nil {
		t.Error("Expected error!")
		return
	}
}

func TestGreaterThanEqual(t *testing.T) {
	fn, err := GetFunction("greaterThanEqual")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := fn.Exec(7.1, 7.1) //floats
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true, got %t`, res[0])
		return
	}

	res, err = fn.Exec("def", "abc") //strings
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true, got %t`, res[0])
		return
	}

	_, err = fn.Exec(7, "abc") //string & number
	if err == nil {
		t.Error("Expected error!")
		return
	}
}

func TestLessThan(t *testing.T) {
	fn, err := GetFunction("lessThan")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := fn.Exec(7.1, 7.12) //floats
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true, got %t`, res[0])
		return
	}

	res, err = fn.Exec("abc", "abd") //strings
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true (strings), got %t`, res[0])
		return
	}

	res, err = fn.Exec("abc", "abc") //strings
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] == true {
		t.Errorf(`expected false (strings), got %t`, res[0])
		return
	}

	_, err = fn.Exec(7, "abc") //string & number
	if err == nil {
		t.Error("Expected error!")
		return
	}
}

func TestLessThanEqual(t *testing.T) {
	fn, err := GetFunction("lessThanEqual")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := fn.Exec(7.1, 7.1) //floats
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true, got %t`, res[0])
		return
	}

	res, err = fn.Exec("abc", "abc") //strings
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != true {
		t.Errorf(`expected true (strings), got %t`, res[0])
		return
	}

	_, err = fn.Exec(7, "abc") //string & number
	if err == nil {
		t.Error("Expected error!")
		return
	}
}

func TestAddOrConcat(t *testing.T) {
	fn, err := GetFunction("addOrConcat")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := fn.Exec(7.1, 7.1) //floats
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != 14.2 {
		t.Errorf("expected 14.2, got %d", res)
		return
	}

	res, err = fn.Exec("7.1", "7.1") //strings
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != "7.17.1" {
		t.Errorf("expected 14.2, got %d", res)
		return
	}

}
