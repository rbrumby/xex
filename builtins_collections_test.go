package xex

import (
	"errors"
	"testing"
)

type Object struct {
	Value string
}

func TestSelectArray(t *testing.T) {
	sel, err := GetFunction("select")
	if err != nil {
		t.Error(err)
		return
	}
	eq, err := GetFunction("equals")
	if err != nil {
		t.Error(err)
		return
	}
	var arr [5]Object
	arr[0] = Object{"9"}
	arr[1] = Object{"1"}
	arr[2] = Object{"2"}
	arr[3] = Object{"1"}
	arr[4] = Object{"1"}
	fc := NewFunctionCall(
		sel,
		[]Node{
			NewProperty("array", nil),
			NewLiteral("entry"),
			NewExpression(
				NewFunctionCall(eq, []Node{NewProperty("Value", NewProperty("entry", nil)), NewLiteral("1")}, 0),
			),
		},
		0,
	)
	res, err := fc.Evaluate(Values{"array": arr})
	if err != nil {
		t.Error(err)
		return
	}
	if r, ok := res.([]Object); ok {
		if len(r) != 3 {
			t.Errorf("Expected 3 results, got %d", len(r))
			return
		}
	} else {
		t.Error("Did not get a slice of Object")
		return
	}
}

func TestSelectSlice(t *testing.T) {

	sel, err := GetFunction("select")
	if err != nil {
		t.Error(err)
		return
	}
	eq, err := GetFunction("equals")
	if err != nil {
		t.Error(err)
		return
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
			NewProperty("list", nil),
			NewLiteral("$"),
			NewExpression(
				NewFunctionCall(eq, []Node{NewProperty("Value", NewProperty("$", nil)), NewLiteral("1")}, 0),
			),
		},
		0,
	)
	res, err := fc.Evaluate(Values{"list": s})
	if err != nil {
		t.Error(err)
		return
	}

	if r, ok := res.([]Object); ok {
		if len(r) != 3 {
			t.Errorf("Expected 3 results, got %d", len(r))
			return
		}
	} else {
		t.Error("Did not get a slice of Object")
		return
	}
}

func TestSelectMap(t *testing.T) {
	sel, err := GetFunction("select")
	if err != nil {
		t.Error(err)
		return
	}

	eq, err := GetFunction("equals")
	if err != nil {
		t.Error(err)
		return
	}
	m := make(map[string]Object, 3)
	m["0"] = Object{"Zero"}
	m["1"] = Object{"One"}
	m["2"] = Object{"Two"}
	fc := NewFunctionCall(
		sel,
		[]Node{
			NewProperty("listOfValues", nil),
			NewLiteral("$entry"),
			NewExpression(
				NewFunctionCall(eq, []Node{NewProperty("Value", NewProperty("$entry", nil)), NewProperty("$0", nil)}, 0),
			),
			NewProperty("filter", nil),
		},
		0,
	)
	res, err := fc.Evaluate(Values{"listOfValues": m, "filter": "One"})
	if err != nil {
		t.Error(err)
		return
	}

	if r, ok := res.(map[string]Object); ok {
		if len(r) != 1 || r["1"].Value != "One" {
			t.Errorf("Unexpected result %v", r)
			return
		}
	} else {
		t.Error("Did not get a map[string]Object")
		return
	}

}

func TestSelectInvalidType(t *testing.T) {
	sel, err := GetFunction("select")
	if err != nil {
		t.Error(err)
		return
	}
	fc := NewFunctionCall(
		sel,
		[]Node{
			NewProperty("env", nil),
			NewLiteral("$"),
			NewExpression(nil),
		},
		0,
	)
	_, err = fc.Evaluate(Values{"env": "I'm no collection"})
	logger.Debug(err)
	if err == nil {
		t.Error(errors.New("Should have failed with invalid type for select"))
		return
	}

}

func TestSimpleArrayCount(t *testing.T) {
	cnt, err := GetFunction("count")
	if err != nil {
		t.Error(err)
		return
	}

	arr := [5]string{"a", "b", "c", "d", "e"}

	fc := NewFunctionCall(
		cnt,
		[]Node{NewProperty("env", nil)},
		0,
	)
	res, err := fc.Evaluate(Values{"env": arr})
	if err != nil {
		t.Error(err)
		return
	}
	if res != 5 {
		t.Error("array count should be 5")
		return
	}
}

