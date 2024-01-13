package parser

import (
	"fmt"
	"strings"
)

// expression -> comparison ;
// comparison -> term ( ( "==" | "!=" | ">" | ">=" | "<" | "<=" ) term )* ;
// term       -> factor ( ( "-" | "+" ) factor )* ;
// factor     -> unary ( ( "/" | "*" | "^" | "%" ) unary )* ;
// unary      -> ( "!" | "-" ) unary | group ;
// group      ->  "(" expression ")" ;
// method     -> expression "." IDENT "(" arguments ")"
// function   -> IDENT "(" arguments ")"
// property   -> expression ("." IDENT)*
// arguments  -> expression ( "," expression )* ;
// literal    -> INT | FLOAT | STRING | BOOLEAN | NIL ;

// todo - array index & return index

type Parser struct {
	scanner *Scanner
	pos     int
}

type ASTNode interface {
	Accept(visitor ASTVisitor)
	String() string
}

type ASTProperty struct {
	name   string
	parent ASTNode
}

func (n *ASTProperty) Accept(v ASTVisitor) {
	v.Visit(n)
}

func (n *ASTProperty) String() string {
	if n.parent != nil {
		return fmt.Sprintf("%s.%s", n.parent.String(), n.name)
	}
	return fmt.Sprintf("%s", n.name)
}

type ASTFunction struct {
	name string
	args *ASTArguments
}

func (n *ASTFunction) Accept(v ASTVisitor) {
	v.Visit(n)
}

func (n *ASTFunction) String() string {
	return fmt.Sprintf("%s(%s)", n.name, n.args.String())
}

type ASTMethod struct {
	name   string
	args   *ASTArguments
	parent ASTNode
}

func (n *ASTMethod) Accept(v ASTVisitor) {
	v.Visit(n)
}

func (n *ASTMethod) String() string {
	return fmt.Sprintf("%s.%s(%s)", n.parent.String(), n.name, n.args.String())
}

type ASTArguments struct {
	values []ASTNode
}

func (n *ASTArguments) Accept(v ASTVisitor) {
	v.Visit(n)
}

func (n *ASTArguments) String() string {
	var bld strings.Builder
	for i, a := range n.values {
		bld.WriteString(a.String())
		if i < len(n.values)-1 {
			bld.WriteString(", ")
		}
	}
	return bld.String()
}

type ASTLiteral struct {
	token *Token
}

func (n *ASTLiteral) Accept(v ASTVisitor) {
	v.Visit(n)
}

func (n *ASTLiteral) String() string {
	return fmt.Sprintf("%s{%s}", n.token.TokenType.String(), n.token.Value)
}

type ASTVisitor interface {
	Visit(n ASTNode)
}

func (p *Parser) Parse(scanner *Scanner) (ASTNode, error) {
	p.scanner = scanner.Scan()
	exp, err := p.Expression()
	if err != nil {
		return nil, err
	}
	if p.peek().TokenType != TOKEN_EOF {
		return nil, fmt.Errorf("unexpected token %s", p.consume())
	}
	return exp, nil
}

func (p *Parser) Expression() (ASTNode, error) {
	return p.Comparison()
}

func (p *Parser) Comparison() (ASTNode, error) {
	comp, err := p.Term()
	if err != nil {
		return nil, err
	}
	for p.match(TOKEN_EQUALS, TOKEN_NOT_EQUALS, TOKEN_GREATER_THAN, TOKEN_GREATER_THAN_EQUAL, TOKEN_LESS_THAN, TOKEN_LESS_THAN_EQUAL) {
		logInf.Printf("Found Comparison %s", p.peek())
		operator := p.consume()
		right, err := p.Term()
		if err != nil {
			return nil, err
		}
		fnName, err := builtInFuncName(operator.TokenType)
		if err != nil {
			return nil, err
		}
		comp = &ASTFunction{
			name: fnName,
			args: &ASTArguments{[]ASTNode{comp, right}},
		}
	}
	return comp, err
}

func (p *Parser) Term() (ASTNode, error) {
	term, err := p.Factor()
	if err != nil {
		return nil, err
	}
	for p.match(TOKEN_PLUS, TOKEN_MINUS) {
		logInf.Printf("Found Term %s", p.peek())
		operator := p.consume()
		right, err := p.Factor()
		if err != nil {
			return nil, err
		}
		fnName, err := builtInFuncName(operator.TokenType)
		if err != nil {
			return nil, err
		}
		term = &ASTFunction{
			name: fnName,
			args: &ASTArguments{[]ASTNode{term, right}},
		}
	}
	return term, err
}

func (p *Parser) Factor() (ASTNode, error) {
	term, err := p.Unary()
	if err != nil {
		return nil, err
	}
	for p.match(TOKEN_MULTIPLY, TOKEN_DIVIDE, TOKEN_MODULUS, TOKEN_POWER) {
		logInf.Printf("Found Factor %s", p.peek())
		operator := p.consume()
		right, err := p.Unary()
		if err != nil {
			return nil, err
		}
		fnName, err := builtInFuncName(operator.TokenType)
		if err != nil {
			return nil, err
		}
		term = &ASTFunction{
			name: fnName,
			args: &ASTArguments{[]ASTNode{term, right}},
		}
	}
	return term, err
}

