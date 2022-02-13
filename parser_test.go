package xex

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestEmptyExpression(t *testing.T) {
	lex := NewDefaultLexer(bufio.NewReader(strings.NewReader(
		"",
	)))
	par := DefaultParser{
		lexer: lex,
	}
	_, err := par.Parse()
	if err.Error() != "empty expression" {
		t.Error(err)
		return
	}
}

func TestPeekNext(t *testing.T) {
	p := &DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader("test 123 abc"))),
		buff:  make([]*Token, 0),
	}

	p.lexer.Run()

	if tst := p.peek(0).Value; tst != "test" {
		t.Errorf("peek expected %q. Got %q", "test", tst)
		return
	}

	if tst := p.next().Value; tst != "test" {
		t.Errorf("next expected %q. Got %q", "test", tst)
		return
	}

	if num := p.peek(0).Value; num != "123" {
		t.Errorf("peek after peek after next expected %q. Got %q", "123", num)
		return
	}

	if num := p.peek(1).Value; num != "abc" {
		t.Errorf("peek after peek after next expected %q. Got %q", "abc", num)
		return
	}

	if num := p.next().Value; num != "123" {
		t.Errorf("final next expected %q. Got %q", "123", num)
		return
	}

	if num := p.next().Value; num != "abc" {
		t.Errorf("final next expected %q. Got %q", "abc", num)
		return
	}

	if eof := p.next().Typ; eof != TOKEN_EOF {
		t.Errorf("expected eof. Got %v", eof)
		return
	}
}

func TestParseEquation(t *testing.T) {
	err := justParseAndCheckString("(4 + 3.5) * 2", "Expression: multiply(nil(addOrConcat(4,3.5)),2)")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseFunction(t *testing.T) {
	err := justParseAndCheckString(`concat("a","b")`, `Expression: concat("a","b")`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUnaryOperator(t *testing.T) {
	err := justParseAndCheckString(`!true`, `Expression: not(true)`)
	if err != nil {
		t.Error(err)
		return
	}
	//Bad arg
	err = justParseAndCheckString(`!~`, "")
	if err == nil {
		t.Error("should have failed with bad arg for unary operator")
		return
	}
	//delete the function
	delete(functions, "not")
	err = justParseAndCheckString(`!true`, "")
	if err == nil {
		t.Error("expected function does not exist error")
		return
	}
	//delete from unaryFuncMap
	delete(unaryFuncMap, "!")
	err = justParseAndCheckString(`!true`, "")
	if err == nil {
		t.Error("expected error for missing unary function map")
		return
	}
}

func TestParseSubProp(t *testing.T) {
	err := justParseAndCheckString(`person.Address`, `Expression: person.Address`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseFuncNoArgs(t *testing.T) {
	err := justParseAndCheckString(`nil()`, `Expression: nil()`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseFunctionWithReturnIndex(t *testing.T) {
	err := justParseAndCheckString(`concat("a","b"){55}`, `Expression: concat("a","b")`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestFunctionWithSubProp(t *testing.T) {
	err := justParseAndCheckString(`concat("a","b").child`, `Expression: concat("a","b").child`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseFunctionWithReturnIndexAndSubProp(t *testing.T) {
	err := justParseAndCheckString(`concat("a","b"){55}.child`, `Expression: concat("a","b").child`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseMethodNoArgs(t *testing.T) {
	err := justParseAndCheckString(`something.Do()`, `Expression: something.Do()`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseMethodWithReturnIndex(t *testing.T) {
	err := justParseAndCheckString(`something.Do("a","b"){55}`, `Expression: something.Do("a","b")`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMethodWithSubProp(t *testing.T) {
	err := justParseAndCheckString(`something.Do("a","b").child`, `Expression: something.Do("a","b").child`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseMethodWithReturnIndexAndSubProp(t *testing.T) {
	err := justParseAndCheckString(`something.Do("a","b"){55}.child`, `Expression: something.Do("a","b").child`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseCollectionIndex(t *testing.T) {
	err := justParseAndCheckString(`collection[5]`, `Expression: indexOf(collection,5)`)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestParseBadCollectionIndex(t *testing.T) {
	err := justParseAndCheckString(`collection[5 f]`, `Expression: indexOf(collection,5)`)
	if err != nil && strings.HasPrefix(err.Error(), "unexpected token") {
		return
	}
	t.Error("should have failed with unexpected token")
}

func TestParseCollectionIndexSubProp(t *testing.T) {
	err := justParseAndCheckString(`collection[5].prop`, `Expression: indexOf(collection,5).prop`)
	if err != nil {
		t.Error(err)
		return
	}
}

func justParseAndCheckString(exStr string, expected string) error {
	p := DefaultParser{lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(exStr)))}
	ex, err := p.Parse()
	if err != nil {
		return err
	}
	if ex.String() != expected {
		return fmt.Errorf("expected %q, Got %q", expected, ex.String())
	}
	return nil
}

type dummyParser struct{}

func (dp *dummyParser) Parse() (node *Expression, err error) {
	return nil, errors.New("not a real implementation")
}

func TestCustomParser(t *testing.T) {
	_, err := NewStr(
		"add(4,1)",
		func(p Parser) Parser {
			return &dummyParser{}
		},
	)
	if err.Error() != "not a real implementation" {
		t.Error("Expected dummyParser Parse() error")
	}
}
