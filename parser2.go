package xex

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

//ParserOoption is a function which can be passed to New (or NewStr) to modify default behaviour
//(such as using different Parser or Lexer implementations & configuring them).
//To replace the default parser, pass a function which returns a different Parser.
//To replace the Lexer pass a function which modifies changes p.Lexer & returns p.
type ParserOption func(p *Parser) *Parser

//New creates an Expression from a *bufio.Readerby default using a DefaultParser & DefaultLexer.
//ParserOption functions can be passed to change this behaviour.
func New(r *bufio.Reader, opts ...ParserOption) (ex *Expression, err error) {
	p := NewParser(NewDefaultLexer(r))
	for _, opt := range opts {
		p = opt(p)
	}
	return p.Parse()
}

//NewStr creates an Expression from a string by default using a DefaultParser & DefaultLexer.
//ParserOption functions can be passed to change this behaviour.
func NewStr(s string, opts ...ParserOption) (ex *Expression, err error) {
	b := bufio.NewReader(strings.NewReader(s))
	return New(b, opts...)
}

var unaryFuncMap map[string]string = map[string]string{
	"!": "not",
}

var binaryFuncMap map[string]string = map[string]string{
	"+":  "addOrConcat",
	"-":  "subtract",
	"*":  "multiply",
	"/":  "divide",
	"^":  "pow",
	"%":  "mod",
	"==": "equals",
	"!=": "notEquals",
	">":  "greaterThan",
	">=": "greaterThanEqual",
	"<":  "lessThan",
	"<=": "lessThanEqual",
	"&&": "and",
	"||": "or",
}

// NewParser returns an initialised Parser
func NewParser(l Lexer) *Parser {
	return &Parser{
		lexer: l,
	}
}

type Parser struct {
	lexer Lexer
	buff  []*Token
}

func (p *Parser) Parse() (node *Expression, err error) {
	ast := NewASTNode(p)
	err = ast.Build(TOKEN_EOF)
	if err != nil {
		return nil, err
	}
	fmt.Println(ast)
	return nil, errors.New("not implemented yet")
}

//next gets & consumes the next non-whitepace *Token either from the buffer or the lexer
func (p *Parser) next() (tok *Token) {
	if len(p.buff) <= 0 { //nothing in the buffer - call peek(0)
		p.peek(0)
	}
	tok = p.buff[0] //return the first token from the buffer

	p.buff = p.buff[1:]
	return tok
}

// peek gets the nth non-whitespace *Token without consuming any tokens
// It will store read tokens in a buffer for future access
func (p *Parser) peek(index int) (tok *Token) {
	for len(p.buff) <= index {
		tok = p.lexer.NextToken()
		if tok.Typ != TOKEN_WHITESPACE {
			p.buff = append(p.buff, tok)
			break
		}
	}
	tok = p.buff[index]
	return tok
}

type ASTNode struct {
	parser   *Parser
	tokens   []*Token
	next     *ASTNode
	children []*ASTNode
}

func NewASTNode(parser *Parser) *ASTNode {
	return &ASTNode{
		parser:   parser,
		tokens:   make([]*Token, 0),
		children: make([]*ASTNode, 0),
	}
}

func (ast *ASTNode) Build(until TokenType) (err error) {
	tok := ast.parser.peek(0)
loop:
	for tok.Typ != until {
		switch tok.Typ {
		case TOKEN_WHITESPACE: //ignore whitespace
			continue loop
		case TOKEN_BINARY_OPERATOR, TOKEN_UNARY_OPERATOR, TOKEN_DELIMITER: //separator tokens
			ast.next = NewASTNode(ast.parser)      //create a node to hold the token
			ast.next.tokens[0] = ast.parser.next() //populate the operator / delimiter node with 1 token
			ast.next.next = NewASTNode(ast.parser) //continue parsing after the operator / delimiter
			ast.next.next.Build(until)
			break loop
		case TOKEN_LPAREN, TOKEN_LINDEX, TOKEN_LRESULT: //create child groups
			tok = ast.parser.next() //consume the previously peeked Token
			ast.tokens = append(ast.tokens, tok)
			child := NewASTNode(ast.parser)
			err := ast.Build(tok.Typ + 1) //build child until the TOKEN_R... equivalent of the token we found
			if err != nil {
				return err
			}
			ast.children = append(ast.children, child)
			tok = ast.parser.peek(0)
			if tok.Typ == until {
				ast.tokens = append(ast.tokens, tok)
				continue loop
			}
			return fmt.Errorf("unexpected token %s before %s", tok.Typ.String(), until.String())
		default:
			ast.tokens = append(ast.tokens, ast.parser.next())
		}
		tok = ast.parser.peek(0)
	}
	return
}
