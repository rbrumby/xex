package xex

import (
	"bufio"
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

func TestSimpleProperties(t *testing.T) {
	p := DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`lib.Address.City`))),
	}
	ex, err := p.Parse()
	if err != nil {
		t.Error(err)
	}
	r, err := ex.Evaluate(Values{"lib": testLib})
	if err != nil {
		t.Error(err)
	}
	if r != "London" {
		t.Errorf("Expected London. Got %d", r)
		return
	}
}

func TestBinaryOperators(t *testing.T) {
	lex := NewDefaultLexer(bufio.NewReader(strings.NewReader(`string(add(4.5, 10.5){0} * float64(3)) + "_hello"`)))
	par := DefaultParser{
		lexer: lex,
	}
	ex, err := par.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	answer, err := ex.Evaluate(Values{})
	if err != nil {
		t.Error(err)
		return
	}
	if answer != "45_hello" {
		t.Errorf("Expected 45_hello. Got %d", answer)
		return
	}
}

func TestUnaryOperator(t *testing.T) {
	p := DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`!true`))),
	}
	ex, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	r, err := ex.Evaluate(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if r != false {
		t.Errorf("Expected false, got %v", r)
		return
	}
	//Bad arg
	p = DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`!~`))),
	}
	_, err = p.Parse()
	if err == nil {
		t.Error("Expected error for unary operator with bad arg")
		return
	}
	//delete the function
	delete(functions, "not")
	p = DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`!true`))),
	}
	_, err = p.Parse()
	if err == nil {
		t.Error("Expected function does not exist error")
		return
	}
	//delete from unaryFuncMap
	delete(unaryFuncMap, "!")
	p = DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`!true`))),
	}
	_, err = p.Parse()
	if err == nil {
		t.Error("Expected error for missing unary function map")
		return
	}
}

func TestIndexOfMapFromMethod(t *testing.T) {
	p := DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`lib.Authors()[2]`))),
	}
	ex, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ex)
	r, err := ex.Evaluate(Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
	if r != "George Orwell" {
		t.Errorf("Expected George Orwell, got %v", r)
		return
	}
}

func TestIndexOfSliceFromMethod(t *testing.T) {
	p := DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`lib.GetBooks()[2].Title`))),
	}
	ex, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ex)
	r, err := ex.Evaluate(Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
	if r != "1984" {
		t.Errorf("Expected 1984, got %v", r)
		return
	}
}
