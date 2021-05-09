package xex

import (
	"reflect"
	"testing"
)

func TestLiteral(t *testing.T) {
	l := &Literal{value: "testing123"}
	e := &Expression{root: l}
	if e, err := e.Evaluate(struct{}{}); err != nil {
		t.Fatal(err)
	} else if e != "testing123" {
		t.Fatalf("literal Eval() returned %v", e)
	}
}

func TestProperty(t *testing.T) {
	//our test graph
	x := struct {
		Str string
		Sub struct{ Str2 string }
	}{
		Str: "something",
		Sub: struct{ Str2 string }{
			Str2: "somethingelse",
		},
	}

	p := &Property{
		name: "Sub",
	}

	ps := &Property{
		name:   "Str2",
		parent: p,
	}
	e := &Expression{root: ps}
	result, err := e.Evaluate(x)
	if err != nil {
		t.Fatal(err)
	}
	if result != "somethingelse" {
		t.Fatalf("Expect \"somethingelse\". Got %q", result)
	}
}

type Car struct {
	Engine Engine
}

type Engine struct {
	Gearbox Gearbox
}

type Gearbox struct {
	Gear uint8
}

func (e *Engine) GetRPM(mph int) (rpm int) {
	return mph * 150 / int(e.Gearbox.Gear)
}

func TestMethodCall(t *testing.T) {
	c := Car{
		Engine: Engine{
			Gearbox: Gearbox{
				Gear: 4,
			},
		},
	}

	eProp := &Property{name: "Engine"}

	mphArg := &Literal{value: 30}

	rpm := &MethodCall{
		name:      "GetRPM",
		parent:    eProp,
		arguments: []Node{mphArg},
	}

	exp := Expression{root: rpm}

	res, err := exp.Evaluate(c)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.TypeOf(res).Kind() != reflect.Int {
		t.Fatal("RPM's are not an int")
	}
	if res != 1125 {
		t.Fatalf("Expected 1125, got %d", res)
	}
}

func TestFunctionCall(t *testing.T) {
	c := Car{
		Engine: Engine{
			Gearbox: Gearbox{
				Gear: 4,
			},
		},
	}

	eProp := &Property{name: "Engine"}

	gbProp := &Property{name: "Gearbox", parent: eProp}

	gProp := &Property{name: "Gear", parent: gbProp}
	gProp.SetParent(gbProp)

	fnMultiply, err := GetFunction("multiply")
	if err != nil {
		t.Fatal(err)
	}

	fnInt, err := GetFunction("int")
	if err != nil {
		t.Fatal(err)
	}

	fnDivide, err := GetFunction("divide")
	if err != nil {
		t.Fatal(err)
	}
	fnDivideCall := FunctionCall{
		function: fnDivide,
		arguments: []Node{
			&FunctionCall{
				function: fnMultiply,
				arguments: []Node{
					&Literal{value: 30},
					&Literal{value: 150}},
			},
			&FunctionCall{
				function: fnInt,
				arguments: []Node{
					gProp,
				},
			}},
	}

	result, err := fnDivideCall.Evaluate(c)
	if err != nil {
		t.Fatal(err)
	}

	if result != 1125 {
		t.Fatalf("Expected 1125, got %d", result)
	}

}
