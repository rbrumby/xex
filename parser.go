package xex

import (
	"errors"
	"fmt"
	"strconv"
)

var unaryFuncMap map[string]string = map[string]string{
	"!": "not",
}

var binaryFuncMap map[string]string = map[string]string{
	"+":  "addOrConcat",
	"-":  "subtract",
	"*":  "multiply",
	"/":  "divide",
	"^":  "pow",
	"%":  "mod",
	"==": "equals",
	"!=": "notEquals",
	">":  "greaterThan",
	">=": "greaterThanEqual",
	"<":  "lessThan",
	"<=": "lessThanEqual",
	"&&": "and",
	"||": "or",
}

type Parser interface {
	Parse() (node *Expression, err error)
}

type DefaultParser struct {
	lexer Lexer
	buff  []*Token
}

//next gets & consumes the next non-whitepace *Token either from the buffer or the lexer
func (p *DefaultParser) next() (tok *Token) {
	if len(p.buff) <= 0 { //nothing in the buffer - call peek(0)
		p.peek(0)
	}
	tok = p.buff[0] //return the first token from the buffer

	p.buff = p.buff[1:]
	return tok
}

// peek gets the nth non-whitespace *Token without consuming any tokens
// It will store read tokens in a buffer for future access
func (p *DefaultParser) peek(index int) (tok *Token) {
	for len(p.buff) <= index {
		tok = p.lexer.NextToken()
		if tok.Typ != TOKEN_WHITESPACE {
			p.buff = append(p.buff, tok)
			break
		}
	}
	tok = p.buff[index]
	return tok
}

func (p *DefaultParser) Parse() (ex *Expression, err error) {
	go p.lexer.Run()
	root, err := p.parse(nil)
	if err != nil {
		return nil, err
	}
	if root == nil {
		err = errors.New("empty expression")
	}
	return NewExpression(root), err
}

func (p *DefaultParser) parse(parent Node) (node Node, err error) {
	this := p.next()
	if this.Typ == TOKEN_EOF {
		return nil, nil
	}

	if this.Typ == TOKEN_ERROR {
		return nil, fmt.Errorf("Lexer found invalid token: %s", this)
	}

	next := p.peek(0)
	switch {

	case this.Typ == TOKEN_IDENT && next.Typ == TOKEN_LPAREN: //This is a call (method or function - not Foo Fighters)
		logger.Debugf("Func or method, %v followed by %v\n", this, next)
		call, err := p.parseCall(parent, this)
		if err != nil {
			return nil, err
		}
		return p.functionalize(call)

	case this.Typ == TOKEN_LPAREN: //This is a use of parentheses for logical grouping - make a nil function
		this = &Token{
			Typ:   TOKEN_IDENT,
			Start: this.Start,
			Value: "nil",
		}
		logger.Debugf("Nil function, %v followed by %v\n", this, next)
		call, err := p.parseCall(nil, this)
		if err != nil {
			return nil, err
		}
		return p.functionalize(call)

	case this.Typ == TOKEN_UNARY_OPERATOR:
		unFnName, ok := unaryFuncMap[this.Value]
		if !ok {
			return nil, fmt.Errorf("unary function %q not found", this.Value)
		}
		unFn, err := GetFunction(unFnName)
		if err != nil {
			return nil, err
		}
		unArg, err := p.parse(nil)
		if err != nil {
			return nil, err
		}
		return NewFunctionCall(unFn, []Node{unArg}, 0), nil //unary functions can only use first return val (hardcoded zero)

	case this.Typ == TOKEN_IDENT && next.Typ == TOKEN_SEPARATOR: //This is a Parent Property
		logger.Debugf("Property with children, %v followed by %v\n", this, next)
		p.next() //consume the SEPARATOR
		prop, err := p.parseCollectionIndex(NewProperty(this.Value, parent))
		if err != nil {
			return nil, err
		}
		return p.parse(prop) //Don't functionalize - we already know "next" isn't a binary operator!!

	case this.Typ == TOKEN_IDENT: //This is a Property with no children
		logger.Debugf("Property without children, %v followed by %v\n", this, next)
		prop, err := p.parseCollectionIndex(NewProperty(this.Value, parent))
		if err != nil {
			return nil, err
		}
		return p.functionalize(prop)

	case this.Typ == TOKEN_STRING:
		logger.Debugf("String literal, %v \n", this)
		return p.functionalize(NewLiteral(this.Value))

	case this.Typ == TOKEN_INT:
		logger.Debugf("Int literal, %v \n", this)
		val, err := strconv.ParseInt(this.Value, 0, 0) //use default int bitsize
		v := int(val)
		if err != nil {
			return nil, err
		}
		return p.functionalize(NewLiteral(v))

	case this.Typ == TOKEN_FLOAT:
		logger.Debugf("Float literal, %v \n", this)
		val, err := strconv.ParseFloat(this.Value, 64)
		if err != nil {
			return nil, err
		}
		return p.functionalize(NewLiteral(val))

	case this.Typ == TOKEN_BOOL:
		logger.Debugf("Bool literal, %v \n", this)
		val, err := strconv.ParseBool(this.Value)
		if err != nil {
			return nil, err
		}
		return p.functionalize(NewLiteral(val))
	}
	return nil, fmt.Errorf("unexpected token: %v (%s)", this.Value, this.Typ.String())
}

