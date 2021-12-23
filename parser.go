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
	"!=": "notEqual",         //TODO
	">":  "greaterThan",      //TODO
	">=": "greaterThanEqual", //TODO
	"<":  "lessThan",         //TODO
	"<=": "lessThanEqual",    //TODO
}

type Parser interface {
	Parse() (node *Expression, err error)
	Lexer(l Lexer)
}

type DefaultParser struct {
	lexer Lexer
	buff  []*Token
}

func (p *DefaultParser) Lexer(l Lexer) {
	p.lexer = l
}

func (p *DefaultParser) next() (tok *Token) {
	if len(p.buff) <= 0 {
		//nothing in the buffer - call peek(0)
		p.peek(0)
	}
	//return the first token from the buffer
	tok = p.buff[0]
	p.buff = p.buff[1:]
	return tok
}

func (p *DefaultParser) peek(index int) (tok *Token) {
	// Get the next non-whitespace *Token
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
	next := p.peek(0)
	switch {
	case this.Typ == TOKEN_IDENT && next.Typ == TOKEN_LPAREN:
		//This is a call (method or function - not Foo Fighters)
		logger.Debugf("Func or method, %v followed by %v\n", this, next)
		call, err := p.parseCall(parent, this)
		if err != nil {
			return nil, err
		}
		return p.functionalize(call)

	case this.Typ == TOKEN_LPAREN:
		//This is a use of parentheses for logical grouping - make a nil function
		this = &Token{
			Typ:   TOKEN_IDENT,
			Start: this.Start,
			Value: "nil",
		}
		//we already consumed the LPAREN
		logger.Debugf("Nil function, %v followed by %v\n", this, next)
		call, err := p.parseCall(nil, this)
		if err != nil {
			return nil, err
		}
		return p.functionalize(call)

	case this.Typ == TOKEN_IDENT && next.Typ == TOKEN_SEPARATOR:
		//This is a Parent Property
		logger.Debugf("Property with children, %v followed by %v\n", this, next)
		p.next() //consume the SEPARATOR
		prop := NewProperty(this.Value, parent)
		//Don't try to functionalize - we already know "next" isn't a binary operator!!
		return p.parse(prop)

	case this.Typ == TOKEN_IDENT:
		//This is a Property with no children
		logger.Debugf("Property without children, %v followed by %v\n", this, next)
		return p.functionalize(NewProperty(this.Value, parent))

	case this.Typ == TOKEN_STRING:
		logger.Debugf("String literal, %v \n", this)
		return p.functionalize(NewLiteral(this.Value))

	case this.Typ == TOKEN_INT:
		logger.Debugf("Int literal, %v \n", this)
		val, err := strconv.ParseInt(this.Value, 0, 64)
		if err != nil {
			return nil, err
		}
		return p.functionalize(NewLiteral(val))

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
	//build args for the method or function call
	logger.Debugf("Collecting args for %s\n", ident.Value)
	for {
		next := p.peek(0)
		if next.Typ == TOKEN_RPAREN {
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
	if parent == nil {
		//FunctionCall
		fn, err := GetFunction(ident.Value)
		if err != nil {
			return nil, fmt.Errorf("parse error (%s): %s,", ident.String(), err)
		}
		//TODO: Indexing output - remove hardcoded zero
		fc := NewFunctionCall(fn, args, callRtnIdx)
		if p.peek(0).Typ == TOKEN_SEPARATOR {
			//we have a child
			p.next()           //consume the separator
			return p.parse(fc) //parse the child passing this *FunctionCall as the parent
		}
		return fc, nil
	} else {
		//MethodCall
		//TODO: Indexing output - remove hardcoded zero
		mc := NewMethodCall(ident.Value, parent, args, callRtnIdx)
		if p.peek(0).Typ == TOKEN_SEPARATOR {
			//we have a child
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
			//TODO: indexing - remove hardcoded zero
			return NewFunctionCall(fn, args, 0), nil
		}
		return nil, fmt.Errorf("unrecognized binary operator %s", next)
	}
	return node, nil
}

func (p *DefaultParser) parseReturnIndex() (index int, err error) {
	//user specified the argument they want returned from the call
	lin := p.next() //consume lindex
	logger.Debugf("Found %s, skipped & consumed it\n", lin)
	if p.peek(0).Typ != TOKEN_INT {
		return 0, fmt.Errorf("unexpected call index: %s - Expected integer", p.peek(0))
	}
	index, err = strconv.Atoi(p.next().Value)
	if err != nil {
		return 0, fmt.Errorf("error parsing call index: %s", err)
	}
	if p.peek(0).Typ != TOKEN_RRESULT {
		return 0, fmt.Errorf("expeceted end of call index. found %s", err)
	}
	rin := p.next() //consume rindex
	logger.Debugf("Found %s, skipped & consumed it\n", rin)
	return
}
