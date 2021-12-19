package xex

import (
	"fmt"
	"reflect"

	"github.com/coreos/capnslog"
)

var logger = capnslog.NewPackageLogger("github.com/rbrumby/xex", "xex")
var Repolog capnslog.RepoLogger

func init() {
	Repolog = capnslog.MustRepoLogger("github.com/rbrumby/xex")
}

//Node is a node in the compiled expression tree.
type Node interface {
	Name() string
	Evaluate(values Values) (interface{}, error)
	String() string
}

type Call interface {
	Node
	Arg(arg Node)
}

type Values map[string]interface{}

type Expression struct {
	root Node
}

func NewExpression(root Node) *Expression {
	return &Expression{root}
}

func (e *Expression) Name() string {
	return "<expression>"
}

func (e *Expression) Evaluate(values Values) (interface{}, error) {
	if values == nil {
		values = make(Values)
	}
	return e.root.Evaluate(values)
}

func (e *Expression) String() string {
	return fmt.Sprintf("Expression:\n%s\n", e.root)
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

func (fc *FunctionCall) Arg(arg Node) {
	if fc.arguments == nil {
		fc.arguments = make([]Node, 0)
	}
	fc.arguments = append(fc.arguments, arg)
}

func (fc *FunctionCall) Evaluate(values Values) (interface{}, error) {
	args := make([]interface{}, len(fc.arguments))
	for i, argNode := range fc.arguments {
		if reflect.TypeOf(fc.function.impl).NumIn() > i &&
			reflect.TypeOf(fc.function.impl).In(i).Kind() == reflect.Ptr &&
			reflect.TypeOf(fc.function.impl).In(i) == reflect.ValueOf(&Expression{}).Type() {
			//Arg is a pointer to an expression. Don't evaluate it, pass the expression to the function for it to evaluate.
			args[i] = argNode
			continue
		}
		arg, err := argNode.Evaluate(values)
		if err != nil {
			return nil, fmt.Errorf("function %q: %s", fc.Name(), err)
		}
		args[i] = arg
	}
	results, err := fc.function.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("function %q: %s", fc.Name(), err)
	}
	if len(results) <= fc.Index() {
		return nil, fmt.Errorf("index %d out of range. Function %s returned %d values (indexex start at zero)", fc.Index(), fc.Name(), len(results))
	}
	return results[fc.Index()], nil
}

func (f *FunctionCall) String() string {
	return fmt.Sprintf("FunctionCall: %s(%s)", f.function.name, f.arguments)
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

func (l *Literal) Evaluate(values Values) (interface{}, error) {
	return l.value, nil
}

func (l *Literal) String() string {
	return fmt.Sprintf("Literal: %s", l.value)
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

func (mc *MethodCall) Arg(arg Node) {
	if mc.arguments == nil {
		mc.arguments = make([]Node, 0)
	}
	mc.arguments = append(mc.arguments, arg)
}

func (m *MethodCall) String() string {
	return fmt.Sprintf("MethodCall: %s(%s)", m.name, m.arguments)
}

//Evaluate calls the method on the MethodCalls parent or a pointer to the MethodCalls parent if the method isn't found on the parent itself.
//It will call Evaluate on the parent & the arguments passed to the MethodCall before invoking the underlying method.
func (mc *MethodCall) Evaluate(values Values) (result interface{}, err error) {
	if mc.parent == nil {
		return nil, fmt.Errorf("cannot call method %q on nil parent", mc.Name())
	}
	args := make([]reflect.Value, len(mc.arguments))
	for i, argNode := range mc.arguments {
		arg, err := argNode.Evaluate(values)
		if err != nil {
			return nil, fmt.Errorf("method %q: %s", mc.Name(), err)
		}
		args[i] = reflect.ValueOf(arg)
	}

	//Evaluate the parent Node & execute the named method on the result.
	parent, err := mc.parent.Evaluate(values)
	if err != nil {
		return nil, fmt.Errorf("method %s: %s", mc.Name(), err)
	}
	meth := reflect.ValueOf(parent).MethodByName(mc.Name())
	if !meth.IsValid() {
		//The method isn't valid - maybe the method has a pointer receiver?
		//If parent isn't already a pointer, try creating one & getting the method on it.
		if reflect.ValueOf(parent).Kind() != reflect.Ptr {
			parentVal := reflect.ValueOf(parent)
			ptr := reflect.New(parentVal.Type())
			ptr.Elem().Set(parentVal)
			meth = ptr.MethodByName(mc.Name())
		}
		if !meth.IsValid() {
			//The method still isn't valid after trying a pointer receiver
			if mc.parent == nil {
				err = fmt.Errorf("top level object does not have method %q", mc.Name())
			} else {
				err = fmt.Errorf("value retrieved from %q does not have method %q", mc.parent.Name(), mc.Name())
			}
			return
		}
	}
	results := meth.Call(args)
	//If last result is an error, split it from the result slice & return as a separate error.
	if errchk, ok := results[len(results)-1].Interface().(error); ok {
		err = errchk
	}
	if len(results) <= mc.Index() {
		return nil, fmt.Errorf("index %d out of range. Function %s returned %d values (indexex start at zero)", mc.Index(), mc.Name(), len(results))
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
func (p *Property) Evaluate(values Values) (interface{}, error) {
	if p.parent == nil {
		//If there is no parent, we must be referring to a map key in values
		if values[p.Name()] == nil {
			return nil, fmt.Errorf("unable to get property - no value named %q exists in Values passed to expression", p.Name())
		}
		return values[p.Name()], nil
	}
	prnt, err := p.parent.Evaluate(values)
	if err != nil {
		return nil, fmt.Errorf("error evaluating parent of %q: %s", p.Name(), err)
	}
	return p.evaluate(prnt)
}

//evaluate does the real evaluation dereferencing pointers along the way.
func (p *Property) evaluate(env interface{}) (result interface{}, err error) {
	var propVal reflect.Value
	//If property name is null, return top-level env object
	if p.Name() == "" {
		return env, nil
	}
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

func (p *Property) String() string {
	return fmt.Sprintf("Property: %s", p.name)
}
