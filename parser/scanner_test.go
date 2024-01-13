package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestReadAndScan(t *testing.T) {
	s := newTestReaderScanner(strings.NewReader("&&"), t).Scan()
	reportTokens(t, s.Tokens)
	err := compareResults(
		[]*Token{
			{TokenType: TOKEN_AND, Start: 0, Value: "&&"},
			{TokenType: TOKEN_EOF, Start: 2, Value: ""},
		},
		s.Tokens,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestScanReservedWords(t *testing.T) {
	s := newTestByteScanner([]byte("&& == ! != ( [ ) } < <= > >= || +-/* ., {] ^ %"), t).Scan()
	reportTokens(t, s.Tokens)
	err := compareResults(
		[]*Token{
			{TokenType: TOKEN_AND, Start: 0, Value: "&&"},
			{TokenType: TOKEN_EQUALS, Start: 3, Value: "=="},
			{TokenType: TOKEN_NOT, Start: 6, Value: "!"},
			{TokenType: TOKEN_NOT_EQUALS, Start: 8, Value: "!="},
			{TokenType: TOKEN_START_ARGS, Start: 11, Value: "("},
			{TokenType: TOKEN_START_ARRAY_INDEX, Start: 13, Value: "["},
			{TokenType: TOKEN_END_ARGS, Start: 15, Value: ")"},
			{TokenType: TOKEN_END_RETURN_INDEX, Start: 17, Value: "}"},
			{TokenType: TOKEN_LESS_THAN, Start: 19, Value: "<"},
			{TokenType: TOKEN_LESS_THAN_EQUAL, Start: 21, Value: "<="},
			{TokenType: TOKEN_GREATER_THAN, Start: 24, Value: ">"},
			{TokenType: TOKEN_GREATER_THAN_EQUAL, Start: 26, Value: ">="},
			{TokenType: TOKEN_OR, Start: 29, Value: "||"},
			{TokenType: TOKEN_PLUS, Start: 32, Value: "+"},
			{TokenType: TOKEN_MINUS, Start: 33, Value: "-"},
			{TokenType: TOKEN_DIVIDE, Start: 34, Value: "/"},
			{TokenType: TOKEN_MULTIPLY, Start: 35, Value: "*"},
			{TokenType: TOKEN_SEPARATOR, Start: 37, Value: "."},
			{TokenType: TOKEN_DELIMITER, Start: 38, Value: ","},
			{TokenType: TOKEN_START_RETURN_INDEX, Start: 40, Value: "{"},
			{TokenType: TOKEN_END_ARRAY_INDEX, Start: 41, Value: "]"},
			{TokenType: TOKEN_POWER, Start: 43, Value: "^"},
			{TokenType: TOKEN_MODULUS, Start: 45, Value: "%"},
			{TokenType: TOKEN_EOF, Start: 46, Value: ""},
		},
		s.Tokens,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestDefaultErrorHandler(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic")
		}
		if err, ok := r.(error); ok {
			if err.Error() != "unexpected character '@' at position 1" {
				t.Errorf("expected unexpected char, got %q", err)
			}
		}
	}()
	NewByteScanner([]byte("@")).Scan() //don't use newTestScanner as it reports errors instead of panicking
}

func TestStringLiteral(t *testing.T) {
	s := newTestByteScanner([]byte(`"li\\t\"e\tr\na\rl\\"`), t).Scan()
	reportTokens(t, s.Tokens)
	err := compareResults(
		[]*Token{
			{TokenType: TOKEN_STRING, Start: 1, Value: "li\\t\"e\tr\na\rl\\"}, //skip index 0 - `"`
			{TokenType: TOKEN_EOF, Start: 21, Value: ""},
		},
		s.Tokens,
	)
	if err != nil {
		t.Error(err)
	}

}

func TestBoolean(t *testing.T) {
	s := newTestByteScanner([]byte(`true false`), t).Scan()
	reportTokens(t, s.Tokens)
	err := compareResults(
		[]*Token{
			{TokenType: TOKEN_BOOL, Start: 0, Value: "true"},
			{TokenType: TOKEN_BOOL, Start: 5, Value: "false"},
			{TokenType: TOKEN_EOF, Start: 10, Value: ""},
		},
		s.Tokens,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestUnclosedStringLiteral(t *testing.T) {
	var fail string
	s := newTestByteScanner([]byte(`"literal`), t).
		OnError(func(pos int, err error) {
			fail = fmt.Sprintf("custom error: %s", err)
		}).
		Scan()
	reportTokens(t, s.Tokens)
	if fail != "custom error: unterminated string literal at position 8" {
		t.Errorf(`Should have failed with "custom error: unterminated string literal at position 8", got %q`, fail)
	}
}

func TestIdent(t *testing.T) {
	s := newTestByteScanner([]byte(`property.subprop.method1(function1())`), t).
		Scan()
	reportTokens(t, s.Tokens)
	err := compareResults(
		[]*Token{
			{TokenType: TOKEN_IDENT, Start: 0, Value: "property"},
			{TokenType: TOKEN_SEPARATOR, Start: 8, Value: "."},
			{TokenType: TOKEN_IDENT, Start: 9, Value: "subprop"},
			{TokenType: TOKEN_SEPARATOR, Start: 16, Value: "."},
			{TokenType: TOKEN_IDENT, Start: 17, Value: "method1"},
			{TokenType: TOKEN_START_ARGS, Start: 24, Value: "("},
			{TokenType: TOKEN_IDENT, Start: 25, Value: "function1"},
			{TokenType: TOKEN_START_ARGS, Start: 34, Value: "("},
			{TokenType: TOKEN_END_ARGS, Start: 35, Value: ")"},
			{TokenType: TOKEN_END_ARGS, Start: 36, Value: ")"},
			{TokenType: TOKEN_EOF, Start: 37, Value: ""},
		},
		s.Tokens,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestNumbers(t *testing.T) {
	s := newTestByteScanner([]byte(`967 954.123 0.1 -45 -33.333`), t).Scan()
	reportTokens(t, s.Tokens)
	err := compareResults(
		[]*Token{
			{TokenType: TOKEN_INT, Start: 0, Value: "967"},
			{TokenType: TOKEN_FLOAT, Start: 4, Value: "954.123"},
			{TokenType: TOKEN_FLOAT, Start: 12, Value: "0.1"},
			{TokenType: TOKEN_INT, Start: 16, Value: "-45"},
			{TokenType: TOKEN_FLOAT, Start: 20, Value: "-33.333"},
			{TokenType: TOKEN_EOF, Start: 27, Value: ""},
		},
		s.Tokens,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestBadNumbers(t *testing.T) {
	expect := "unexpected '.' in number at position 3"
	s := newTestByteScanner([]byte(`0.2.3.4`), t).
		OnError(func(pos int, err error) {
			if err.Error() != expect {
				t.Errorf("expected %q, got %q", expect, err)
			}
		}).
		Scan()
	reportTokens(t, s.Tokens)
}

type badReader struct{}

func (br badReader) Read(b []byte) (n int, err error) {
	return 0, errors.New("error from the bad reader!!")
}

func TestReadFailure(t *testing.T) {
	_, err := NewReaderScanner(badReader{})
	if err == nil || err.Error() != "read error: error from the bad reader!!" {
		t.Error("expected bad reader error")
	}
}

//utils --------------------------------------------

func newTestReaderScanner(src io.Reader, t *testing.T) *Scanner {
	s, err := NewReaderScanner(src)
	if err != nil {
		t.Fatal(err)
	}
	s.Error = func(pos int, err error) {
		t.Error(err)
	}
	return s
}
func newTestByteScanner(src []byte, t *testing.T) *Scanner {
	s := NewByteScanner(src)
	s.Error = func(pos int, err error) {
		t.Error(err)
	}
	return s
}

func compareResults(expected []*Token, actual []*Token) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("expected %d tokens, got %d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if *actual[i] != *expected[i] {
			return fmt.Errorf("expected %s, got %s", expected[i], actual[i])
		}
	}
	return nil
}

func reportTokens(t *testing.T, toks []*Token) {
	for _, tok := range toks {
		t.Log(tok)
	}
}
