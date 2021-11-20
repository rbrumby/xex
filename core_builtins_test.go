package xex

import (
	"testing"
)

type testSubStruct struct {
	value string
}

type testStruct struct {
	sub testSubStruct
}

func TestEquals(t *testing.T) {
	fn, err := GetFunction("equals")
	if err != nil {
		t.Fatal(err)
	}
	res, err := fn.Exec("Hello", 5)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] == true {
		t.Fatal("values shouldn't match")
	}

	//Using a FunctionCall
	t1 := testStruct{sub: testSubStruct{value: "AbC"}}
	t2 := testStruct{sub: testSubStruct{value: "AbC"}}
	fnc := NewFunctionCall(
		fn,
		[]Node{
			NewProperty("tOne", nil),
			NewProperty("tTwo", nil),
		},
		0,
	)
	res2, err := fnc.Evaluate(Values{"tOne": t1, "tTwo": t2})

	// res, err = fn.Exec(t1, t2)
	if err != nil {
		t.Fatal(err)
	}
	if res2 != true {
		t.Fatal("values should match")
	}
}

func TestSwitch(t *testing.T) {
	fn, err := GetFunction("switch")
	if err != nil {
		t.Fatal(err)
	}

	values := []interface{}{"testval", "x", "y", "testval", testSubStruct{value: "answer"}}
	res, err := fn.Exec(values...)
	if err != nil {
		t.Fatal(err)
	}
	if val, ok := res[0].(testSubStruct); !ok {
		t.Fatal("val is not a testSubStruct")
	} else if val.value != "answer" {
		t.Fatalf(`expected "answer", got %q`, res[0].(testSubStruct).value)
	}

	values = []interface{}{
		"testval",
		"x", "y",
		"z", testSubStruct{value: "answer"},
		"default",
	}
	res, err = fn.Exec(values...)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != "default" {
		t.Fatalf(`expected "default", got %q`, res[0])
	}
}

func TestNot(t *testing.T) {
	fn, err := GetFunction("not")
	if err != nil {
		t.Fatal(err)
	}

	res, err := fn.Exec(false)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != true {
		t.Fatalf(`expected true, got %t`, res[0])
	}

}
