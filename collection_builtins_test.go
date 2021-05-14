package xex

import (
	"errors"
	"testing"
)

type env struct {
	Data interface{}
}
type Object struct {
	Value string
}

func TestSelectArray(t *testing.T) {
	sel, err := GetFunction("select")
	if err != nil {
		t.Fatal(err)
	}
	eq, err := GetFunction("equals")
	if err != nil {
		t.Fatal(err)
	}
	var arr [5]Object
	arr[0] = Object{"0"}
	arr[1] = Object{"1"}
	arr[2] = Object{"2"}
	arr[3] = Object{"1"}
	arr[4] = Object{"1"}
	fc := NewFunctionCall(
		sel,
		[]Node{
			NewProperty("Data", nil),
			NewExpression(
				NewFunctionCall(eq, []Node{NewProperty("Value", nil), NewLiteral("1")}, 0),
			),
		},
		0,
	)
	res, err := fc.Evaluate(env{arr})
	if err != nil {
		t.Fatal(err)
	}

	if r, ok := res.([]Object); ok {
		if len(r) != 3 {
			t.Errorf("Expected 3 results, got %d", len(r))
		}
	} else {
		t.Fatal("Did not get a slice of Object")
	}
}

func TestSelectSlice(t *testing.T) {

	sel, err := GetFunction("select")
	if err != nil {
		t.Fatal(err)
	}
	eq, err := GetFunction("equals")
	if err != nil {
		t.Fatal(err)
	}
	s := make([]Object, 5)
	s[0] = Object{"1"}
	s[1] = Object{"0"}
	s[2] = Object{"1"}
	s[3] = Object{"2"}
	s[4] = Object{"1"}
	fc := NewFunctionCall(
		sel,
		[]Node{
			NewProperty("Data", nil),
			NewExpression(
				NewFunctionCall(eq, []Node{NewProperty("Value", nil), NewLiteral("1")}, 0),
			),
		},
		0,
	)
	res, err := fc.Evaluate(env{s})
	if err != nil {
		t.Fatal(err)
	}

	if r, ok := res.([]Object); ok {
		if len(r) != 3 {
			t.Errorf("Expected 3 results, got %d", len(r))
		}
	} else {
		t.Fatal("Did not get a slice of Object")
	}
}

func TestSelectMap(t *testing.T) {
	fn, err := GetFunction("select")
	if err != nil {
		t.Fatal(err)
	}

	eq, err := GetFunction("equals")
	if err != nil {
		t.Fatal(err)
	}
	m := make(map[string]Object, 3)
	m["0"] = Object{"Zero"}
	m["1"] = Object{"One"}
	m["2"] = Object{"Two"}
	res, err := fn.Exec(
		m,
		NewExpression(
			NewFunctionCall(eq, []Node{NewProperty("Value", nil), NewLiteral("One")}, 0),
		),
	)
	if err != nil {
		t.Fatal(err)
	}
	if r, ok := res[0].(map[string]Object); ok {
		if len(r) != 1 || r["1"].Value != "One" {
			t.Fatalf("Unexpected result %v", r)
		}
	} else {
		t.Fatal("Did not get a map[string]Object")
	}

}

func TestSelectInvalidType(t *testing.T) {
	fn, err := GetFunction("select")
	if err != nil {
		t.Fatal(err)
	}
	_, err = fn.Exec("Nonsense", NewExpression(NewLiteral("Nothing")))
	if err == nil {
		t.Fatal(errors.New("Should have failed with invalid type for select"))
	}

}