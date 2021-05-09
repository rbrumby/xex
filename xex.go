package xex

import (
	"fmt"
	"reflect"
)

//Node is a node in the compiled expression tree.
type Node interface {
	Name() string
	Evaluate(object interface{}) (interface{}, error)
}

type Expression struct {
	root Node
}

func (e *Expression) Evaluate(object interface{}) (result interface{}, error error) {
	return e.root.Evaluate(object)
}

//FunctionCall is a Node in the compiled expression tree which represents a call to a funtion with Nodes as its arguments.
type FunctionCall struct {
	function  *Function
	index     int
	arguments []Node
}

func (fc *FunctionCall) Name() string {
	return fc.function.name
}

func (fc *FunctionCall) Index() int {
	return fc.index
}

func (fc FunctionCall) Evaluate(object interface{}) (interface{}, error) {
	args := make([]interface{}, len(fc.arguments))
	for i, argNode := range fc.arguments {
		arg, err := argNode.Evaluate(object)
		if err != nil {
			return nil, fmt.Errorf("function %s: %s", fc.function.Name(), err)
		}
		args[i] = arg
	}
	results, err := fc.function.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("function %s: %s", fc.function.Name(), err)
	}
	return results[fc.Index()], nil
}

//Literal is a Node in the compiled expression tree which represents a literal value.
type Literal struct {
	value interface{}
}

func (l *Literal) Name() string {
	return "literal"
}

func (l *Literal) Evaluate(env interface{}) (interface{}, error) {
	return l.value, nil
}

//MethodCall is a Node in the compiled expression tree which represents a call to a method on a parent object.
type MethodCall struct {
	name      string
	parent    Node
	index     int
	arguments []Node
}

func (mc *MethodCall) SetParent(parent Node) {
	mc.parent = parent
}

func (mc *MethodCall) Name() string {
	return mc.name
}

func (mc *MethodCall) Index() int {
	return mc.index
}

//Evaluate calls the method on the MethodCalls parent or a pointer to the MethodCalls parent if the method isn't found on the parent itself.
//It will call Evaluate on the parent & the arguments passed to the MethodCall before invoking the underlying method.
func (mc *MethodCall) Evaluate(env interface{}) (result interface{}, err error) {
	defer func() {
		if recv := recover(); recv != nil {
			err = fmt.Errorf("method %s: %s", mc.Name(), recv)
		}
	}()

	args := make([]reflect.Value, len(mc.arguments))
	for i, argNode := range mc.arguments {
		arg, err := argNode.Evaluate(env)
		if err != nil {
			return nil, fmt.Errorf("method %s: %s", mc.name, err)
		}
		args[i] = reflect.ValueOf(arg)
	}

	if mc.parent == nil {
		return nil, fmt.Errorf("method %s: method has null parent", mc.name)
	}
	parent, err := mc.parent.Evaluate(env)
	if err != nil {
		return nil, fmt.Errorf("method %s: %s", mc.name, err)
	}
	meth := reflect.ValueOf(parent).MethodByName(mc.name)
	if !meth.IsValid() {
		//If parent isn't already a pointer, try creating one & getting the method on it.
		//This handles situations where the method has a pointer receiver but the object in the graph is not a pointer.
		//Probably bad code in the first place but we can handle it so why not!
		if reflect.ValueOf(parent).Kind() != reflect.Ptr {
			parentVal := reflect.ValueOf(parent)
			ptr := reflect.New(parentVal.Type())
			ptr.Elem().Set(parentVal)
			meth = ptr.MethodByName(mc.name)
		}
		if !meth.IsValid() {
			err = fmt.Errorf("%s does not have method %s", mc.parent.Name(), mc.name)
			return
		}
	}
	results := meth.Call(args)
	//If last result is an error, split it from the result slice & return as a separate error.
	if errchk, ok := results[len(results)-1].Interface().(error); ok {
		err = errchk
	}
	result = results[mc.Index()].Interface()
	return
}

//Property is a Node in the compiled expression tree which represents a reference to a property.
type Property struct {
	name   string
	parent Node
}

func (p *Property) SetParent(parent Node) {
	p.parent = parent
}

func (p *Property) Name() string {
	return p.name
}

//Evaluate will evaluate the chain of parent nodes if parent is not null.
//If parent is null it will evaluate the proprty from the env object.
func (p *Property) Evaluate(env interface{}) (result interface{}, err error) {
	defer func() {
		if recv := recover(); recv != nil {
			err = fmt.Errorf("property %s: %s", p.Name(), recv)
		}
	}()
	if p.parent == nil {
		//No parent, this is a top-level property reference - use env
		field := reflect.ValueOf(env).FieldByName(p.Name())
		if (field == reflect.Value{}) {
			return nil, fmt.Errorf("property %s: not found", p.Name())
		}
		result = field.Interface()
		return
	}
	//Evaluate parent & then extract the property from the result.
	parent, err := p.parent.Evaluate(env)
	if err != nil {
		return nil, fmt.Errorf("property %s: %s", p.Name(), err)
	}
	field := reflect.ValueOf(parent).FieldByName(p.Name())
	if (field == reflect.Value{}) {
		return nil, fmt.Errorf("property %s: not found", p.Name())
	}
	result = field.Interface()
	return
}