func (p *DefaultParser) parseCall(parent Node, ident *Token) (node Node, err error) {
	if p.peek(0).Typ == TOKEN_LPAREN {
		p.next() //consume the LPAREN
	}
	args := make([]Node, 0)
	logger.Debugf("Collecting args for %s\n", ident.Value)
	for { //build args for the method or function call
		next := p.peek(0)
		if next.Typ == TOKEN_RPAREN { //this check needed in case there are no args - we check again below after finding an arg
			rp := p.next() //consume the rparen
			logger.Debugf("Found %s, skipped & consumed it\n", rp)
			break
		}
		if next.Typ == TOKEN_EOF {
			return nil, fmt.Errorf("unlcosed parenthesis for %s", ident.Value)
		}
		if next.Typ == TOKEN_DELIMITER {
			del := p.next() //consume the delimiter
			logger.Debugf("Found %s, skipped & consumed it\n", del)
			continue
		}
		logger.Debugf("Building arg %d for %s\n", len(args), ident.Value)
		arg, err := p.parse(nil)
		if err != nil {
			return nil, fmt.Errorf("error parsing argument %d for %s: %s", len(args), ident.Value, err)
		}
		args = append(args, arg)
		next = p.peek(0)
		if next.Typ == TOKEN_RPAREN {
			rp := p.next() //consume the rparen
			logger.Debugf("Found %s, skipped & consumed it\n", rp)
			break
		}
	}
	logger.Debugf("Finished collecting args for %s\n", ident.Value)
	//If there is an index declared, set it
	callRtnIdx := 0
	if p.peek(0).Typ == TOKEN_LRESULT {
		callRtnIdx, err = p.parseReturnIndex()
		if err != nil {
			return nil, err
		}
	}
	if parent == nil { //FunctionCall
		fn, err := GetFunction(ident.Value)
		if err != nil {
			return nil, fmt.Errorf("parse error (%s): %s,", ident.String(), err)
		}
		fc, err := p.parseCollectionIndex(NewFunctionCall(fn, args, callRtnIdx)) //wraps call in collection indexOf Node if an index is specified
		if err != nil {
			return nil, err
		}
		if p.peek(0).Typ == TOKEN_SEPARATOR { //we have a child
			p.next()           //consume the separator
			return p.parse(fc) //parse the child passing this *FunctionCall as the parent
		}
		return fc, nil
	} else { //MethodCall
		mc, err := p.parseCollectionIndex(NewMethodCall(ident.Value, parent, args, callRtnIdx)) //wraps call in collection indexOf Node if an index is specified
		if err != nil {
			return nil, err
		}
		if p.peek(0).Typ == TOKEN_SEPARATOR { //we have a child
			p.next()           //consume the separator
			return p.parse(mc) //parse the child passing this *FunctionCall as the parent
		}
		return mc, nil
	}
}

