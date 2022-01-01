package xex

import (
	"fmt"
	"reflect"
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
	Gear  uint8
	Gears []Gear
}

type Gear struct {
	Ratio float32
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

	root := NewProperty("car", nil)
	eProp := NewProperty("Engine", root)
	mphArg := NewLiteral(30)
	rpm := NewMethodCall("GetRPM", eProp, []Node{mphArg}, 0)
	exp := NewExpression(rpm)

	res, err := exp.Evaluate(Values{"car": c})
	if err != nil {
		t.Error(err)
		return
	}
	if reflect.TypeOf(res).Kind() != reflect.Int {
		t.Error("RPM's are not an int")
		return
	}
	if res != 1125 {
		t.Errorf("Expected 1125, got %d", res)
		return
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
	_, err := mc.Evaluate(Values{"car": c})
	if err == nil {
		t.Error("expected cannot call method on inl object")
		return
	}

	e := NewProperty("Engine", NewProperty("car", nil))
	mc2 := NewMethodCall("not_exists", e, nil, 0)
	_, err = mc2.Evaluate(Values{"car": c})
	if err == nil {
		t.Error("expected method not_exists doesn't exist error")
		return
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

	root := NewProperty("car", nil)
	eProp := NewProperty("Engine", root)
	gbProp := NewProperty("Gearbox", eProp)
	gProp := NewProperty("Gear", gbProp)

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
					gProp,
				},
				0,
			),
		},
		0,
	)

	result, err := fnDivideCall.Evaluate(Values{"car": c})
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
			NewProperty("car", nil), //parent
			nil,                     //no args
			0,
		),
	)

	gear, err := gearNode.Evaluate(Values{"car": c})
	if err != nil {
		t.Error(err)
		return
	}

	if gear != uint8(4) {
		t.Errorf("Expected gear 4, got %d", gear)
		return
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

	d := NewProperty("Driver", NewProperty("car", nil))
	n := NewProperty("Name", d)

	ex := NewExpression(n)

	res, err := ex.Evaluate(Values{"car": c})
	if err != nil {
		t.Error(err)
		return
	}
	if res.(string) != "Stig" {
		t.Errorf("Expected Stig, got %s", res)
		return
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

	d := NewProperty("Driver", NewProperty("car", nil))
	n := NewProperty("Name", d)

	ex := NewExpression(n)

	res, err := ex.Evaluate(Values{"car": &c})
	if err != nil {
		t.Error(err)
		return
	}
	if res.(string) != "Stig" {
		t.Errorf("Expected Stig, got %s", res)
		return
	}
}
