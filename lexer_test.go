package xex

import (
	"bufio"
	"strings"
	"testing"
)

func TestIdentAndInt(t *testing.T) {
	//                                                                1         2         3         4         5         6         7         8
	//                                                      012345678901234567890123456789012345678901234567890123456789012345678901234567890
	l := NewDefaultLexer(bufio.NewReader(strings.NewReader(`123 4.5 hello_WORLD.FuncName  + - * / % ^ == < != "a string" "" ()[]., false/true`)))
	expected := []*Token{
		{TOKEN_INT, 0, "123", nil},
		{TOKEN_WHITESPACE, 3, " ", nil},
		{TOKEN_FLOAT, 4, "4.5", nil},
		{TOKEN_WHITESPACE, 7, " ", nil},
		{TOKEN_IDENT, 8, "hello_WORLD", nil},
		{TOKEN_SEPARATOR, 19, ".", nil},
		{TOKEN_IDENT, 20, "FuncName", nil},
		{TOKEN_WHITESPACE, 28, "  ", nil},
		{TOKEN_BINARY_OPERATOR, 30, "+", nil},
		{TOKEN_WHITESPACE, 31, " ", nil},
		{TOKEN_BINARY_OPERATOR, 32, "-", nil},
		{TOKEN_WHITESPACE, 33, " ", nil},
		{TOKEN_BINARY_OPERATOR, 34, "*", nil},
		{TOKEN_WHITESPACE, 35, " ", nil},
		{TOKEN_BINARY_OPERATOR, 36, "/", nil},
		{TOKEN_WHITESPACE, 37, " ", nil},
		{TOKEN_BINARY_OPERATOR, 38, "%", nil},
		{TOKEN_WHITESPACE, 39, " ", nil},
		{TOKEN_BINARY_OPERATOR, 40, "^", nil},
		{TOKEN_WHITESPACE, 41, " ", nil},
		{TOKEN_BINARY_OPERATOR, 42, "==", nil},
		{TOKEN_WHITESPACE, 44, " ", nil},
		{TOKEN_BINARY_OPERATOR, 45, "<", nil},
		{TOKEN_WHITESPACE, 46, " ", nil},
		{TOKEN_BINARY_OPERATOR, 47, "!=", nil},
		{TOKEN_WHITESPACE, 49, " ", nil},
		{TOKEN_STRING, 50, "a string", nil},
		{TOKEN_WHITESPACE, 60, " ", nil},
		{TOKEN_STRING, 61, "", nil},
		{TOKEN_WHITESPACE, 63, " ", nil},
		{TOKEN_LPAREN, 64, "(", nil},
		{TOKEN_RPAREN, 65, ")", nil},
		{TOKEN_LINDEX, 66, "[", nil},
		{TOKEN_RINDEX, 67, "]", nil},
		{TOKEN_SEPARATOR, 68, ".", nil},
		{TOKEN_DELIMITER, 69, ",", nil},
		{TOKEN_WHITESPACE, 70, " ", nil},
		{TOKEN_BOOL, 71, "false", nil},
		{TOKEN_BINARY_OPERATOR, 76, "/", nil},
		{TOKEN_BOOL, 77, "true", nil},
		{TOKEN_EOF, 82, "", nil},
	}
	l.Run()
	for _, exp := range expected {
		tok := l.NextToken()
		if *exp != *tok {
			t.Errorf("Expected %s. Got %s", exp, tok)
		}
		if exp.Typ == TOKEN_EOF {
			break
		}
	}
}