func (p *DefaultParser) functionalize(node Node) (Node, error) {
	next := p.peek(0)
	if next.Typ == TOKEN_BINARY_OPERATOR {
		next = p.next() //consume the binary operator
		if binFunc, isBinFunc := binaryFuncMap[next.Value]; isBinFunc {
			fn, err := GetFunction(binFunc)
			if err != nil {
				return nil, fmt.Errorf("binary operator %q mapped to unknown function %q", next, binFunc)
			}
			args := make([]Node, 2)
			args[0] = node
			args[1], err = p.parse(nil)
			if err != nil {
				return nil, fmt.Errorf("error parsing binary function argument: %s", err)
			}
			logger.Debugf("functionalize created function %q with args %s", fn.name, args)
			return NewFunctionCall(fn, args, 0), nil //binary functions can only sue the first return value (hardcoded zero)
		}
		return nil, fmt.Errorf("unrecognized binary operator %s", next)
	}
	return node, nil
}

func (p *DefaultParser) parseReturnIndex() (index int, err error) {
	//user specified the argument they want returned from the call
	lres := p.peek(0) //check for LRESULT
	if lres.Typ != TOKEN_LRESULT {
		return 0, nil //return default zero-index literal
	}
	lres = p.next() //consume LRETURN
	logger.Debugf("Found %s, skipped & consumed it\n", lres)
	idx, err := p.parse(nil)
	if err != nil {
		return 0, fmt.Errorf("error parsing return index: %s", err)
	}
	if idxlit, ok := idx.(*Literal); ok {
		rres := p.next() //consume rresrult
		if rres.Typ != TOKEN_RRESULT {
			return 0, fmt.Errorf("unexpected token %s, expected %s", rres, TOKEN_RRESULT)
		}
		logger.Debugf("Found %s, skipped & consumed it\n", rres)
		switch idxint := idxlit.value.(type) {
		case int:
			return int(idxint), nil
		case int8:
			return int(idxint), nil
		case int16:
			return int(idxint), nil
		case int32:
			return int(idxint), nil
		case int64:
			return int(idxint), nil
		}
	}
	return 0, fmt.Errorf("unexpected return index %s, expected int literal", idx)
}

func (p *DefaultParser) parseCollectionIndex(in Node) (out Node, err error) {
	lin := p.peek(0) //check forLINDEX
	if lin.Typ != TOKEN_LINDEX {
		return in, nil //return unmodified Node
	}
	lin = p.next() //consume LINDEX
	logger.Debugf("Found %s, skipped & consumed it\n", lin)
	index, err := p.parse(nil) //no parent
	if err != nil {
		return nil, fmt.Errorf("error parsing index: %s", err)
	}
	if p.peek(0).Typ != TOKEN_RINDEX {
		return nil, fmt.Errorf("expeceted end of index. found %s", err)
	}
	rin := p.next() //consume rindex
	if rin.Typ != TOKEN_RINDEX {
		return nil, fmt.Errorf("unexpected token %s, expected %s", rin, TOKEN_RINDEX)
	}
	logger.Debugf("Found %s, skipped & consumed it\n", rin)
	indFn, err := GetFunction("indexOf")
	if err != nil {
		return nil, fmt.Errorf("parseCollectionIndex: %s", err)
	}
	//TODO: Need to call parse with this as parent if next token is a "."
	if p.peek(0).Typ == TOKEN_SEPARATOR {
		sep := p.next() //consume separator
		logger.Debugf("Found %s, skipped & consumed it\n", sep)
		return p.parse(NewFunctionCall(indFn, []Node{in, index}, 0))
	}
	return NewFunctionCall(indFn, []Node{in, index}, 0), nil
}
