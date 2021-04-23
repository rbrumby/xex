package xex

import (
	"testing"
)

func TestAdd(t *testing.T) {
	fn, err := GetFunction("addInt")
	if err != nil {
		t.Error(err)
		return
	}
	sum, err := fn.Exec(4, 5)
	if err != nil {
		t.Error(err)
		return
	}
	if sum[0] != 9 {
		t.Errorf("Expected 9, got %d", sum[0])
	}
}

func TestConcat(t *testing.T) {
	fn, err := GetFunction("concat")
	if err != nil {
		t.Error(err)
		return
	}
	con, err := fn.Exec("Hello ", "world!")
	if err != nil {
		t.Error(err)
		return
	}
	if con[0] != "Hello world!" {
		t.Errorf(`Expected "Hello word!", got %q`, con)
	}
}
