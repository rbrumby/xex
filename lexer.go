package xex

//Lexer is based on Rob Pikes talk (https://www.youtube.com/watch?v=HxaD_trXwRE) about the std lib text template implementation.
import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type TokenType int

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_WHITESPACE
	TOKEN_IDENT
	TOKEN_SEPARATOR
	TOKEN_DELIMITER
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_LINDEX
	TOKEN_RINDEX
	TOKEN_LRESULT
	TOKEN_RRESULT
	TOKEN_BINARY_OPERATOR
	TOKEN_UNARY_OPERATOR
	TOKEN_STRING
	TOKEN_INT
	TOKEN_FLOAT
	TOKEN_BOOL
	TOKEN_ALL_VALUES
	TOKEN_EOF
)

var tokenTypes = []string{
	TOKEN_ERROR:           "ERROR",
	TOKEN_WHITESPACE:      "WHITESPACE",
	TOKEN_IDENT:           "IDENTIFIER",
	TOKEN_SEPARATOR:       "SEPARATOR",
	TOKEN_DELIMITER:       "DELIMITER",
	TOKEN_LPAREN:          "LEFT_PARENTHESIS",
	TOKEN_RPAREN:          "RIGHT_PARENTHESIS",
	TOKEN_LINDEX:          "LEFT_INDEX",
	TOKEN_RINDEX:          "RIGHT_INDEX",
	TOKEN_LRESULT:         "LEFT_RESULT",
	TOKEN_RRESULT:         "RIGHT_RESULT",
	TOKEN_BINARY_OPERATOR: "BINARY_OPERATOR",
	TOKEN_UNARY_OPERATOR:  "UNARY_OPERATOR",
	TOKEN_STRING:          "STRING",
	TOKEN_INT:             "INTEGER",
	TOKEN_FLOAT:           "FLOAT",
	TOKEN_BOOL:            "BOOL",
	TOKEN_ALL_VALUES:      "ALL_VALUES",
	TOKEN_EOF:             "EOF",
}

func (tt TokenType) String() string {
	return tokenTypes[tt]
}

type stateFn func(l *DefaultLexer) stateFn

type Token struct {
	Typ   TokenType
	Start int
	Value string
	Error error
}

func (t *Token) String() string {
	return fmt.Sprintf("%s (%d, %d) = %q", t.Typ.String(), t.Start, len(t.Value), t.Value)
}

type Lexer interface {
	Run()
	NextToken() *Token
	Error() string
}

type DefaultLexer struct {
	Reader *bufio.Reader
	buff   []rune
	err    error
	tokens chan *Token
	pos    int
	start  int
	eof    bool
}

//NewLexer returns a Lexer to read an expression from the provided reader
func NewDefaultLexer(r *bufio.Reader) Lexer {
	return &DefaultLexer{
		Reader: r,
		buff:   make([]rune, 0),
		tokens: make(chan *Token),
	}
}

func (l *DefaultLexer) Error() string {
	return l.err.Error()
}

func (l *DefaultLexer) next() rune {
	r, _, err := l.Reader.ReadRune()
	l.pos++
	if err == io.EOF {
		l.eof = true
	} else if err != nil {
		l.err = err
		l.emit(TOKEN_ERROR)
	}
	return r
}

func (l *DefaultLexer) backup() {
	if !l.eof {
		if err := l.Reader.UnreadRune(); err != nil {
			l.err = err
			l.emit(TOKEN_ERROR)
			return
		}
		l.pos--
	}
}

func (l *DefaultLexer) peek() (r rune) {
	r = l.next()
	l.backup()
	return
}

func (l *DefaultLexer) consume(validFn func(r rune) bool) bool {
	r := l.next()
	if validFn == nil {
		l.buff = append(l.buff, r)
		return true
	}
	if validFn(r) {
		l.buff = append(l.buff, r)
		return true
	}
	l.backup()
	return false
}

func (l *DefaultLexer) consumeUntilInvalid(validFn func(r rune) bool) {
	for l.consume(validFn) {
	}
}

