package xex

import (
	"errors"
	"fmt"
	"reflect"
)

type node struct {
	name        string
	kinds       []reflect.Kind
	description string
	value       interface{}
}

//Function represents a xex function which can be dynamically invoked.
type Function struct {
	Name          string
	Documentation string
	impl          interface{}
}

//SetImplementation ensures that impl is a Go func & sets it inside the Function.
func (f *Function) SetImplementation(impl interface{}) error {
	if reflect.TypeOf(impl).Kind() != reflect.Func {
		return errors.New("Implementation must be a Go function")
	}
	f.impl = impl
	return nil
}

//Exec executes the function implementation.
//If the implementation does not return an error, error will always be nil.
func (f *Function) Exec(args ...interface{}) (results []interface{}, err error) {
	//defer recovers from errors to reflect.ValueOf(...).Call(...) returning an error if,
	//for example arguments do not match the function implementation
	defer func() {
		if recv := recover(); recv != nil {
			err = fmt.Errorf("Error executing %q: %v", f.Name, recv)
		}
	}()

	if reflect.TypeOf(f.impl).Kind() != reflect.Func {
		err = errors.New("exec must be a function")
		return
	}
	vargs := make([]reflect.Value, len(args))
	for i, a := range args {
		vargs[i] = reflect.ValueOf(a)
	}

	vres := reflect.ValueOf(f.impl).Call(vargs)

	//Pick the error out of the result slice if the last arg is an error
	if err, ok := vres[len(vres)-1].Interface().(error); ok {
		results = make([]interface{}, len(vres)-1)
		for i, r := range vres[:len(results)] {
			results[i] = r.Interface()
		}
		return results, err
	}

	//Else just send all of the results back
	results = make([]interface{}, len(vres))
	for i, r := range vres {
		results[i] = r.Interface()
	}
	return
}

type operator string

const (
	add      operator = "+"
	subtract          = "-"
	multiply          = "*"
	divide            = "/"
	modulus           = "%"
)

type comparator string

const (
	equals    comparator = "=="
	notEquals            = "!="
)
