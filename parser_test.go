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
	err := testDoParse(`lib.Address.City`, "London", Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBinaryOperators(t *testing.T) {
	err := testDoParse(`string(add(4.5, 10.5){0} * float64(3)) + "_hello"`, "45_hello", Values{})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMathsWithParentheses(t *testing.T) {
	err := testDoParse(`4 + 3 * 2`, 10, Values{})
	if err != nil {
		t.Error(err)
		return
	}
	err = testDoParse(`(4 + 3) * 2`, 14, Values{})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUnaryOperator(t *testing.T) {
	err := testDoParse(`!true`, false, Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
	//Bad arg
	err = testDoParse(`!~`, nil, Values{"lib": testLib})
	if err == nil {
		t.Error("should have failed with bad arg for unary operator")
		return
	}
	//delete the function
	delete(functions, "not")
	err = testDoParse(`!true`, nil, Values{"lib": testLib})
	if err == nil {
		t.Error("expected function does not exist error")
		return
	}
	//delete from unaryFuncMap
	delete(unaryFuncMap, "!")
	err = testDoParse(`!true`, nil, Values{"lib": testLib})
	if err == nil {
		t.Error("expected error for missing unary function map")
		return
	}
}

func TestIndexOfMapFromMethod(t *testing.T) {
	err := testDoParse(`lib.Authors()[2]`, "George Orwell", Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIndexOfSliceFromMethod(t *testing.T) {
	err := testDoParse(`lib.GetBooks(){0}[2].Title`, "1984", Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestInvalidReturnIndex(t *testing.T) {
	err := testDoParse(`lib.GetBooks(){999}[2].Title`, "1984", Values{"lib": testLib})
	if err == nil {
		t.Error("expected error for invalid return index")
	}
}

func TestSelectAndIterate(t *testing.T) {
	p := DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`select(lib.GetBooks(), "book", book.PublicationYear > 1900)`))),
	}
	ex, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	results, err := ex.Evaluate(Values{"lib": testLib})
	if err != nil {
		t.Error(err)
		return
	}
	list, ok := results.([]*Book)
	if !ok {
		t.Error("did not get a slice of Book")
	}
	chk := map[string]bool{"1984": false, "Animal Farm": false, "The Lion, the With & the Wardrobe": false}
	cnt := 0
	for _, res := range list {
		chk[res.Title] = true
		cnt++
	}
	if cnt != 3 {
		t.Errorf("expected 3 results. Got %d", cnt)
		return
	}
	for book, found := range chk {
		if found != true {
			t.Errorf("did not find %q", book)
		}
	}
}

func TestSelectAndCall(t *testing.T) {
	err := testDoParse(`select(lib.GetBooks(), "book", book.PublicationYear > 1900)[1].Title`, "Animal Farm", Values{"lib": testLib})
	if err != nil {
		t.Error(err)
	}
}

func testDoParse(expression string, expect interface{}, values Values) error {
	p := DefaultParser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(expression))),
	}
	ex, err := p.Parse()
	if err != nil {
		return err
	}
	r, err := ex.Evaluate(values)
	if err != nil {
		return err
	}
	if r != expect {
		return fmt.Errorf("expected %v, got %v", expect, r)
	}
	return nil
}