//emit clears the current buffer &
func (l *DefaultLexer) emit(tType TokenType) {
	val := string(l.buff)
	var err error = nil
	if tType == TOKEN_ERROR {
		err = fmt.Errorf("error reading expression: %s", l.err.Error())
	}
	token := &Token{
		Typ:   tType,
		Start: l.start,
		Value: val,
		Error: err,
	}
	l.buff = l.buff[:0]
	l.start = l.pos
	l.tokens <- token
}

func (l *DefaultLexer) Run() {
	go func() {
		for fn := lexNextToken; fn != nil; {
			fn = fn(l)
		}
	}()
}

func (l *DefaultLexer) NextToken() *Token {
	return <-l.tokens
}

func lexNextToken(l *DefaultLexer) stateFn {
	r := l.peek()
	switch {
	case l.eof:
		l.emit(TOKEN_EOF)
	case unicode.IsSpace(r):
		l.consumeUntilInvalid(unicode.IsSpace)
		l.emit(TOKEN_WHITESPACE)
		return lexNextToken
	case unicode.IsLetter(r): //ident must start with a letter but can contain alphanumerics & underscores
		l.consumeUntilInvalid(isIdentChar)
		if string(l.buff) == "true" || string(l.buff) == "false" {
			l.emit(TOKEN_BOOL)
			return lexNextToken
		}
		l.emit(TOKEN_IDENT)
		return lexNextToken
	case r == '.':
		l.consume(nil)
		l.emit(TOKEN_SEPARATOR)
		return lexNextToken
	case r == ',':
		l.consume(nil)
		l.emit(TOKEN_DELIMITER)
		return lexNextToken
	case r == '(':
		l.consume(nil)
		l.emit(TOKEN_LPAREN)
		return lexNextToken
	case r == ')':
		l.consume(nil)
		l.emit(TOKEN_RPAREN)
		return lexNextToken
	case r == '[':
		l.consume(nil)
		l.emit(TOKEN_LINDEX)
		return lexNextToken
	case r == ']':
		l.consume(nil)
		l.emit(TOKEN_RINDEX)
		return lexNextToken
	case r == '{':
		l.consume(nil)
		l.emit(TOKEN_LRESULT)
		return lexNextToken
	case r == '}':
		l.consume(nil)
		l.emit(TOKEN_RRESULT)
		return lexNextToken
	case r == '#':
		l.consume(nil)
		l.emit(TOKEN_ALL_VALUES)
		return lexNextToken
	case isOperator(r):
		if r == '!' {
			l.consume(nil)
			r = l.peek()
			if r != '=' {
				l.emit(TOKEN_UNARY_OPERATOR)
			}
		} else {
			l.consumeUntilInvalid(isOperator)
			l.emit(TOKEN_BINARY_OPERATOR)
		}
		return lexNextToken
	case isQuote(r):
		return lexStringLiteral
	case unicode.IsDigit(r): //if it starts with a number its a number - but it might get delegated to lexFloat if it contains a "."
		return lexNumber
	default:
		l.err = fmt.Errorf("unrecognized character %q at position %d", r, l.pos)
		l.consume(nil)
		l.emit(TOKEN_ERROR)
	}
	return nil
}

func lexNumber(l *DefaultLexer) stateFn {
	for l.consume(unicode.IsDigit) {
	}
	if l.peek() == '.' {
		//consume the '.' and treat as a float
		l.consume(nil)
		return lexFloat
	}
	l.emit(TOKEN_INT)
	return lexNextToken
}

func lexFloat(l *DefaultLexer) stateFn {
	for l.consume(unicode.IsDigit) {
	}
	l.emit(TOKEN_FLOAT)
	return lexNextToken
}

func lexStringLiteral(l *DefaultLexer) stateFn {
	//Get quote starting character so we know what will close the string
	start := l.peek()
	l.consume(isQuote)                                             //Consume the initial quote
	l.consumeUntilInvalid(func(r rune) bool { return r != start }) //Consume until we find the matching character
	l.consume(isQuote)                                             //Then consume the closing quote
	l.buff = l.buff[1 : len(l.buff)-1]                             //discard leading & trailing quote runes
	l.emit(TOKEN_STRING)
	return lexNextToken
}

func isIdentChar(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsNumber(r)
}
func isOperator(r rune) bool {
	return strings.ContainsRune("+-/*%^<>=!", r)
}

func isQuote(r rune) bool {
	return strings.ContainsRune("\"`'", r)
}
