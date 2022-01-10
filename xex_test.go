package xex

import (
	"testing"
)

func TestMain(m *testing.M) {
	//TODO: Make this configurable
	// Repolog.SetRepoLogLevel(capnslog.INFO)
	cfg, err := Repolog.ParseLogLevelConfig("xex=TRACE")
	if err != nil {
		panic(err)
	}
	Repolog.SetLogLevel(cfg)

	m.Run()
}

func TestLiteral(t *testing.T) {
	l := NewLiteral("testing123")
	e := NewExpression(l)
	if e, err := e.Evaluate(nil); err != nil {
		t.Error(err)
		return
	} else if e != "testing123" {
		t.Errorf("literal Eval() returned %v", e)
		return
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

	root := NewProperty("root", nil)
	p := NewProperty("Sub", root)
	ps := NewProperty("Str2", p)
	e := NewExpression(ps)

	result, err := e.Evaluate(Values{"root": x})
	if err != nil {
		t.Error(err)
		return
	}
	if result != "somethingelse" {
		t.Errorf("Expect \"somethingelse\". Got %q", result)
		return
	}
}

func TestMethodCall(t *testing.T) {

	lib := NewProperty("lib", nil)
	add := NewMethodCall("GetAddress", lib, nil, 0)
	exp := NewExpression(add)

	res, err := exp.Evaluate(Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
	r, ok := res.(Address)
	if !ok {
		t.Errorf("Expected address, got %v", r)
		return
	}
	if r.City != "London" {
		t.Errorf("Expected city London, got %s", r.City)
	}
}

func TestNonExistentMethodCall(t *testing.T) {
	mc := NewMethodCall("not_exists", nil, nil, 0)
	_, err := mc.Evaluate(Values{"lib": testLib})
	if err == nil {
		t.Error("expected cannot call method on nil object")
		return
	}
	lib := NewProperty("lib", nil)
	mc2 := NewMethodCall("not_exists", lib, nil, 0)
	_, err = mc2.Evaluate(Values{"lib": testLib})
	if err == nil {
		t.Error("expected method not_exists doesn't exist error")
		return
	}
}

func TestFunctionCalls(t *testing.T) {
	fnMultiply, err := GetFunction("multiply")
	if err != nil {
		t.Error(err)
		return
	}
	fnInt, err := GetFunction("int")
	if err != nil {
		t.Error(err)
		return
	}
	fnDivide, err := GetFunction("divide")
	if err != nil {
		t.Error(err)
		return
	}
	fnDivideCall := NewFunctionCall(
		fnDivide,
		[]Node{
			NewFunctionCall(
				fnMultiply,
				[]Node{
					NewLiteral(30),
					NewLiteral(150),
				},
				0,
			),
			NewFunctionCall(
				fnInt,
				[]Node{
					NewLiteral(4.76543), //int function will round down
				},
				0,
			),
		},
		0,
	)
	result, err := fnDivideCall.Evaluate(Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
	if result != 1125 {
		t.Errorf("Expected 1125, got %d", result)
		return
	}
}

func TestPropertyOfMethodCall(t *testing.T) {
	cityEx := NewProperty(
		"City",
		NewMethodCall(
			"GetAddress",
			NewProperty("lib", nil), //top level object
			nil,                     //no args to GetAddress
			0,                       //use 1st return value
		),
	)

	city, err := cityEx.Evaluate(Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}

	if city != "London" {
		t.Errorf("Expected London, got %s", city)
		return
	}
}

func TestPropertyOnPointer(t *testing.T) {

	book := NewMethodCall("Book", NewProperty("lib", nil), []Node{NewLiteral("1984")}, 0)
	price := NewProperty("Price", book)

	ex := NewExpression(price)

	res, err := ex.Evaluate(Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
	if p, ok := res.(float32); ok {
		if p != 9.99 {
			t.Errorf("Expected 9.99, got %v", res)
			return
		}
		return
	}
	t.Error("Could not convert price to float32")
}
