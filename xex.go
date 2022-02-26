package xex

import (
	"fmt"
	"reflect"
	"strings"
)

//Node is a node in the compiled expression tree
type Node interface {
	Name() string
	Evaluate(values Values) (interface{}, error)
	String() string
}

//Call represents something that can be called (a function or method)
type Call interface {
	Node
	Arg(arg Node)
}

//Values is a map which an Expression is evaluated against
type Values map[string]interface{}

//Expression will be evaluated to return a value.
//It is the root of the graph of Nodes used to produce a value but can also be
type Expression struct {
	root Node
}

//NewExpression creates an expression
func NewExpression(root Node) *Expression {
	return &Expression{root}
}

//Name always returns "<expression>" - expressions don't have names.
func (e *Expression) Name() string {
	return "<expression>"
}

//Evaluate evaluates the expression
func (e *Expression) Evaluate(values Values) (interface{}, error) {
	if values == nil {
		values = make(Values)
	}
	return e.root.Evaluate(values)
}

//String returns a string representation of the expression
func (e *Expression) String() string {
	return fmt.Sprintf("Expression: %s", e.root.String())
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
	return fc.function.Name
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
		if argNode == nil {
			args[i] = nil
			continue
		}
		if reflect.TypeOf(fc.function.impl).NumIn() > i &&
			reflect.TypeOf(fc.function.impl).In(i).Implements(reflect.TypeOf((*Node)(nil)).Elem()) {
			//This arg shouldn't be evaluated - the function expects a Node
			if ex, ok := argNode.(Node); ok {
				args[i] = ex
				continue
			}
			return nil, fmt.Errorf("%q expected argument %d to be a xex.Node", fc.Name(), i)
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
		return nil, fmt.Errorf("index %d out of range. Function %s returned %d values (indices start at zero)", fc.Index(), fc.Name(), len(results))
	}
	return results[fc.Index()], nil
}

func (f *FunctionCall) String() string {
	out := &strings.Builder{}
	out.WriteString(fmt.Sprintf("%s(", f.Name()))
	for i, arg := range f.arguments {
		out.WriteString(arg.String())
		if i < len(f.arguments)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString(")")
	return out.String()
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
	if s, ok := (l.value).(string); ok {
		return fmt.Sprintf("%q", s)
	}
	if s, ok := (l.value).(fmt.Stringer); ok {
		return fmt.Sprintf("%q", s)
	}
	return fmt.Sprintf("%v", l.value)
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
	out := &strings.Builder{}
	prefix := ""
	if m.parent != nil {
		out.WriteString(m.parent.String())
		prefix = "."
	}
	out.WriteString(fmt.Sprintf("%s%s(", prefix, m.Name()))
	for i, arg := range m.arguments {
		out.WriteString(arg.String())
		if i < len(m.arguments)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString((")"))
	return out.String()
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
		return nil, fmt.Errorf("index %d out of range. Function %s returned %d values (indices start at zero)", mc.Index(), mc.Name(), len(results))
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
	out := &strings.Builder{}
	prefix := ""
	if p.parent != nil {
		out.WriteString(p.parent.String())
		prefix = "."
	}
	out.WriteString(fmt.Sprintf("%s%s", prefix, p.Name()))
	return out.String()
}
