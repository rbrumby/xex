package xex

import (
	"errors"
	"fmt"
	"reflect"
)

func init() {
	functions = make(map[string]*Function)
	registerCoreBuiltins()
	registerNumberBuiltins()
	registerStringBuiltins()
}

type node struct {
	name        string
	kinds       []reflect.Kind
	description string
	value       interface{}
}

//Function represents a xex function which can be dynamically invoked.
type Function struct {
	name          string
	documentation string
	impl          interface{}
}

func NewFunction(name string, documentation string, implementation interface{}) *Function {
	if name == "" {
		panic(errors.New("Attempt to create unnamed function"))
	}
	f := Function{
		name:          name,
		documentation: documentation,
		impl:          implementation,
	}
	if err := f.validate(); err != nil {
		panic(err)
	}
	return &f
}

func (f *Function) Name() string {
	return f.name
}

func (f *Function) Documentation() string {
	return f.documentation
}

//validate validates that the Function implementation
func (f *Function) validate() (err error) {
	if f.Name() == "" {
		err = fmt.Errorf("Attempt to use unnamed function")
		return
	}
	if f.impl == nil || reflect.TypeOf(f.impl).Kind() != reflect.Func {
		err = fmt.Errorf("Implementation of %q is not a Go function", f.Name())
		return
	}
	return
}

//Exec executes the function implementation.
//It maps the input arguments to the implemtation arguments, returning an error if the number or types do not match.
//The implementation return values are returned in results except for error.
//If the implementaion's last return value is an error, it will be returned as the error returned from Exec (it will not be included in the results slice).
//This way, error can be consistently checked whether the function cannot be called or if the functions implementation returns an error
//(in both cases, this is reported by the error return value).
func (f *Function) Exec(args ...interface{}) (results []interface{}, err error) {
	//defer recovers from reflect panicking in reflect.ValueOf(...).Call(...) returning an error if,
	//for example arguments do not match the function implementation
	defer func() {
		if recv := recover(); recv != nil {
			err = fmt.Errorf("Error executing %q: %v", f.Name(), recv)
		}
	}()

	if err = f.validate(); err != nil {
		return
	}

	vargs := make([]reflect.Value, len(args))
	for i, a := range args {
		vargs[i] = reflect.ValueOf(a)
	}

	vres := reflect.ValueOf(f.impl).Call(vargs)

	//Pick the error out of the result slice if the last arg is an error.
	//Errors are returned separately from the slice of values returned.
	switch e := vres[len(vres)-1].Interface().(type) {
	case error:
		results = make([]interface{}, len(vres)-1)
		for i, r := range vres[:len(results)] {
			results[i] = r.Interface()
		}
		err = e
	default:
		results = make([]interface{}, len(vres))
		for i, r := range vres {
			results[i] = r.Interface()
		}
	}

	return
}

var functions map[string]*Function

//RegisterFunction registers the functions name in a map so it can be obtained by name in an expression.
//It will panic if the Function is not valid, or if a Function with that name is already registered.
func RegisterFunction(f *Function) {
	if err := f.validate(); err != nil {
		panic(errors.New("Attempt to register unnamed or unimplemented function - a function must have a name & an implementation"))
	}
	if _, ok := functions[f.Name()]; ok {
		panic(fmt.Errorf("Function %q is already registered", f.Name()))
	}
	functions[f.Name()] = f
}

//GetFunction returns the named Function from the registry or returns an error if the name does not exist.
func GetFunction(name string) (*Function, error) {
	if f, ok := functions[name]; ok {
		return f, nil
	}
	return &Function{}, fmt.Errorf("Function %s does not exist", name)
}
