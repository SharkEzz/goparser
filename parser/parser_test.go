package parser_test

import (
	"encoding/json"
	"strings"
	"testing"

	ps "github.com/SharkEzz/goparser/parser"
)

var parser = ps.Parser{}

func TestAST(t *testing.T) {
	const testProgram string = `
		"oui";
		42;
		'non';
		"blabla";
	`

	ast, err := parser.Parse(testProgram)

	if ast.Type != "PROGRAM" {
		t.Errorf(`Invalid AST program type, expected "PROGRAM" but received "%s"`, ast.Type)
	}

	if len(ast.Body) != 4 {
		t.Errorf(`Invalid AST body length, expected "4" but received "%d"`, len(ast.Body))
	}

	astJson, err := json.Marshal(ast)
	if err != nil {
		t.Error(err)
	}

	const expectedJson = `{"Type":"PROGRAM","Body":[{"Type":"ExpressionStatement","Expression":{"Type":"LiteralExpression","Value":"oui"}},{"Type":"ExpressionStatement","Expression":{"Type":"NumericExpression","Value":"42"}},{"Type":"ExpressionStatement","Expression":{"Type":"LiteralExpression","Value":"non"}},{"Type":"ExpressionStatement","Expression":{"Type":"LiteralExpression","Value":"blabla"}}]}`

	if strings.TrimSpace(expectedJson) != string(astJson) {
		t.Error("JSON does not match")
	}
}
