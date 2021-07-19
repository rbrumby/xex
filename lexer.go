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
	TOKEN_OPERATOR
	TOKEN_COMPARATOR
	TOKEN_STRING
	TOKEN_INT
	TOKEN_FLOAT
	TOKEN_EOF
)

var tokenTypes = []string{
	TOKEN_ERROR:      "ERROR",
	TOKEN_WHITESPACE: "WHITESPACE",
	TOKEN_IDENT:      "IDENT",
	TOKEN_SEPARATOR:  "SEPARATOR",
	TOKEN_DELIMITER:  "COMMA",
	TOKEN_LPAREN:     "LEFT_PARENTHESIS",
	TOKEN_RPAREN:     "RIGHT_PARENTHESIS",
	TOKEN_LINDEX:     "LEFT_INDEX",
	TOKEN_RINDEX:     "RIGHT_INDEX",
	TOKEN_OPERATOR:   "OPERATOR",
	TOKEN_COMPARATOR: "COMPARATOR",
	TOKEN_STRING:     "STRING",
	TOKEN_INT:        "INTEGER",
	TOKEN_FLOAT:      "FLOAT",
	TOKEN_EOF:        "EOF",
}

func (tt TokenType) String() string {
	return tokenTypes[tt]
}

type stateFn func(l *Lexer) stateFn

type Token struct {
	Typ   TokenType
	Start int
	Value string
}

func (t *Token) String() string {
	return fmt.Sprintf("%s (%d, %d) = %q", t.Typ.String(), t.Start, len(t.Value), t.Value)
}

type Lexer struct {
	Reader *bufio.Reader
	buff   []rune
	err    error
	out    chan *Token
	pos    int
	start  int
	eof    bool
}

//NewLexer returns a *Lexer to read an expression from the provided reader
func NewLexer(r *bufio.Reader) *Lexer {
	return &Lexer{
		Reader: r,
		buff:   make([]rune, 0),
		out:    make(chan *Token),
	}
}

func (l *Lexer) Error() string {
	return l.err.Error()
}

func (l *Lexer) next() rune {
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

func (l *Lexer) backup() {
	if !l.eof {
		if err := l.Reader.UnreadRune(); err != nil {
			l.err = err
			l.emit(TOKEN_ERROR)
			return
		}
		l.pos--
	}
}

func (l *Lexer) peek() (r rune) {
	r = l.next()
	l.backup()
	return
}

func (l *Lexer) consume(validFn func(r rune) bool) bool {
	r := l.next()
	if validFn(r) {
		l.buff = append(l.buff, r)
		return true
	}
	l.backup()
	return false
}

func (l *Lexer) consumeUntilInvalid(validFn func(r rune) bool) {
	for l.consume(validFn) {
	}
}

//emit clears the current buffer &
func (l *Lexer) emit(tType TokenType) {
	val := string(l.buff)
	if tType == TOKEN_ERROR {
		val = fmt.Sprintf("error reading expression: %s", l.err.Error())
	}
	token := &Token{
		Typ:   tType,
		Start: l.start,
		Value: val,
	}
	l.buff = l.buff[:0]
	l.start = l.pos
	l.out <- token
}

func (l *Lexer) Run() {
	for fn := lexNextToken; fn != nil; {
		fn = fn(l)
	}
}

func lexNextToken(l *Lexer) stateFn {
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
		l.emit(TOKEN_IDENT)
		return lexNextToken
	case r == '.':
		l.consume(func(r rune) bool { return true })
		l.emit(TOKEN_SEPARATOR)
		return lexNextToken
	case r == ',':
		l.consume(func(r rune) bool { return true })
		l.emit(TOKEN_DELIMITER)
		return lexNextToken
	case r == '(':
		l.consume(func(r rune) bool { return true })
		l.emit(TOKEN_LPAREN)
		return lexNextToken
	case r == ')':
		l.consume(func(r rune) bool { return true })
		l.emit(TOKEN_RPAREN)
		return lexNextToken
	case r == '[':
		l.consume(func(r rune) bool { return true })
		l.emit(TOKEN_LINDEX)
		return lexNextToken
	case r == ']':
		l.consume(func(r rune) bool { return true })
		l.emit(TOKEN_RINDEX)
		return lexNextToken
	case isOperator(r): //operators are single runes so we can do a single consume
		l.consume(isOperator)
		l.emit(TOKEN_OPERATOR)
		return lexNextToken
	case isComparator(r): //operators can be 1 or 2 runes. We know the 1st consume will succeed. We don't care if the 2nd does or not!
		l.consume(isComparator)
		l.consume(isComparator)
		l.emit(TOKEN_COMPARATOR)
		return lexNextToken
	case isQuote(r):
		return lexStringLiteral
	case unicode.IsDigit(r): //if it starts with a number its a number - but it might get delegated to lexFloat if it contains a "."
		return lexNumber
	default:
		l.err = fmt.Errorf("unrecognized character %q at position %d", r, l.pos)
		l.emit(TOKEN_ERROR)
	}
	return nil
}

func lexNumber(l *Lexer) stateFn {
	for l.consume(unicode.IsDigit) {
	}
	if l.peek() == '.' {
		//consume the '.' and treat as a float
		l.consume(func(r rune) bool { return true })
		return lexFloat
	}
	l.emit(TOKEN_INT)
	return lexNextToken
}

func lexFloat(l *Lexer) stateFn {
	for l.consume(unicode.IsDigit) {
	}
	l.emit(TOKEN_FLOAT)
	return lexNextToken
}

func lexStringLiteral(l *Lexer) stateFn {
	//Get quote starting character so we know what will close the string
	start := l.peek()
	//Consume the initial quote
	l.consume(isQuote)
	//Consume until we find the matching character
	l.consumeUntilInvalid(func(r rune) bool { return r != start })
	//Then consume the closing quote
	l.consume(isQuote)
	l.emit(TOKEN_STRING)
	return lexNextToken
}

func isIdentChar(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsNumber(r)
}
func isOperator(r rune) bool {
	return strings.ContainsRune("+-/*%^", r)
}

func isComparator(r rune) bool {
	return strings.ContainsRune("!=<>", r)
}

func isQuote(r rune) bool {
	return strings.ContainsRune("\"`'", r)
}