func TestFilterArrayCount(t *testing.T) {

	sel, err := GetFunction("select")
	if err != nil {
		t.Error(err)
		return
	}

	cnt, err := GetFunction("count")
	if err != nil {
		t.Error(err)
		return
	}

	eq, err := GetFunction("equals")
	if err != nil {
		t.Error(err)
		return
	}

	arr := [5]string{"c", "a", "c", "b", "c"}

	fc := NewFunctionCall(
		cnt,
		[]Node{
			NewFunctionCall(
				sel,
				[]Node{
					NewProperty("env", nil),
					NewLiteral("$"),
					NewExpression(
						NewFunctionCall(
							eq,
							[]Node{
								NewProperty("$", nil),
								NewLiteral("c"),
							},
							0,
						),
					),
				},
				0,
			),
		},
		0,
	)

	res, err := fc.Evaluate(Values{"env": arr})
	if err != nil {
		t.Error(err)
		return
	}

	if i, ok := res.(int); !ok || i != 3 {
		t.Errorf("Expected 3 entries from count slice with sub-select, got %d", i)
		return
	}
}
func TestFilterMapCount(t *testing.T) {
	sel, err := GetFunction("select")
	if err != nil {
		t.Error(err)
		return
	}

	cnt, err := GetFunction("count")
	if err != nil {
		t.Error(err)
		return
	}

	eq, err := GetFunction("equals")
	if err != nil {
		t.Error(err)
		return
	}

	mp := map[string]int{"i": 0, "j": 1, "l": 2, "m": 3, "n": 3}

	fc := NewFunctionCall(
		cnt,
		[]Node{
			NewFunctionCall(
				sel,
				[]Node{
					NewProperty("env", nil),
					NewLiteral("$"),
					NewExpression(
						NewFunctionCall(
							eq,
							[]Node{
								NewProperty("$", nil),
								NewLiteral(3),
							},
							0,
						),
					),
				},
				0,
			),
		},
		0,
	)

	res, err := fc.Evaluate(Values{"env": mp})
	if err != nil {
		t.Error(err)
		return
	}
	if res != 2 {
		t.Errorf("map count should be 2 - got %d", res)
		return
	}
}

func TestSlice(t *testing.T) {
	s, err := GetFunction("slice")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := s.Exec("One", "Two", "Three")
	if err != nil {
		t.Error(err)
		return
	}
	if r, ok := res[0].([]string); ok {
		if len(r) != 3 {
			t.Errorf("Expected 3 - got %d: %v", len(res), res)
			return
		}
	} else {
		t.Errorf("Could not assert res[0] as []string - %v", res[0])
		return
	}

	resi, err := s.Exec(11, 12, 13, 14)
	if err != nil {
		t.Error(err)
		return
	}
	if r, ok := resi[0].([]int); ok {
		if len(r) != 4 {
			t.Errorf("Expected 4 - got %d: %v", len(r), r)
			return
		}
	} else {
		t.Errorf("Could not assert res[0] as []string - %v", resi[0])
		return
	}
}
func TestBadSlice(t *testing.T) {
	s, err := GetFunction("slice")
	if err != nil {
		t.Error(err)
		return
	}
	_, err = s.Exec(11, 12, "13")
	logger.Debug(err)
	if err == nil {
		t.Error("slice creation should fail due to inconsistent types")
		return
	}

}

func TestMap(t *testing.T) {
	s, err := GetFunction("map")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := s.Exec(MapEntry{"One", 1}, MapEntry{"Two", 2}, MapEntry{"Three", 3})
	if err != nil {
		t.Error(err)
		return
	}
	if r, ok := res[0].(map[string]int); ok {
		if len(r) != 3 {
			t.Errorf("Expected 3 - got %d: %v", len(r), r)
			return
		}
	} else {
		t.Errorf("Could not assert res[0] as []string - %v", res[0])
		return
	}

	resi, err := s.Exec(MapEntry{0, 200}, MapEntry{1, 401}, MapEntry{2, 500})
	if err != nil {
		t.Error(err)
		return
	}
	if r, ok := resi[0].(map[int]int); ok {
		if r[1] != 401 {
			t.Errorf("Expected 401 - got %d: %v", r[1], r)
			return
		}
	} else {
		t.Errorf("Could not assert resi[0] as map[int]int - %v", resi[0])
		return
	}
}

func TestBadMap(t *testing.T) {
	s, err := GetFunction("map")
	if err != nil {
		t.Error(err)
		return
	}
	_, err = s.Exec(MapEntry{"One", 1}, MapEntry{"Two", "2"}, MapEntry{"Three", 3})
	logger.Debug(err)
	if err == nil {
		t.Error("map creation should fail due to inconsistent types")
		return
	}
}

func TestIndexOfSlice(t *testing.T) {
	s := []string{"zero", "one", "two", "three"}
	f, err := GetFunction("indexOf")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := f.Exec(s, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != "two" {
		t.Errorf("Expected two, got %s", res[0])
		return
	}
}

func TestIndexOfMap(t *testing.T) {
	s := map[int]string{0: "zero", 1: "one", 2: "two", 3: "three"}
	f, err := GetFunction("indexOf")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := f.Exec(s, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != "two" {
		t.Errorf("Expected two, got %s", res[0])
		return
	}
	res, err = f.Exec(s, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != "zero" {
		t.Errorf("Expected zero, got %s", res[0])
		return
	}
}
