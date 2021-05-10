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

func NewExpression(root Node) *Expression {
	return &Expression{root}
}

func (e *Expression) Evaluate(object interface{}) (result interface{}, error error) {
	return e.root.Evaluate(object)
}

//FunctionCall is a Node in the compiled expression tree which represents a call to a funtion with Nodes as its arguments.
type FunctionCall struct {
	function  *Function
	arguments []Node
	index     int
}

func NewFunctionCall(function *Function, arguments []Node, index int) *FunctionCall {
	return &FunctionCall{function, arguments, index}
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
			return nil, fmt.Errorf("function %q: %s", fc.function.Name(), err)
		}
		args[i] = arg
	}
	results, err := fc.function.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("function %q: %s", fc.function.Name(), err)
	}
	return results[fc.Index()], nil
}

//Literal is a Node in the compiled expression tree which represents a literal value.
type Literal struct {
	value interface{}
}

func NewLiteral(value interface{}) *Literal {
	return &Literal{value}
}

func (l *Literal) Name() string {
	return "<literal>"
}

func (l *Literal) Evaluate(env interface{}) (interface{}, error) {
	return l.value, nil
}

//MethodCall is a Node in the compiled expression tree which represents a call to a method on a parent object.
type MethodCall struct {
	name      string
	parent    Node
	arguments []Node
	index     int
}

func NewMethodCall(name string, parent Node, arguments []Node, index int) *MethodCall {
	return &MethodCall{name, parent, arguments, index}
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
	args := make([]reflect.Value, len(mc.arguments))
	for i, argNode := range mc.arguments {
		arg, err := argNode.Evaluate(env)
		if err != nil {
			return nil, fmt.Errorf("method %q: %s", mc.name, err)
		}
		args[i] = reflect.ValueOf(arg)
	}

	//Assume no parent so set parent to top-level env
	parent := env
	if mc.parent != nil {
		//A parent Node is configured on this methodCall (the method is not on the top level env object)
		//Evaluate the parent Node & use its result in place of the top level env.
		parent, err = mc.parent.Evaluate(env)
		if err != nil {
			return nil, fmt.Errorf("method %s: %s", mc.name, err)
		}
	}
	meth := reflect.ValueOf(parent).MethodByName(mc.name)
	if !meth.IsValid() {
		//The method isn't valid - maybe the method has a pointer receiver?
		//If parent isn't already a pointer, try creating one & getting the method on it.
		if reflect.ValueOf(parent).Kind() != reflect.Ptr {
			parentVal := reflect.ValueOf(parent)
			ptr := reflect.New(parentVal.Type())
			ptr.Elem().Set(parentVal)
			meth = ptr.MethodByName(mc.name)
		}
		if !meth.IsValid() {
			//The method still isn't valid after trying a pointer receiver
			if mc.parent == nil {
				err = fmt.Errorf("top level object does not have method %q", mc.name)
			} else {
				err = fmt.Errorf("%s does not have method %q", mc.parent.Name(), mc.name)
			}
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

func NewProperty(name string, parent Node) *Property {
	return &Property{name, parent}
}

func (p *Property) Name() string {
	return p.name
}

//Evaluate will evaluate the chain of parent nodes if parent is not null.
//If parent is null it will evaluate the property from the env object.
func (p *Property) Evaluate(env interface{}) (result interface{}, err error) {
	if p.parent == nil {
		//No parent, this is a top-level property reference - use env
		return p.evaluate(env)
	}
	res, err := p.parent.Evaluate(env)
	if err != nil {
		return nil, fmt.Errorf("error evaluating parent of %q: %s", p.Name(), err)
	}
	return p.evaluate(res)
}

//evaluate does the real evaluation dereferencing pointers along the way.
func (p *Property) evaluate(env interface{}) (result interface{}, err error) {
	var propVal reflect.Value
	if reflect.ValueOf(env).Kind() == reflect.Ptr {
		//use the dereferenced value
		propVal = reflect.ValueOf(reflect.ValueOf(env).Elem().Interface()).FieldByName(p.Name())
	} else {
		propVal = reflect.ValueOf(env).FieldByName(p.Name())
	}
	if !propVal.IsValid() {
		return nil, fmt.Errorf("property %q not found", p.Name())
	}
	result = propVal.Interface()
	return

}
