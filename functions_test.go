package xex

import (
	"errors"
	"testing"
)

func TestHappyPathRegisterGetAndExec(t *testing.T) {
	RegisterFunction(
		NewFunction("test", "just a test", testFunc),
	)
	f, err := GetFunction("test")
	if err != nil {
		t.Fatal(err)
		return
	}
	if f.Name() != "test" || f.Documentation() != "just a test" {
		t.Fatal("Function didn't contain expected values")
	}
	values, err := f.Exec("XXX")
	if err != nil {
		t.Fatal(err)
		return
	}
	if values[0] != "Hello world!" {
		t.Fatalf("Expected \"Hello world!\", got %q", values[0])
	}
}

func TestHappyPathRegisterGetAndExecNoError(t *testing.T) {
	RegisterFunction(
		NewFunction("testnoerr", "just another test", testFuncNoError),
	)
	f, err := GetFunction("testnoerr")
	if err != nil {
		t.Fatal(err)
		return
	}
	values, err := f.Exec()
	if err != nil {
		t.Fatal(err)
		return
	}
	if values[0] != "Hello world!" {
		t.Fatalf("Expected \"Helloworld!\", got %q", values[0])
	}
}

func TestGetNonExistentFunction(t *testing.T) {
	_, err := GetFunction("xxx")
	if err == nil {
		t.Fatal("sould have got error sayign function does not exist")
	}
}

type myCustomError struct {
	msg string
}

func (e myCustomError) Error() string {
	return e.msg
}

func TestCustomErrorTypeFunction(t *testing.T) {
	fn := NewFunction("fail", "always returns an error", func() (string, myCustomError) {
		return "nothing", myCustomError{msg: "It always fails"}
	})
	_, err := fn.Exec()
	if err == nil || err.Error() != "It always fails" {
		t.Fatalf("Didn't get expected error. Got %q", err.Error())
	}
}

func TestUnnamedFunction(t *testing.T) {
	defer assertPanic(t)
	_ = NewFunction("", "just a test", testFunc)
}

func TestDuplicateFunction(t *testing.T) {
	defer assertPanic(t)
	RegisterFunction(NewFunction("duplicate", "just a test", testFunc))
	RegisterFunction(NewFunction("duplicate", "just a test", testFunc))
}

func TestNilFunctionImplementation(t *testing.T) {
	defer assertPanic(t)
	_ = NewFunction("test", "just a test", nil)
}

func TestRegisterInternallyBrokenFunction(t *testing.T) {
	defer assertPanic(t)
	f := NewFunction("test", "just a test", testFuncNoError)
	f.name = ""
	RegisterFunction(f)
}

func TestCallInternallyBrokenFunction(t *testing.T) {
	f := NewFunction("test", "just a test", testFuncNoError)
	f.name = ""
	_, err := f.Exec()
	if err == nil {
		t.Fatal(err)
		return
	}
}

func TestIncorrectArgs(t *testing.T) {
	f := NewFunction("test", "just a test", testFunc)
	_, err := f.Exec(5)
	if err == nil {
		t.Fatal("Should have failed using int as string")
		return
	}
	_, err = f.Exec("x", "y")
	if err == nil {
		t.Fatal("Should have failed with too many input arguments")
		return
	}

}

func assertPanic(t *testing.T) {
	err := recover()
	if err == nil {
		t.Fatal("Should have panicked")
	}

}

func testFunc(val string) (out string, err error) {
	if val == "" {
		err = errors.New("A test error")
	}
	out = "Hello world!"
	return
}

func testFuncNoError() (out string) {
	out = "Hello world!"
	return
}
