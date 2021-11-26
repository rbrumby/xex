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
		t.Fatal(err)
	}
}
func TestFuncMethProp(t *testing.T) {
	lex := NewDefaultLexer(bufio.NewReader(strings.NewReader(
		`add(
			float64(
				multiply( car.GetGearBox().Gear, uint8(99) )
			),
			3.5)`,
	)))
	par := DefaultParser{
		lexer: lex,
	}
	ex, err := par.Parse()
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}
	if res != 201.5 {
		t.Errorf("Expected 201.5, got %f", res)
	}
}

func TestVariadicConcat(t *testing.T) {
	lex := NewDefaultLexer(bufio.NewReader(strings.NewReader(
		`concat("123","-","456","-","789",": ", car.Driver.Name)`,
	)))
	par := DefaultParser{
		lexer: lex,
	}
	ex, err := par.Parse()
	if err != nil {
		t.Fatal(err)
	}

	c := Car{
		Driver: &Driver{Name: "Lando"},
	}
	res, err := ex.Evaluate(Values{"car": c})
	if err != nil {
		t.Fatal(err)
	}
	if res != "123-456-789: Lando" {
		t.Errorf(`Expected "123-456-789: Lando", got %s`, res)
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
	}

	if tst := p.next().Value; tst != "test" {
		t.Errorf("next expected %q. Got %q", "test", tst)
	}

	if num := p.peek(0).Value; num != "123" {
		t.Errorf("peek after peek after next expected %q. Got %q", "123", num)
	}

	if num := p.peek(1).Value; num != "abc" {
		t.Errorf("peek after peek after next expected %q. Got %q", "abc", num)
	}

	if num := p.next().Value; num != "123" {
		t.Errorf("final next expected %q. Got %q", "123", num)
	}

	if num := p.next().Value; num != "abc" {
		t.Errorf("final next expected %q. Got %q", "abc", num)
	}

	if eof := p.next().Typ; eof != TOKEN_EOF {
		t.Errorf("expected eof. Got %v", eof)
	}
}

func TestFunctionalize(t *testing.T) {
	lex := NewDefaultLexer(bufio.NewReader(strings.NewReader(`string((4 + 10) * 3) + "_hello"`)))
	par := DefaultParser{
		lexer: lex,
	}
	ex, err := par.Parse()
	if err != nil {
		t.Fatal(err)
	}
	answer, err := ex.Evaluate(Values{})
	if err != nil {
		t.Fatal(err)
	}
	if answer != "42_hello" {
		t.Errorf("Expected 42_hello. Got %d", answer)
	}
}
