package xex

import "testing"

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

	t1 := testStruct{sub: testSubStruct{value: "AbC"}}
	t2 := testStruct{sub: testSubStruct{value: "AbC"}}
	res, err = fn.Exec(t1, t2)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != true {
		t.Fatal("values should match")
	}
}

func TestCount(t *testing.T) {
	fn, err := GetFunction("count")
	if err != nil {
		t.Fatal(err)
	}

	arr := [5]string{"a", "b", "c", "d", "e"}
	res, err := fn.Exec(arr)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != 5 {
		t.Fatal("array count should be 5")
	}

	slc := []string{"f", "g", "h"}
	res, err = fn.Exec(slc)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != 3 {
		t.Fatal("slice count should be 3")
	}

	mp := map[string]int{"i": 0, "j": 1, "l": 2, "m": 3}
	res, err = fn.Exec(mp)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != 4 {
		t.Fatal("map count should be 4")
	}

	_, err = fn.Exec("not a countable type")
	if err == nil {
		t.Fatal("should have failed to count string")
	}
}

func TestDecode(t *testing.T) {
	fn, err := GetFunction("decode")
	if err != nil {
		t.Fatal(err)
	}

	values := []interface{}{"testval", "x", "y", "testval", testSubStruct{value: "answer"}}
	res, err := fn.Exec(values)
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
	res, err = fn.Exec(values)
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
