package xex

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

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

func TestReturnIndexOutOfRange(t *testing.T) {
	err := testDoParse(`lib.GetBooks(){999}[2].Title`, "1984", Values{"lib": testLib})
	if err != nil && strings.Contains(err.Error(), "out of range") {
		return
	}
	t.Errorf("should have got err about return index out of range, got %q", err)
}

func TestSelectAndIterate(t *testing.T) {
	p := &Parser{
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

func TestChildOfMethodResult(t *testing.T) {
	err := testDoParse(`lib.GetAddress().City`, "London", Values{"lib": testLib})
	if err != nil {
		t.Error(err)
	}
}

func TestChildOfFunctionResult(t *testing.T) {
	err := testDoParse(`select(lib.GetBooks(), "book", book.PublicationYear > 1900)[1].Title`, "Animal Farm", Values{"lib": testLib})
	if err != nil {
		t.Error(err)
	}
}

func TestBadMethodCollectionIndex(t *testing.T) {
	p := &Parser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`lib.GetBooks()[[]]`))),
	}
	_, err := p.Parse()
	if err == nil {
		t.Error("should have failed with unexpected LEFT_INDEX")
		return
	}
}

func TestBadFunctionCollectionIndex(t *testing.T) {
	p := &Parser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`concat("a","b")[[]]`))),
	}
	_, err := p.Parse()
	if err == nil {
		t.Error("should have failed with unexpected LEFT_INDEX")
		return
	}
}

func TestMethodCollectionIndexOutOfRange(t *testing.T) {
	p := &Parser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`lib.GetBooks()[9999]`))),
	}
	ex, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}
	_, err = ex.Evaluate(Values{"lib": testLib})
	if err != nil && strings.Contains(err.Error(), "slice index out of range") {
		return
	}
	t.Errorf("should have got err about slice index out of range, got %q", err)
}

func TestBadFunctionName(t *testing.T) {
	p := &Parser{
		lexer: NewDefaultLexer(bufio.NewReader(strings.NewReader(`XXX("a","b")`))),
	}
	_, err := p.Parse()
	if err == nil {
		t.Error("should have fail with function not found")
		return
	}
}

func TestPropertyOfFunctionReturnVal(t *testing.T) {
	RegisterFunction(
		&Function{
			"test_getABook",
			FunctionDocumentation{},
			func() *Book {
				return testLib.Books[0]
			},
		},
	)

	err := testDoParse(`test_getABook().Title`, "Sense & Sensibility", nil)
	if err != nil {
		t.Error(err)
	}

}

func testDoParse(expression string, expect interface{}, values Values) error {
	ex, err := NewStr(expression)
	if err != nil {
		return err
	}
	logger.Debugf(ex.String())
	r, err := ex.Evaluate(values)
	if err != nil {
		return err
	}
	if r != expect {
		return fmt.Errorf("expected %v, got %v", expect, r)
	}
	return nil
}
