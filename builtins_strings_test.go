package xex

import "testing"

func TestConcat(t *testing.T) {
	concat, err := GetFunction("concat")
	if err != nil {
		t.Fatal(err)
		return
	}
	con, err := concat.Exec("Hello ", "world!")
	if err != nil {
		t.Fatal(err)
	}
	if con[0] != "Hello world!" {
		t.Fatalf(`Expected "Hello world!", got %q`, con[0])
	}
}

func TestLen(t *testing.T) {
	fn, err := GetFunction("len")
	if err != nil {
		t.Fatal(err)
		return
	}
	res, err := fn.Exec("Hello world!")
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != 12 {
		t.Fatalf("Expected 12, got %d", res[0])
	}
}

func TestSubstring(t *testing.T) {
	fn, err := GetFunction("substring")
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := fn.Exec("Hello world!", 0, 4)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != "Hell" {
		t.Fatalf(`Expected "Hell", got %q`, res[0])
	}

	res, err = fn.Exec("Hello world!", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != "Hello world!" {
		t.Fatalf(`Expected "Hello world!", got %q`, res[0])
	}

	res, err = fn.Exec("Hello world!", 3, 5)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != "lo" {
		t.Fatalf(`Expected "lo", got %q`, res[0])
	}

	res, err = fn.Exec("Hello world!", 6, 0)
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != "world!" {
		t.Fatalf(`Expected "world!", got %q`, res[0])
	}

}

func TestInstring(t *testing.T) {
	fn, err := GetFunction("instring")
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := fn.Exec("Hello world!", "orl")
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != 7 {
		t.Fatalf(`Expected 7, got %d`, res[0])
	}

	res, err = fn.Exec("Hello world!", "xyz")
	if err != nil {
		t.Fatal(err)
	}
	if res[0] != -1 {
		t.Fatalf(`Expected -1, got %d`, res[0])
	}

}
