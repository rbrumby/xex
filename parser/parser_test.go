package parser

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	logInf.SetOutput(ioutil.Discard)
	m.Run()
}

func TestFuncNoArgs(t *testing.T) {
	p := &Parser{}
	ast, err := p.Parse(NewByteScanner([]byte("fn1()")))
	if err != nil {
		t.Fatal(err)
	}
	if ast.String() != "fn1()" {
		t.Fatalf("unexpected result %s", ast.String())
	}
}

func TestFuncWithArgs(t *testing.T) {
	p := &Parser{}
	ast, err := p.Parse(NewByteScanner([]byte(`fn1(prop, "string", 99)`)))
	if err != nil {
		t.Fatal(err)
	}
	if ast.String() != "fn1(prop, STRING{string}, INT{99})" {
		t.Fatalf("unexpected result %s", ast.String())
	}
}

func TestMethodCallOnFuncResult(t *testing.T) {
	p := &Parser{}
	ast, err := p.Parse(NewByteScanner([]byte("fn1().Meth1()")))
	if err != nil {
		t.Fatal(err)
	}
	if ast.String() != "fn1().Meth1()" {
		t.Fatalf("unexpected result %s", ast.String())
	}
}

func TestSimpleArithmetic(t *testing.T) {
	p := &Parser{}
	ast, err := p.Parse(NewByteScanner([]byte("5 + 8 / 2")))
	if err != nil {
		t.Fatal(err)
	}
	if ast.String() != "addOrConcat(INT{5}, divide(INT{8}, INT{2}))" {
		t.Fatalf("unexpected result %s", ast.String())
	}
}

func TestNegativeNumber(t *testing.T) {
	p := &Parser{}
	ast, err := p.Parse(NewByteScanner([]byte("-777.7")))
	if err != nil {
		t.Fatal(err)
	}
	if ast.String() != "FLOAT{-777.7}" {
		t.Fatalf("unexpected result %s", ast.String())
	}
}

func TestUnaryOpOnGroup(t *testing.T) {
	p := &Parser{}
	ast, err := p.Parse(NewByteScanner([]byte("-(5 + 8)")))
	if err != nil {
		t.Fatal(err)
	}
	if ast.String() != "subtract(INT{0}, nil(addOrConcat(INT{5}, INT{8})))" {
		t.Fatalf("unexpected result %s", ast.String())
	}
}

func TestComparison(t *testing.T) {
	p := &Parser{}
	ast, err := p.Parse(NewByteScanner([]byte(`x.y == 1`)))
	if err != nil {
		t.Fatal(err)
	}
	if ast.String() != "equals(x.y, INT{1})" {
		t.Fatalf("unexpected result %s", ast.String())
	}
}

func TestConsecutiveLiterals(t *testing.T) {
	p := &Parser{}
	_, err := p.Parse(NewByteScanner([]byte("123 567")))
	if err == nil {
		t.Fatal("should have failed with unexpected token error")
	}
}

func TestUnclosedParentheses(t *testing.T) {
	p := &Parser{}
	_, err := p.Parse(NewByteScanner([]byte("func1(abc")))
	if err == nil {
		t.Fatal("should have failed with expected right parenthesis error")
	}
}

func TestEmptyExpression(t *testing.T) {
	p := &Parser{}
	_, err := p.Parse(NewByteScanner([]byte("")))
	if err == nil {
		t.Fatal("should have failed with unexpected EOF")
	}
}

// utils

type ASTPrinter struct {
	depth int
	t     *testing.T
}

func (p *ASTPrinter) Visit(n ASTNode) {
	switch t := n.(type) {
	case *ASTFunction:
		p.t.Logf("%sFunc: %s:", p.prefix(), t.name)
		p.depth++
		for _, arg := range t.args.values {
			arg.Accept(p)
		}
		p.depth--
	case *ASTMethod:
		if t.parent != nil {
			t.parent.Accept(p)
			p.depth++
		}
		p.t.Logf("%sMethod: %s:\n", p.prefix(), t.name)
		p.depth++
		for _, arg := range t.args.values {
			arg.Accept(p)
		}
		p.depth--
	case *ASTArguments:
		for _, arg := range t.values {
			arg.Accept(p)
		}
	case *ASTProperty:
		if t.parent != nil {
			t.parent.Accept(p)
			p.depth++
		}
		p.t.Logf("%sProp: %s\n", p.prefix(), t.name)
	case *ASTLiteral:
		if t.token.TokenType == TOKEN_STRING {
			p.t.Logf("%sLiteral (%s): %q\n", p.prefix(), t.token.TokenType.String(), t.token.Value)
		} else {
			p.t.Logf("%sLiteral(%s): %s\n", p.prefix(), t.token.TokenType.String(), t.token.Value)
		}
	default:
		p.t.Logf("DEFAULT: %s\n", reflect.TypeOf(n).Elem().Name())
	}
}

func (p *ASTPrinter) prefix() string {
	var b strings.Builder
	for d := 0; d < p.depth; d++ {
		b.WriteString("-")
	}
	return b.String()
}