func (p *Parser) Unary() (ASTNode, error) {
	if p.match(TOKEN_NOT, TOKEN_MINUS) {
		logInf.Printf("Found Unary %s", p.peek())
		operator := p.consume()
		operand, err := p.Group()
		if err != nil {
			return nil, err
		}
		fnName, err := builtInFuncName(operator.TokenType)
		if err != nil {
			return nil, err
		}
		args := &ASTArguments{make([]ASTNode, 0)}
		if operator.TokenType == TOKEN_MINUS {
			args.values = append(args.values, &ASTLiteral{&Token{TokenType: TOKEN_INT, Value: "0"}})
		}
		args.values = append(args.values, operand)
		return &ASTFunction{
				name: fnName,
				args: args,
			},
			nil
	}
	return p.Group()
}

func (p *Parser) Group() (ASTNode, error) {
	if p.match(TOKEN_START_ARGS) { //parentheses without function / method name
		logInf.Printf("Found Group %s", p.peek())
		p.consume() //consume start args
		args, err := p.Expression()
		if err != nil {
			return nil, err
		}
		if !p.match(TOKEN_END_ARGS) {
			return nil, fmt.Errorf("expected right parenthesis but got %s", p.peek())
		}
		p.consume() //consume end args
		return &ASTFunction{
				name: "nil",
				args: &ASTArguments{
					[]ASTNode{args},
				},
			},
			nil
	}
	return p.Ident(nil)
}

func (p *Parser) Ident(parent ASTNode) (ret ASTNode, err error) {
	if p.match(TOKEN_IDENT) {
		id := p.consume()
		logInf.Printf("Consumed %s", id)
		if p.match(TOKEN_START_ARGS) { //method or function call
			args, err := p.Arguments()
			if err != nil {
				return nil, err
			}
			if parent == nil { //function call
				ret =
					&ASTFunction{
						name: id.Value,
						args: args,
					}
			} else {
				//method call
				ret = &ASTMethod{
					name:   id.Value,
					args:   args,
					parent: parent,
				}
			}
		} else {
			//property
			ret = &ASTProperty{
				name:   id.Value,
				parent: parent,
			}
		}
		for p.match(TOKEN_SEPARATOR) {
			p.consume() //consume separator
			ret, err = p.Ident(ret)
		}
		return ret, err
	}
	return p.Literal()
}

func (p *Parser) Arguments() (*ASTArguments, error) {
	a := &ASTArguments{
		values: make([]ASTNode, 0),
	}
	for p.peek().TokenType == TOKEN_START_ARGS || p.peek().TokenType == TOKEN_DELIMITER {
		logInf.Printf("Consuming arg delimiter %s", p.peek())
		p.consume()                   //consume the start, delim or end
		if !p.match(TOKEN_END_ARGS) { //skip empty parentheses
			logInf.Printf("Parsing argument beginning with token %s", p.peek())
			arg, err := p.Expression()
			if err != nil {
				return nil, err
			}
			a.values = append(a.values, arg)
		}
	}
	if !p.match(TOKEN_END_ARGS) {
		return nil, fmt.Errorf("expected right parenthesis but got %s", p.peek())
	}
	p.consume() //consume TOKEN_DELIMITER or TOKEN_END_ARGS
	return a, nil
}

func (p *Parser) Literal() (ASTNode, error) {
	if p.match(TOKEN_BOOL, TOKEN_NIL, TOKEN_STRING, TOKEN_INT, TOKEN_FLOAT) {
		logInf.Printf("Found literal %s", p.peek())
		return &ASTLiteral{token: p.consume()}, nil
	}
	return nil, fmt.Errorf("unexpected token %s", p.consume())
}

//private utils
func (p *Parser) match(types ...TokenType) bool {
	if len(p.scanner.Tokens) > p.pos {
		for _, tt := range types {
			if p.scanner.Tokens[p.pos].TokenType == tt {
				return true
			}
		}
	}
	return false
}

func (p *Parser) consume() (t *Token) {
	if len(p.scanner.Tokens) > p.pos {
		t = p.scanner.Tokens[p.pos]
		p.pos++
		return
	}
	return &Token{TokenType: TOKEN_UNKNOWN}
}

func (p *Parser) peek() *Token {
	return p.peekAhead(0)
}

func (p *Parser) peekAhead(count int) *Token {
	if len(p.scanner.Tokens) == 0 {
		return &Token{TokenType: TOKEN_UNKNOWN}
	}
	if len(p.scanner.Tokens) <= p.pos+count {
		return p.scanner.Tokens[len(p.scanner.Tokens)-1]
	}
	return p.scanner.Tokens[p.pos+count]
}

func (p *Parser) peekBehind(count int) *Token {
	return p.peekAhead(count * -1)
}

var builtInFuncMap = map[TokenType]string{
	TOKEN_PLUS:               "addOrConcat",
	TOKEN_MINUS:              "subtract",
	TOKEN_MULTIPLY:           "multiply",
	TOKEN_DIVIDE:             "divide",
	TOKEN_POWER:              "pow",
	TOKEN_MODULUS:            "mod",
	TOKEN_EQUALS:             "equals",
	TOKEN_NOT_EQUALS:         "notEquals",
	TOKEN_GREATER_THAN:       "greaterThan",
	TOKEN_GREATER_THAN_EQUAL: "greaterThanEqual",
	TOKEN_LESS_THAN:          "lessThan",
	TOKEN_LESS_THAN_EQUAL:    "lessThanEqual",
	TOKEN_AND:                "and",
	TOKEN_OR:                 "or",
	TOKEN_NOT:                "not",
}

func builtInFuncName(tokenType TokenType) (string, error) {
	if fn, ok := builtInFuncMap[tokenType]; ok {
		return fn, nil
	}
	return "", fmt.Errorf("%s is not mapped to a built in function", tokenType)
}
