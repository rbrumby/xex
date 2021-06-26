package xex

import (
	"fmt"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	// Repolog.SetRepoLogLevel(capnslog.TRACE)
	l := Lexer{
		Reader: strings.NewReader(`test.thing + something("Hello, \"world\".",
		reference) + 50 - 1.23`),
	}

	tokens, err := l.Lex()
	if err != nil {
		t.Fatal(err)
	}

	expectedTokens := []Token{
		{TOKEN_IDENT, 1, "test"},
		{TOKEN_SEPARATOR, 5, "."},
		{TOKEN_IDENT, 6, "thing"},
		{TOKEN_WHITESPACE, 11, " "},
		{TOKEN_OPERATOR, 12, "+"},
		{TOKEN_WHITESPACE, 13, " "},
		{TOKEN_IDENT, 14, "something"},
		{TOKEN_LPAREN, 23, "("},
		{TOKEN_STRING, 24, "\"Hello, \\\"world\\\".\""},
		{TOKEN_COMMA, 43, ","},
		{TOKEN_WHITESPACE, 44, "\n		"},
		{TOKEN_IDENT, 47, "reference"},
		{TOKEN_RPAREN, 56, ")"},
		{TOKEN_WHITESPACE, 57, " "},
		{TOKEN_OPERATOR, 58, "+"},
		{TOKEN_WHITESPACE, 59, " "},
		{TOKEN_INT, 60, "50"},
		{TOKEN_WHITESPACE, 62, " "},
		{TOKEN_OPERATOR, 63, "-"},
		{TOKEN_WHITESPACE, 64, " "},
		{TOKEN_FLOAT, 65, "1.23"},
	}
	fmt.Println(tokens)
	if len(tokens) != len(expectedTokens) {
		t.Fatalf("got %d tokens, expect %d", len(tokens), len(expectedTokens))
	}

	for i := range tokens {
		if tokens[i] != expectedTokens[i] {
			t.Fatalf("token %d was %s, expected %s", i, tokens[i], expectedTokens[i])
		}
	}
}

func TestLexerBadExpr(t *testing.T) {
	// Repolog.SetRepoLogLevel(capnslog.TRACE)
	l := Lexer{
		Reader: strings.NewReader(`This should fail because $ "isn't" valid`),
	}
	_, err := l.Lex()
	if err == nil {
		t.Fatal("Should have failed with unrecognized token")
	}
}
