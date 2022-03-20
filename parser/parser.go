package parser

import (
	"fmt"

	"github.com/SharkEzz/goparser/parser/ast"
	tk "github.com/SharkEzz/goparser/tokenizer"
)

type Parser struct {
	tokenizer tk.Tokenizer
	lookahead *tk.Token
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
		Body: p.statements(),
	}, nil
}

func (p *Parser) eat(tokenType string) *tk.Token {
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

func (p *Parser) statements() []ast.Expression {
	var nodes []ast.Expression

	for p.lookahead != nil {
		nodes = append(nodes, p.generateExpression())
	}

	return nodes
}

func (p *Parser) generateExpression() ast.Expression {
	expression := p.literal()

	return ast.Expression{
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
