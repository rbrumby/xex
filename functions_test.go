package xex

import (
	"errors"
	"testing"
)

func TestHappyPathRegisterGetAndExec(t *testing.T) {
	RegisterFunction(
		NewFunction("test", FunctionDocumentation{Text: "just a test", Parameters: map[string]string{"test1": "param1"}}, testFunc),
	)
	f, err := GetFunction("test")
	if err != nil {
		t.Error(err)
		return
	}
	if f.Name() != "test" || f.DocumentationString() != "just a test\ntest1: param1" {
		t.Errorf("unexpected function documentation: %s - %s", f.Name(), f.DocumentationString())
		return
	}
	values, err := f.Exec("XXX")
	if err != nil {
		t.Error(err)
		return
	}
	if values[0] != "Hello world!" {
		t.Errorf("Expected \"Hello world!\", got %q", values[0])
		return
	}
}

func TestHappyPathRegisterGetAndExecNoError(t *testing.T) {
	RegisterFunction(
		NewFunction("testnoerr", FunctionDocumentation{Text: "just another test"}, testFuncNoError),
	)
	f, err := GetFunction("testnoerr")
	if err != nil {
		t.Error(err)
		return
	}
	values, err := f.Exec()
	if err != nil {
		t.Error(err)
		return
	}
	if values[0] != "Hello world!" {
		t.Errorf("Expected \"Helloworld!\", got %q", values[0])
		return
	}
}

func TestGetNonExistentFunction(t *testing.T) {
	_, err := GetFunction("xxx")
	if err == nil {
		t.Error("should have got error sayign function does not exist")
		return
	}
}

type myCustomError struct {
	msg string
}

func (e myCustomError) Error() string {
	return e.msg
}

func TestCustomErrorTypeFunction(t *testing.T) {
	fn := NewFunction("fail", FunctionDocumentation{Text: "always returns an error"}, func() (string, myCustomError) {
		return "nothing", myCustomError{msg: "It always fails"}
	})
	_, err := fn.Exec()
	if err == nil || err.Error() != "It always fails" {
		t.Errorf("Didn't get expected error. Got %q", err.Error())
		return
	}
}

func TestUnnamedFunction(t *testing.T) {
	defer assertPanic(t)
	_ = NewFunction("", FunctionDocumentation{Text: "just a test"}, testFunc)
}

func TestInvalidFunctionName(t *testing.T) {
	defer assertPanic(t)
	_ = NewFunction("xy1_2+x", FunctionDocumentation{Text: "just a test"}, testFunc)
}

func TestInvalidFuncNameregex(t *testing.T) {
	defer assertPanic(t)
	fn := NewFunction("valid", FunctionDocumentation{Text: "just a test"}, testFunc)
	fn.name = "in-valid"
	fn.validate("regexpisinvalid(")
}

func TestDuplicateFunction(t *testing.T) {
	defer assertPanic(t)
	RegisterFunction(NewFunction("duplicate", FunctionDocumentation{Text: "just a test"}, testFunc))
	RegisterFunction(NewFunction("duplicate", FunctionDocumentation{Text: "just a test"}, testFunc))
}

func TestNilFunctionImplementation(t *testing.T) {
	defer assertPanic(t)
	_ = NewFunction("test", FunctionDocumentation{Text: "just a test"}, nil)
}

func TestRegisterInternallyBrokenFunction(t *testing.T) {
	defer assertPanic(t)
	f := NewFunction("test", FunctionDocumentation{Text: "just a test"}, testFuncNoError)
	f.name = ""
	RegisterFunction(f)
}

func TestCallInternallyBrokenFunction(t *testing.T) {
	f := NewFunction("test", FunctionDocumentation{Text: "just a test"}, testFuncNoError)
	f.name = ""
	_, err := f.Exec()
	if err == nil {
		t.Error(err)
		return
	}
}

func TestIncorrectArgs(t *testing.T) {
	f := NewFunction("test", FunctionDocumentation{Text: "just a test"}, testFunc)
	_, err := f.Exec(5)
	if err == nil {
		t.Error("Should have failed using int as string")
		return
	}
	_, err = f.Exec("x", "y")
	if err == nil {
		t.Error("Should have failed with too many input arguments")
		return
	}

}

func assertPanic(t *testing.T) {
	err := recover()
	if err == nil {
		t.Error("Should have panicked")
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
