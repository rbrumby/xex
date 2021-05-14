package xex

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/coreos/capnslog"
)

func TestMain(m *testing.M) {
	//TODO: Make this configurable
	Repolog.SetRepoLogLevel(capnslog.TRACE)
	// cfg, err := Repolog.ParseLogLevelConfig(logCfg)
	// if err != nil {
	// 	panic(err)
	// }
	// Repolog.SetLogLevel(cfg)

	m.Run()
}

func TestLiteral(t *testing.T) {
	l := NewLiteral("testing123")
	e := NewExpression(l)
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

	p := NewProperty("Sub", nil)

	ps := NewProperty("Str2", p)
	e := NewExpression(ps)
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
	Driver *Driver
}

func (c Car) GetGearBox() Gearbox {
	return c.Engine.Gearbox
}

func (c Car) String() string {
	return fmt.Sprintf("Car driven by %s (%d) is in gear %d", c.Driver.Name, c.Driver.Age, c.Engine.Gearbox.Gear)
}

type Driver struct {
	Name string
	Age  int
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

	eProp := NewProperty("Engine", nil)

	mphArg := NewLiteral(30)

	rpm := NewMethodCall("GetRPM", eProp, []Node{mphArg}, 0)

	exp := NewExpression(rpm)

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

func TestNonExistentMethodCall(t *testing.T) {
	c := Car{
		Engine: Engine{
			Gearbox: Gearbox{
				Gear: 4,
			},
		},
	}
	mc := NewMethodCall("not_exists", nil, nil, 0)
	_, err := mc.Evaluate(c)
	if err == nil {
		t.Fatal("expected method not_exists doesn't exist error")
	}

	e := NewProperty("Engine", nil)
	mc2 := NewMethodCall("not_exists", e, nil, 0)
	_, err = mc2.Evaluate(c)
	if err == nil {
		t.Fatal("expected method not_exists doesn't exist error")
	}
}

func TestFunctionCalls(t *testing.T) {
	c := Car{
		Engine: Engine{
			Gearbox: Gearbox{
				Gear: 4,
			},
		},
		Driver: &Driver{Name: "Stig"},
	}

	eProp := NewProperty("Engine", nil)

	gbProp := NewProperty("Gearbox", eProp)

	gProp := NewProperty("Gear", gbProp)

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
					gProp,
				},
				0,
			),
		},
		0,
	)

	result, err := fnDivideCall.Evaluate(c)
	if err != nil {
		t.Fatal(err)
	}

	if result != 1125 {
		t.Fatalf("Expected 1125, got %d", result)
	}

}

func TestPropertyOfMethodCall(t *testing.T) {
	c := Car{
		Engine: Engine{
			Gearbox: Gearbox{
				Gear: 4,
			},
		},
	}

	gearNode := NewProperty(
		"Gear",
		NewMethodCall(
			"GetGearBox",
			nil, //no parent
			nil, //no args
			0,
		),
	)

	gear, err := gearNode.Evaluate(c)
	if err != nil {
		t.Fatal(err)
	}

	if gear != uint8(4) {
		t.Fatalf("Expected gear 4, got %d", gear)
	}
}

func TestPropertiesWithPointers(t *testing.T) {
	name := "Stig"
	c := Car{
		Driver: &Driver{
			Name: name,
			Age:  999,
		},
	}

	d := NewProperty("Driver", nil)
	n := NewProperty("Name", d)

	ex := NewExpression(n)

	res, err := ex.Evaluate(c)
	if err != nil {
		t.Fatal(err)
	}
	if res.(string) != "Stig" {
		t.Fatalf("Expected Stig, got %s", res)
	}
}

func TestEnvAsPointer(t *testing.T) {
	name := "Stig"
	c := Car{
		Driver: &Driver{
			Name: name,
			Age:  999,
		},
	}

	d := NewProperty("Driver", nil)
	n := NewProperty("Name", d)

	ex := NewExpression(n)

	res, err := ex.Evaluate(&c)
	if err != nil {
		t.Fatal(err)
	}
	if res.(string) != "Stig" {
		t.Fatalf("Expected Stig, got %s", res)
	}
}
