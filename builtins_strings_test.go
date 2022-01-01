package xex

import "testing"

func TestString(t *testing.T) {
	str, err := GetFunction("string")
	if err != nil {
		t.Error(err)
		return
	}
	con, err := str.Exec(555)
	if err != nil {
		t.Error(err)
		return
	}
	if con[0] != "555" {
		t.Errorf(`Expected "555", got %q`, con[0])
		return
	}
}

func TestConcat(t *testing.T) {
	concat, err := GetFunction("concat")
	if err != nil {
		t.Error(err)
		return
	}
	con, err := concat.Exec("Hello ", "world!")
	if err != nil {
		t.Error(err)
		return
	}
	if con[0] != "Hello world!" {
		t.Errorf(`Expected "Hello world!", got %q`, con[0])
		return
	}
}

func TestLen(t *testing.T) {
	fn, err := GetFunction("len")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := fn.Exec("Hello world!")
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != 12 {
		t.Errorf("Expected 12, got %d", res[0])
		return
	}
}

func TestSubstring(t *testing.T) {
	fn, err := GetFunction("substring")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := fn.Exec("Hello world!", 0, 4)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != "Hell" {
		t.Errorf(`Expected "Hell", got %q`, res[0])
		return
	}

	res, err = fn.Exec("Hello world!", 0, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != "Hello world!" {
		t.Errorf(`Expected "Hello world!", got %q`, res[0])
		return
	}

	res, err = fn.Exec("Hello world!", 3, 5)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != "lo" {
		t.Errorf(`Expected "lo", got %q`, res[0])
		return
	}

	res, err = fn.Exec("Hello world!", 6, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != "world!" {
		t.Errorf(`Expected "world!", got %q`, res[0])
		return
	}

}

func TestInstring(t *testing.T) {
	fn, err := GetFunction("instring")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := fn.Exec("Hello world!", "orl")
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != 7 {
		t.Errorf(`Expected 7, got %d`, res[0])
		return
	}

	res, err = fn.Exec("Hello world!", "xyz")
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != -1 {
		t.Errorf(`Expected -1, got %d`, res[0])
		return
	}

}
