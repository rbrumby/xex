package xex

import (
	"bufio"
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
func TestFuncMethProp(t *testing.T) {
	lex := NewDefaultLexer(bufio.NewReader(strings.NewReader(
		`add(
			float64(
				multiply( car.GetGearBox(){0}.Gear, uint8(99) )
			),
			3.5)`,
	)))
	par := DefaultParser{
		lexer: lex,
	}
	ex, err := par.Parse()
	if err != nil {
		t.Error(err)
		return
	}

	c := Car{
		Engine: Engine{
			Gearbox{
				Gear:  2,
				Gears: []Gear{{1.8}, {3.4}, {6.2}, {11.9}},
			},
		},
	}
	res, err := ex.Evaluate(Values{"car": c})
	if err != nil {
		t.Error(err)
		return
	}
	if res != 201.5 {
		t.Errorf("Expected 201.5, got %f", res)
		return
	}
}

func TestVariadicConcat(t *testing.T) {
	lex := NewDefaultLexer(bufio.NewReader(strings.NewReader(
		`concat("123","-","456","-","789",": ", car.Driver.Name){0}`,
	)))
	par := DefaultParser{
		lexer: lex,
	}
	ex, err := par.Parse()
	if err != nil {
		t.Error(err)
		return
	}

	c := Car{
		Driver: &Driver{Name: "Lando"},
	}
	res, err := ex.Evaluate(Values{"car": c})
	if err != nil {
		t.Error(err)
		return
	}
	if res != "123-456-789: Lando" {
		t.Errorf(`Expected "123-456-789: Lando", got %s`, res)
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

func TestFunctionalize(t *testing.T) {
	lex := NewDefaultLexer(bufio.NewReader(strings.NewReader(`string((4 + 10) * 3) + "_hello"`)))
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
	if answer != "42_hello" {
		t.Errorf("Expected 42_hello. Got %d", answer)
		return
	}
}
