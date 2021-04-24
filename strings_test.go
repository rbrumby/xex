package xex

import "testing"

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
	}
}
