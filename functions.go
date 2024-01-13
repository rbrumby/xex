package xex

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var functions map[string]*Function

const FuncNameRegex = "^[a-z][a-zA-Z0-9_]*$"

func init() {
	functions = make(map[string]*Function)
	registerCoreBuiltins()
	registerNumberBuiltins()
	registerStringBuiltins()
	registerCollectionBuiltins()
}

type FunctionDocParam struct {
	Name        string
	Description string
}

type FunctionDocumentation struct {
	Text       string
	Parameters []FunctionDocParam
}

//Function represents a xex function which can be dynamically invoked.
type Function struct {
	Name          string
	Documentation FunctionDocumentation
	impl          interface{}
}

func GetFunctionNames() (names []string) {
	for n := range functions {
		names = append(names, n)
	}
	return
}

//NewFunction returns a pointer to a new Function.
func NewFunction(name string, documentation FunctionDocumentation, implementation interface{}) *Function {
	f := Function{
		Name:          name,
		Documentation: documentation,
		impl:          implementation,
	}
	if err := f.validate(FuncNameRegex); err != nil {
		panic(err)
	}
	return &f
}

func (f *Function) DocumentationString() string {
	out := strings.Builder{}
	out.WriteString(f.Documentation.Text)
	for _, p := range f.Documentation.Parameters {
		out.WriteRune('\n')
		out.WriteString(p.Name + ": " + p.Description)
	}
	return out.String()
}

//validate validates that the Function implementation
func (f *Function) validate(fNameRegex string) (err error) {
	if f.Name == "" {
		err = fmt.Errorf("attempt to use unnamed function")
		return
	}
	if ok, err := regexp.MatchString(fNameRegex, f.Name); !ok {
		if err != nil {
			panic(fmt.Errorf("error applying regexp %q to %q: %s", fNameRegex, f.Name, err))
		}
		panic(fmt.Errorf("invalid function name %q: function names must match regular expression %q", f.Name, fNameRegex))
	}
	if f.impl == nil || reflect.TypeOf(f.impl).Kind() != reflect.Func {
		err = fmt.Errorf("implementation of %q is not a Go function", f.Name)
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
			logger.Debugf("recovering from call to %q with args %v: %s", f.Name, args, recv)
			err = fmt.Errorf("error executing %q with args %v: %v", f.Name, args, recv)
		}
	}()

	if err = f.validate(FuncNameRegex); err != nil {
		return
	}

	vargs := make([]reflect.Value, len(args))
	for i, a := range args {
		if a == nil {
			vargs[i] = reflect.New(reflect.TypeOf(f.impl).In(i)).Elem()
			continue
		}
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

//RegisterFunction registers the functions name in a map so it can be obtained by name in an expression.
//It will panic if the Function is not valid, or if a Function with that name is already registered.
func RegisterFunction(f *Function) {
	if err := f.validate(FuncNameRegex); err != nil {
		panic(errors.New("attempt to register unnamed or unimplemented function - a function must have a name & an implementation"))
	}
	if _, ok := functions[f.Name]; ok {
		panic(fmt.Errorf("function %q is already registered", f.Name))
	}
	functions[f.Name] = f
}

//GetFunction returns the named Function from the registry or returns an error if the name does not exist.
func GetFunction(name string) (*Function, error) {
	if f, ok := functions[name]; ok {
		return f, nil
	}
	return &Function{}, fmt.Errorf("function %q does not exist", name)
}
