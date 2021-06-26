package xex

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

var DefaultGrammar = StandardGrammar{}

type TokenType int

const (
	TOKEN_NOT_RECOGNIZED TokenType = iota
	TOKEN_NOT_DETERMINED
	TOKEN_WHITESPACE
	TOKEN_IDENT
	TOKEN_SEPARATOR
	TOKEN_COMMA
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_LINDEX
	TOKEN_RINDEX
	TOKEN_OPERATOR
	TOKEN_COMPARATOR
	TOKEN_STRING
	TOKEN_INT
	TOKEN_FLOAT
)

var tokenTypes = []string{
	TOKEN_NOT_RECOGNIZED: "NOT_RECOGNIZED",
	TOKEN_NOT_DETERMINED: "UNDETERMINED",
	TOKEN_WHITESPACE:     "WHITESPACE",
	TOKEN_IDENT:          "IDENT",
	TOKEN_SEPARATOR:      "SEPARATOR",
	TOKEN_COMMA:          "COMMA",
	TOKEN_LPAREN:         "LEFT_PARENTHESIS",
	TOKEN_RPAREN:         "RIGHT_PARENTHESIS",
	TOKEN_LINDEX:         "LEFT_INDEX",
	TOKEN_RINDEX:         "RIGHT_INDEX",
	TOKEN_OPERATOR:       "OPERATOR",
	TOKEN_COMPARATOR:     "COMPARATOR",
	TOKEN_STRING:         "STRING",
	TOKEN_INT:            "INTEGER",
	TOKEN_FLOAT:          "FLOAT",
}

func (tt TokenType) String() string {
	return tokenTypes[tt]
}

type Token struct {
	Typ   TokenType
	Start int
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("%s: (%d, %d), %q", t.Typ.String(), t.Start, len(t.Value), t.Value)
}

type Lexer struct {
	Reader       io.Reader
	currentToken Token
}

func (l Lexer) Lex() (tokens []Token, err error) {
	br := bufio.NewReader(l.Reader)
	tokens = make([]Token, 0)
	runes := make([]rune, 0)

	l.currentToken.Typ = TOKEN_NOT_DETERMINED
	l.currentToken.Start = 1

	for {
		if r, _, err := br.ReadRune(); err == nil {
			logger.Println("Looping")
			runes = append(runes, r)
			tType := l.getType(runes)

			logger.Tracef("%s is %s and current token is %s\n", string(runes), tType, l.currentToken)

			switch tType {

			case TOKEN_NOT_RECOGNIZED:
				//We have run out of token. Back up and add this token to the output slice.
				if err = br.UnreadRune(); err != nil {
					return nil, err
				}
				if l.currentToken.Typ == TOKEN_NOT_DETERMINED {
					return nil, fmt.Errorf("unrecognized token %q", string(runes))
				}
				tokens = append(tokens, l.currentToken)
				runes = runes[:0]
				logger.Tracef("got token %s\n", l.currentToken)
				l.currentToken = Token{
					Typ:   TOKEN_NOT_DETERMINED,
					Start: l.currentToken.Start + len(string(l.currentToken.Value)),
				}
				logger.Tracef("reset current token to %s\n", l.currentToken)
			default:
				//Set token type & value
				l.currentToken.Typ = tType
				l.currentToken.Value = string(runes)
			}
		} else {
			if err == io.EOF {
				//Store the last token
				tokens = append(tokens, l.currentToken)
			}
			break
		}
	}
	if err != nil && err != io.EOF {
		return nil, err
	}
	return tokens, nil

}

func (l Lexer) getType(runes []rune) TokenType {
	m := DefaultGrammar.Matches([]byte(string(runes)))
	logger.Tracef("%q matched %v", string(runes), m)
	switch len(m) {
	case 0:
		return TOKEN_NOT_RECOGNIZED
	case 1:
		return m[0]
	default:
		return TOKEN_NOT_DETERMINED
	}
}

type Grammar interface {
	Matches([]byte) []TokenType
}

type StandardGrammar struct{}

func (sg StandardGrammar) Get() map[TokenType]*regexp.Regexp {
	return map[TokenType]*regexp.Regexp{
		TOKEN_WHITESPACE: regexp.MustCompile(`^[\n\r\t\s\x00]+$`),
		TOKEN_IDENT:      regexp.MustCompile(`^[a-zA-Z_]+[a-zA-Z0-9_]*$`),
		TOKEN_SEPARATOR:  regexp.MustCompile(`^\.$`),
		TOKEN_COMMA:      regexp.MustCompile(`^\,$`),
		TOKEN_LPAREN:     regexp.MustCompile(`^[(]$`),
		TOKEN_RPAREN:     regexp.MustCompile(`^[)]$`),
		TOKEN_LINDEX:     regexp.MustCompile(`^[[]]$`),
		TOKEN_RINDEX:     regexp.MustCompile(`^[]]$`),
		TOKEN_OPERATOR:   regexp.MustCompile(`^[\+\-\/\*]$`),
		TOKEN_COMPARATOR: regexp.MustCompile(`^>|>=|<|<=|==$`),
		TOKEN_STRING:     regexp.MustCompile(`^"((\\")|[^"])*$|^"((\\")|[^"])*"$`),
		TOKEN_INT:        regexp.MustCompile(`^[0-9]+$`),
		TOKEN_FLOAT:      regexp.MustCompile(`^[0-9]+\.([0-9]+)?$`),
	}
}

func (sg StandardGrammar) Matches(b []byte) []TokenType {
	types := make([]TokenType, 0)
	for tt, reg := range sg.Get() {
		if reg.Match(b) {
			logger.Tracef("grammar match: %q matches %q\n", string(b), tt)
			types = append(types, tt)
		}
	}
	return types
}
