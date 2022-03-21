package parser

import (
	"fmt"

	"github.com/SharkEzz/goparser/parser/types/ast"
	"github.com/SharkEzz/goparser/parser/types/token"
)

type Parser struct {
	tokenizer Tokenizer
	lookahead *token.Token
}

func (p *Parser) Parse(input string) (*ast.Program, error) {
	p.tokenizer.Init(input)

	token, err := p.tokenizer.GetNextToken()
	if err != nil {
		panic(err)
	}

	p.lookahead = token

	return &ast.Program{
		Type: "PROGRAM",
		Body: p.statementList(""),
	}, nil
}

func (p *Parser) eat(tokenType string) *token.Token {
	token := p.lookahead

	if token == nil {
		panic(fmt.Errorf(`Unexpected EOF, expected "%s"`, tokenType))
	}

	if token.Type != tokenType {
		panic(fmt.Errorf(`Unexpected token: "%s", expected: "%s"`, token.Type, tokenType))
	}

	p.lookahead, _ = p.tokenizer.GetNextToken()

	return token
}

func (p *Parser) statementList(stopLookAhead string) *[]ast.Expression {
	nodes := []ast.Expression{*p.statement()}

	for p.lookahead != nil && p.lookahead.Type != stopLookAhead {
		nodes = append(nodes, *p.statement())
	}

	return &nodes
}

func (p *Parser) statement() *ast.Expression {
	switch p.lookahead.Type {
	case "{":
		return p.blockStatement()
	default:
		return p.expressionStatement()
	}
}

func (p *Parser) blockStatement() *ast.Expression {
	p.eat("{")

	body := &[]ast.Expression{}

	if p.lookahead != nil && p.lookahead.Type != "}" {
		body = p.statementList("}")
	}

	p.eat("}")

	return &ast.Expression{
		Type:       "BlockStatement",
		Expression: body,
	}
}

func (p *Parser) expressionStatement() *ast.Expression {
	expression := p.literal()

	return &ast.Expression{
		Type:       "ExpressionStatement",
		Expression: expression,
	}
}

func (p *Parser) literal() ast.Literal {
	switch p.lookahead.Type {
	case "NUMBER":
		token := p.eat("NUMBER")
		return ast.Literal{
			Type:  "NumericExpression",
			Value: token.Value.(string),
		}
	case "STRING":
		token := p.eat("STRING")
		return ast.Literal{
			Type:  "LiteralExpression",
			Value: token.Value.(string),
		}
	default:
		panic(fmt.Errorf("Unexpected literal"))
	}
}
