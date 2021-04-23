package xex

import (
	"errors"
	"fmt"
	"strings"
)

var functions map[string]Function

func RegisterFunction(f Function) error {
	if f.Name == "" || f.impl == nil {
		return errors.New("Attempt to register unnamed or unimplemented function - a function must have a name & an implementation")
	}
	if _, ok := functions[f.Name]; ok {
		return fmt.Errorf("Function %q is already registered", f.Name)
	}
	functions[f.Name] = f
	return nil
}

func GetFunction(name string) (Function, error) {
	if f, ok := functions[name]; ok {
		return f, nil
	}
	return Function{}, fmt.Errorf("Function %s does not exist", name)
}

func init() {
	functions = make(map[string]Function)

	addInt := Function{
		Name:          "addInt",
		Documentation: `addInt ads two ints returning a single int result`,
	}
	err := addInt.SetImplementation(func(num1, num2 int) int {
		return num1 + num2
	})
	if err != nil {
		panic(err)
	}
	RegisterFunction(addInt)

	concat := Function{
		Name:          "concat",
		Documentation: `concat concatenates any number of strings returning a single string result`,
	}
	err = concat.SetImplementation(func(strs ...string) string {
		sb := strings.Builder{}
		for _, s := range strs {
			sb.WriteString(s)
		}
		return sb.String()
	})
	if err != nil {
		panic(err)
	}
	RegisterFunction(concat)
}
