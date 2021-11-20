package xex

import (
	"errors"
	"fmt"
	"strconv"
)

/*
	ref  = prop | func | meth
	node = lit | ref
	lit  = int | float | string | bool
	idx  = lindex int rindex
	func = ident lparen ( node ( delim node )* )+ rparen ( idx )+
	prop = ( node sep )* ident ( idx )+
	meth = ( node sep )* ident lparen ( node ( delim node )* )+ rparen ( idx )+

	6 * some.prop.method(5, 3 - 1) + 3
	INT WHITE OPER WHITE IDENT SEP IDENT SEP IDENT LPAREN INT SEP WHITE INT WHITE OPER WHITE INT RPAREN WHITE OPER WHITE INT
*/

type NodeType int

const (
	TypeProperty NodeType = iota
	TypeString
	TypeInteger
	TypeFLoat
	TypeFunction
	TypeMethod
)

var binaryFuncMap map[string]string = map[string]string{
	"+":  "add",
	"-":  "subtract",
	"*":  "multiply",
	"/":  "divide",
	"^":  "pow",
	"%":  "mod",
	"==": "equals",
	//TODO: Greater than, etc
}

///TODO: Not (unaryFuncMap)

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
		return p.parseCall(parent, this)

	case this.Typ == TOKEN_IDENT && next.Typ == TOKEN_SEPARATOR:
		//This is a Parent Property
		logger.Debugf("Property with children, %v followed by %v\n", this, next)
		p.next() //consume the SEPARATOR
		prop := NewProperty(this.Value, parent)
		return p.parse(prop)

	case this.Typ == TOKEN_IDENT:
		//This is a Property with no children
		logger.Debugf("Property without children, %v followed by %v\n", this, next)
		return NewProperty(this.Value, parent), nil

	case this.Typ == TOKEN_STRING:
		logger.Debugf("String literal, %v \n", this)
		return NewLiteral(this.Value), nil

	case this.Typ == TOKEN_INT:
		logger.Debugf("Int literal, %v \n", this)
		val, err := strconv.ParseInt(this.Value, 0, 64)
		return NewLiteral(val), err

	case this.Typ == TOKEN_FLOAT:
		logger.Debugf("Float literal, %v \n", this)
		val, err := strconv.ParseFloat(this.Value, 64)
		return NewLiteral(val), err

	case this.Typ == TOKEN_BOOL:
		logger.Debugf("Bool literal, %v \n", this)
		val, err := strconv.ParseBool(this.Value)
		return NewLiteral(val), err
		// case this.Typ == TOKEN_DELIMITER:
		// 	logger.Debugf("Found %s, nothing to do\n", this.Typ)
		// 	return nil, nil
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
	if parent == nil {
		//FunctionCall
		fn, err := GetFunction(ident.Value)
		if err != nil {
			return nil, fmt.Errorf("parse error (%s): %s,", ident.String(), err)
		}
		//TODO: Indexing output - remove hardcoded zero
		fc := NewFunctionCall(fn, args, 0)
		if p.peek(0).Typ == TOKEN_SEPARATOR {
			//we have a child
			p.next()           //consume the separator
			return p.parse(fc) //parse the child passing this *FunctionCall as the parent
		}
		return fc, nil
	} else {
		//MethodCall
		//TODO: Indexing output - remove hardcoded zero
		mc := NewMethodCall(ident.Value, parent, args, 0)
		if p.peek(0).Typ == TOKEN_SEPARATOR {
			//we have a child
			p.next()           //consume the separator
			return p.parse(mc) //parse the child passing this *FunctionCall as the parent
		}
		return mc, nil
	}
}
