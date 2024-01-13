package parser

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode"
)

//Scanner scans expresion text into *Tokens.
//The default OnError function panics with the error thrown. This can be overriden to do whatever you want
type Scanner struct {
	Error  func(pos int, err error)
	pos    int
	src    []rune
	Tokens []*Token
}

func (s *Scanner) OnError(fn func(pos int, err error)) *Scanner {
	s.Error = fn
	return s
}

func NewReaderScanner(src io.Reader) (*Scanner, error) {
	bytes, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, fmt.Errorf("read error: %s", err)
	}
	return NewByteScanner(bytes), nil
}

func NewByteScanner(src []byte) *Scanner {
	return &Scanner{
		src: []rune(string(src)),
		Error: func(pos int, err error) {
			logErr.Panic(err)
		},
	}
}

func (s *Scanner) Scan() *Scanner {
	for !s.eof() {
		err := s.scanToken()
		if err != nil {
			s.Error(s.pos, err)
		}
	}
	s.appendTokens(&Token{
		TokenType: TOKEN_EOF,
		Start:     s.pos,
	})
	return s
}

func (s *Scanner) scanToken() (err ScanError) {
	pos := s.pos
	r := s.consume()
	switch {
	//REMEMBER, when we're here, we are not mid-token. Each case decides from
	//the 1st character what to do to evaluate & consume the rest of the token.
	case unicode.IsSpace(r): //Whitespace - ignore
	case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'): //TOKEN_IDENT / TOKEN_BOOL
		buff := []rune{r}
		s.scanIdent(pos, buff)
	case r >= '0' && r <= '9': //TOKEN_INT / TOKEN_FLOAT
		buff := []rune{r}
		err = s.scanNumberLiteral(pos, buff)
		if err != nil {
			return err
		}
	case r == '"' || r == '`': //TOKEN_STRING
		err := s.scanStringLiteral(pos, r)
		if err != nil {
			return err
		}
	case r == '-':
		if unicode.IsDigit(s.peek()) {
			buff := []rune{r}
			err = s.scanNumberLiteral(pos, buff)
			return
		}
		s.appendTokens(&Token{Start: pos, TokenType: TOKEN_MINUS, Value: string(r)})
	default:
		t := &Token{
			Start: pos,
		}
		if !s.eof() {
			next := string([]rune{r, s.peek()}) //see if there is a 2-char match
			if typ, ok := symbolMap[next]; ok {
				s.consume()
				t.TokenType = typ
				t.Value = next
				s.appendTokens(t)
				return nil
			}
		}
		if this, ok := symbolMap[string(r)]; ok { //if not, try a 1-char match
			t.TokenType = this
			t.Value = string(r)
			s.appendTokens(t)
			return nil
		}
		err = ScanErrorf(s.pos, "unexpected character %q", r)
	}
	return
}

func (s *Scanner) eof() bool {
	return s.pos >= len(s.src)
}

func (s *Scanner) consume() rune {
	if s.eof() {
		panic(errors.New("attempt to consume at EOF"))
	}
	r := s.src[s.pos]
	s.pos++
	return r

}

func (s *Scanner) skip() {
	if !s.eof() {
		s.pos++
	}
}

func (s *Scanner) backup() {
	if s.pos > 0 {
		s.pos--
	}
}

func (s *Scanner) peek() (r rune) {
	r = s.consume()
	s.backup()
	return
}

func (s *Scanner) scanIdent(pos int, buff []rune) {
	for !s.eof() {
		n := s.peek()
		if (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z') || (n >= '0' && n <= '9') {
			s.consume()
			buff = append(buff, n)
			continue
		}
		break
	}
	tok := &Token{
		Start: pos,
	}
	if strings.ToLower(string(buff)) == "true" ||
		strings.ToLower(string(buff)) == "false" {
		tok.TokenType = TOKEN_BOOL
	} else if strings.ToLower(string(buff)) == "nil" {
		tok.TokenType = TOKEN_NIL
	} else {
		tok.TokenType = TOKEN_IDENT
	}
	tok.Value = string(buff)
	s.appendTokens(tok)
}

func (s *Scanner) scanStringLiteral(pos int, terminator rune) (err ScanError) {
	buff := make([]rune, 0)
	esc := false
	done := false
	tok := &Token{Start: s.pos, TokenType: TOKEN_STRING}
loop:
	for !s.eof() {
		switch s.peek() {
		case terminator: //not in escape mode & this is the closing quote / backtick
			if !esc {
				s.skip()
				done = true
				break loop
			}
		case '\\': //flip escape mode
			if !esc { //if we're not in escape mode, that's all we need to do - otherwise we'll consume outside the switch
				esc = !esc
				s.skip()
				continue
			}
		case 't':
			if esc {
				buff = append(buff, '\t')
				esc = !esc
				s.skip()
				continue
			}
		case 'n':
			if esc {
				buff = append(buff, '\n')
				esc = !esc
				s.skip()
				continue
			}
		case 'r':
			if esc {
				buff = append(buff, '\r')
				esc = !esc
				s.skip()
				continue
			}
		}
		n := s.consume()
		buff = append(buff, n)
		esc = false
	}
	if s.eof() && !done {
		err = &scanErr{s.pos, "unterminated string literal"}
	}
	tok.Value = string(buff)
	s.appendTokens(tok)
	return
}

