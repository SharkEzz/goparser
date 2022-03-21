package parser_test

import (
	"encoding/json"
	"testing"

	ps "github.com/SharkEzz/goparser/parser"
)

var parser = ps.Parser{}

func TestAST(t *testing.T) {
	const testProgram string = `
		"oui"
		42
		'non'
		
		"blabla"
	`

	ast, err := parser.Parse(testProgram)
	if err != nil {
		t.Error(err)
	}

	if ast.Type != "PROGRAM" {
		t.Errorf(`Invalid AST program type, expected "PROGRAM" but received "%s"`, ast.Type)
	}

	if len(*ast.Body) != 4 {
		t.Errorf(`Invalid AST body length, expected "4" but received "%d"`, len(*ast.Body))
	}

	astJson, err := json.Marshal(ast)
	if err != nil {
		t.Error(err)
	}

	const expectedJson string = `{"Type":"PROGRAM","Body":[{"Type":"ExpressionStatement","Expression":{"Type":"LiteralExpression","Value":"oui"}},{"Type":"ExpressionStatement","Expression":{"Type":"NumericExpression","Value":"42"}},{"Type":"ExpressionStatement","Expression":{"Type":"LiteralExpression","Value":"non"}},{"Type":"ExpressionStatement","Expression":{"Type":"LiteralExpression","Value":"blabla"}}]}`

	if expectedJson != string(astJson) {
		t.Error("JSON does not match")
	}
}

func TestBlock(t *testing.T) {
	const testProgram string = `
		{
			42

			"hello"
		}
	`

	ast, err := parser.Parse(testProgram)

	if err != nil {
		t.Error(err)
	}

	astJson, err := json.Marshal(ast)
	if err != nil {
		t.Error(err)
	}

	const expectedJson string = `{"Type":"PROGRAM","Body":[{"Type":"BlockStatement","Expression":[{"Type":"ExpressionStatement","Expression":{"Type":"NumericExpression","Value":"42"}},{"Type":"ExpressionStatement","Expression":{"Type":"LiteralExpression","Value":"hello"}}]}]}`

	if string(astJson) != expectedJson {
		t.Error("JSON does not match")
	}
}

func TestEmptyBlock(t *testing.T) {
	const testProgram string = `{
	}
	`

	ast, err := parser.Parse(testProgram)

	if err != nil {
		t.Error(err)
	}

	astJson, err := json.Marshal(ast)
	if err != nil {
		t.Error(err)
	}

	const expectedJson string = `{"Type":"PROGRAM","Body":[{"Type":"BlockStatement","Expression":[]}]}`

	if string(astJson) != expectedJson {
		t.Error("JSON does not match")
	}
}

func TestInvalidBlock(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	const testProgram string = `{
	`

	_, err := parser.Parse(testProgram)
	if err != nil {
		t.Error(err)
	}
}

func TestNestedBlocks(t *testing.T) {
	const testProgram string = `
		{
			42

			"hello"
			{
				"nested"
				41
			}
		}
	`

	ast, err := parser.Parse(testProgram)
	if err != nil {
		t.Error(err)
	}

	astJson, err := json.Marshal(ast)
	if err != nil {
		t.Error(err)
	}

	const expectedJson string = `{"Type":"PROGRAM","Body":[{"Type":"BlockStatement","Expression":[{"Type":"ExpressionStatement","Expression":{"Type":"NumericExpression","Value":"42"}},{"Type":"ExpressionStatement","Expression":{"Type":"LiteralExpression","Value":"hello"}},{"Type":"BlockStatement","Expression":[{"Type":"ExpressionStatement","Expression":{"Type":"LiteralExpression","Value":"nested"}},{"Type":"ExpressionStatement","Expression":{"Type":"NumericExpression","Value":"41"}}]}]}]}`

	if expectedJson != string(astJson) {
		t.Error("JSON does not match")
	}
}
