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
		t.Fatal(err)
	}
	eq, err := GetFunction("equals")
	if err != nil {
		t.Fatal(err)
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
	sel, err := GetFunction("select")
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
		t.Fatal(err)
	}

	if r, ok := res.(map[string]Object); ok {
		if len(r) != 1 || r["1"].Value != "One" {
			t.Fatalf("Unexpected result %v", r)
		}
	} else {
		t.Fatal("Did not get a map[string]Object")
	}

}

func TestSelectInvalidType(t *testing.T) {
	sel, err := GetFunction("select")
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(errors.New("Should have failed with invalid type for select"))
	}

}

func TestSimpleArrayCount(t *testing.T) {
	cnt, err := GetFunction("count")
	if err != nil {
		t.Fatal(err)
	}

	arr := [5]string{"a", "b", "c", "d", "e"}

	fc := NewFunctionCall(
		cnt,
		[]Node{NewProperty("env", nil)},
		0,
	)
	res, err := fc.Evaluate(Values{"env": arr})
	if err != nil {
		t.Fatal(err)
	}
	if res != 5 {
		t.Fatal("array count should be 5")
	}
}

func TestFilterArrayCount(t *testing.T) {

	sel, err := GetFunction("select")
	if err != nil {
		t.Fatal(err)
	}

	cnt, err := GetFunction("count")
	if err != nil {
		t.Fatal(err)
	}

	eq, err := GetFunction("equals")
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}

	if i, ok := res.(int); !ok || i != 3 {
		t.Fatalf("Expected 3 entries from count slice with sub-select, got %d", i)
	}
}
func TestFilterMapCount(t *testing.T) {
	sel, err := GetFunction("select")
	if err != nil {
		t.Fatal(err)
	}

	cnt, err := GetFunction("count")
	if err != nil {
		t.Fatal(err)
	}

	eq, err := GetFunction("equals")
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}
	if res != 2 {
		t.Fatalf("map count should be 2 - got %d", res)
	}
}

func TestSlice(t *testing.T) {
	s, err := GetFunction("slice")
	if err != nil {
		t.Fatal(err)
	}
	res, err := s.Exec("One", "Two", "Three")
	if err != nil {
		t.Fatal(err)
	}
	if r, ok := res[0].([]string); ok {
		if len(r) != 3 {
			t.Fatalf("Expected 3 - got %d: %v", len(res), res)
		}
	} else {
		t.Fatalf("Could not assert res[0] as []string - %v", res[0])
	}

	resi, err := s.Exec(11, 12, 13, 14)
	if err != nil {
		t.Fatal(err)
	}
	if r, ok := resi[0].([]int); ok {
		if len(r) != 4 {
			t.Fatalf("Expected 4 - got %d: %v", len(r), r)
		}
	} else {
		t.Fatalf("Could not assert res[0] as []string - %v", resi[0])
	}
}
func TestBadSlice(t *testing.T) {
	s, err := GetFunction("slice")
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.Exec(11, 12, "13")
	logger.Debug(err)
	if err == nil {
		t.Fatal("slice creation should fail due to inconsistent types")
	}

}

func TestMap(t *testing.T) {
	s, err := GetFunction("map")
	if err != nil {
		t.Fatal(err)
	}
	res, err := s.Exec(MapEntry{"One", 1}, MapEntry{"Two", 2}, MapEntry{"Three", 3})
	if err != nil {
		t.Fatal(err)
	}
	if r, ok := res[0].(map[string]int); ok {
		if len(r) != 3 {
			t.Fatalf("Expected 3 - got %d: %v", len(r), r)
		}
	} else {
		t.Fatalf("Could not assert res[0] as []string - %v", res[0])
	}

	resi, err := s.Exec(MapEntry{0, 200}, MapEntry{1, 401}, MapEntry{2, 500})
	if err != nil {
		t.Fatal(err)
	}
	if r, ok := resi[0].(map[int]int); ok {
		if r[1] != 401 {
			t.Fatalf("Expected 401 - got %d: %v", r[1], r)
		}
	} else {
		t.Fatalf("Could not assert resi[0] as map[int]int - %v", resi[0])
	}
}

func TestBadMap(t *testing.T) {
	s, err := GetFunction("map")
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.Exec(MapEntry{"One", 1}, MapEntry{"Two", "2"}, MapEntry{"Three", 3})
	logger.Debug(err)
	if err == nil {
		t.Fatal("map creation should fail due to inconsistent types")
	}
}