func (s *Scanner) scanNumberLiteral(pos int, buff []rune) (err ScanError) {
	isFloat := false
	for !s.eof() {
		n := s.peek()
		if unicode.IsDigit(n) || n == '.' {
			if n == '.' {
				if isFloat { //found a second '.'
					err = ScanErrorf(s.pos, "unexpected %q in number", n)
					break
				}
				isFloat = true
			}
			s.consume()
			buff = append(buff, n)
			continue
		}
		break
	}
	tok := &Token{
		Start: pos,
	}
	if isFloat {
		tok.TokenType = TOKEN_FLOAT
	} else {
		tok.TokenType = TOKEN_INT
	}
	tok.Value = string(buff)
	s.appendTokens(tok)
	return
}

func (s *Scanner) appendTokens(t ...*Token) {
	s.Tokens = append(s.Tokens, t...)
}

type Token struct {
	TokenType TokenType
	Start     int
	Value     string
}

func (t *Token) String() string {
	return fmt.Sprintf("%s(%d, %d): %q", t.TokenType, t.Start, len(t.Value), t.Value)
}

type TokenType uint8

func (tt TokenType) String() string {
	if tn, ok := TokenTypeNames[tt]; ok {
		return tn
	}
	panic("Missing TokenTypeName") //We've missed an antry in the map!
}

type ScanError error

type scanErr struct {
	pos int
	msg string
}

func ScanErrorf(pos int, msg string, vars ...interface{}) ScanError {
	return &scanErr{
		pos: pos,
		msg: fmt.Sprintf(msg, vars...),
	}
}

func (e *scanErr) Error() string {
	return fmt.Sprintf("%s at position %d", e.msg, e.pos)
}

const (
	TOKEN_UNKNOWN TokenType = iota //zero-value
	TOKEN_NIL
	TOKEN_IDENT
	TOKEN_START_ARGS
	TOKEN_END_ARGS
	TOKEN_SEPARATOR
	TOKEN_STRING
	TOKEN_FLOAT
	TOKEN_INT
	TOKEN_BOOL
	TOKEN_DELIMITER
	TOKEN_START_ARRAY_INDEX
	TOKEN_END_ARRAY_INDEX
	TOKEN_START_RETURN_INDEX
	TOKEN_END_RETURN_INDEX
	TOKEN_EOF
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_MULTIPLY
	TOKEN_DIVIDE
	TOKEN_POWER
	TOKEN_MODULUS
	TOKEN_NOT
	TOKEN_EQUALS
	TOKEN_NOT_EQUALS
	TOKEN_GREATER_THAN
	TOKEN_GREATER_THAN_EQUAL
	TOKEN_LESS_THAN
	TOKEN_LESS_THAN_EQUAL
	TOKEN_AND
	TOKEN_OR
)

var TokenTypeNames = map[TokenType]string{
	TOKEN_IDENT:              "IDENTIFIER",
	TOKEN_START_ARGS:         "START_ARGS",
	TOKEN_END_ARGS:           "END_ARGS",
	TOKEN_SEPARATOR:          "SEPARATOR",
	TOKEN_STRING:             "STRING",
	TOKEN_FLOAT:              "FLOAT",
	TOKEN_INT:                "INT",
	TOKEN_BOOL:               "BOOLEAN",
	TOKEN_DELIMITER:          "DELIMITER",
	TOKEN_START_ARRAY_INDEX:  "START_ARRAY",
	TOKEN_END_ARRAY_INDEX:    "END_ARRAY",
	TOKEN_START_RETURN_INDEX: "START_RETURN_INDEX",
	TOKEN_END_RETURN_INDEX:   "END_RETURN_INDEX",
	TOKEN_EOF:                "EOF",
	TOKEN_PLUS:               "PLUS",
	TOKEN_MINUS:              "MINUS",
	TOKEN_MULTIPLY:           "MULTIPLY",
	TOKEN_DIVIDE:             "DIVIDE",
	TOKEN_POWER:              "POWER",
	TOKEN_MODULUS:            "MODULUS",
	TOKEN_NOT:                "NOT",
	TOKEN_EQUALS:             "EQUAL",
	TOKEN_NOT_EQUALS:         "NOT_EQUAL",
	TOKEN_GREATER_THAN:       "GREATER_THAN",
	TOKEN_GREATER_THAN_EQUAL: "GREATER_THAN_EQUAL",
	TOKEN_LESS_THAN:          "LESS_THAN",
	TOKEN_LESS_THAN_EQUAL:    "LESS_THAN_EQUAL",
	TOKEN_AND:                "AND",
	TOKEN_OR:                 "OR",
	TOKEN_NIL:                "NIL",
	TOKEN_UNKNOWN:            "UNKNOWN",
}

var symbolMap = map[string]TokenType{
	"(":  TOKEN_START_ARGS,
	")":  TOKEN_END_ARGS,
	".":  TOKEN_SEPARATOR,
	",":  TOKEN_DELIMITER,
	"[":  TOKEN_START_ARRAY_INDEX,
	"]":  TOKEN_END_ARRAY_INDEX,
	"{":  TOKEN_START_RETURN_INDEX,
	"}":  TOKEN_END_RETURN_INDEX,
	"+":  TOKEN_PLUS,
	"-":  TOKEN_MINUS,
	"*":  TOKEN_MULTIPLY,
	"/":  TOKEN_DIVIDE,
	"^":  TOKEN_POWER,
	"%":  TOKEN_MODULUS,
	"==": TOKEN_EQUALS,
	"!":  TOKEN_NOT,
	"!=": TOKEN_NOT_EQUALS,
	">":  TOKEN_GREATER_THAN,
	">=": TOKEN_GREATER_THAN_EQUAL,
	"<":  TOKEN_LESS_THAN,
	"<=": TOKEN_LESS_THAN_EQUAL,
	"&&": TOKEN_AND,
	"||": TOKEN_OR,
}
