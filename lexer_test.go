package xex

import (
	"bufio"
	"strings"
	"testing"
)

func TestIdentAndInt(t *testing.T) {
	//                                                         1         2         3         4         5         6
	//                                               0123456789012345678901234567890123456789012345678901234567890123456789
	l := NewDefaultLexer(bufio.NewReader(strings.NewReader(`123 4.5 hello_WORLD.FuncName  + - * / % ^ == < != "a string" "" ()[]., false/true`)))
	expected := []*Token{
		{TOKEN_INT, 0, "123"},
		{TOKEN_WHITESPACE, 3, " "},
		{TOKEN_FLOAT, 4, "4.5"},
		{TOKEN_WHITESPACE, 7, " "},
		{TOKEN_IDENT, 8, "hello_WORLD"},
		{TOKEN_SEPARATOR, 19, "."},
		{TOKEN_IDENT, 20, "FuncName"},
		{TOKEN_WHITESPACE, 28, "  "},
		{TOKEN_BINARY_OPERATOR, 30, "+"},
		{TOKEN_WHITESPACE, 31, " "},
		{TOKEN_BINARY_OPERATOR, 32, "-"},
		{TOKEN_WHITESPACE, 33, " "},
		{TOKEN_BINARY_OPERATOR, 34, "*"},
		{TOKEN_WHITESPACE, 35, " "},
		{TOKEN_BINARY_OPERATOR, 36, "/"},
		{TOKEN_WHITESPACE, 37, " "},
		{TOKEN_BINARY_OPERATOR, 38, "%"},
		{TOKEN_WHITESPACE, 39, " "},
		{TOKEN_BINARY_OPERATOR, 40, "^"},
		{TOKEN_WHITESPACE, 41, " "},
		{TOKEN_BINARY_OPERATOR, 42, "=="},
		{TOKEN_WHITESPACE, 44, " "},
		{TOKEN_BINARY_OPERATOR, 45, "<"},
		{TOKEN_WHITESPACE, 46, " "},
		{TOKEN_BINARY_OPERATOR, 47, "!="},
		{TOKEN_WHITESPACE, 49, " "},
		{TOKEN_STRING, 50, "a string"},
		{TOKEN_WHITESPACE, 60, " "},
		{TOKEN_STRING, 61, ""},
		{TOKEN_WHITESPACE, 63, " "},
		{TOKEN_LPAREN, 64, "("},
		{TOKEN_RPAREN, 65, ")"},
		{TOKEN_LINDEX, 66, "["},
		{TOKEN_RINDEX, 67, "]"},
		{TOKEN_SEPARATOR, 68, "."},
		{TOKEN_DELIMITER, 69, ","},
		{TOKEN_WHITESPACE, 70, " "},
		{TOKEN_BOOL, 71, "false"},
		{TOKEN_BINARY_OPERATOR, 76, "/"},
		{TOKEN_BOOL, 77, "true"},
		{TOKEN_EOF, 82, ""},
	}
	l.Run()
	for _, exp := range expected {
		tok := l.NextToken()
		if *exp != *tok {
			t.Fatalf("Expected %s. Got %s", exp, tok)
		}
		if exp.Typ == TOKEN_EOF {
			break
		}
	}
}
