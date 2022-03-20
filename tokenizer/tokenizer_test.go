package tokenizer_test

import (
	"testing"

	tk "github.com/SharkEzz/goparser/tokenizer"
)

var tokenizer = tk.Tokenizer{}

func TestNumeric(t *testing.T) {
	const testData string = "42"

	tokenizer.Init(testData)
	token, err := tokenizer.GetNextToken()

	if err != nil {
		t.Error(err)
	}

	if token.Type != "NUMBER" || token.Value != "42" {
		t.Errorf(`Invalid token data, expected "NUMBER" but received "%s" for Type and expected "42" but received "%s" for value`, token.Type, token.Value)
	}
}

func TestString(t *testing.T) {
	const testData string = `"42"`

	tokenizer.Init(testData)
	token, err := tokenizer.GetNextToken()

	if err != nil {
		t.Error(err)
	}

	if token.Type != "STRING" || token.Value != "42" {
		t.Errorf(`Invalid token data, expected "STRING" but received "%s" for Type and expected "42" but received "%s" for value`, token.Type, token.Value)
	}

	const testDataSingleQuote string = `'42'`

	tokenizer.Init(testDataSingleQuote)
	token, err = tokenizer.GetNextToken()

	if err != nil {
		t.Error(err)
	}

	if token.Type != "STRING" || token.Value != "42" {
		t.Errorf(`Invalid token data, expected "STRING" but received "%s" for Type and expected "42" but received "%s" for value`, token.Type, token.Value)
	}
}

func TestLineBreak(t *testing.T) {
	const testData string = `
		24
				"test"

		"OUI"



		12
	`

	tokenizer.Init(testData)
	token, err := tokenizer.GetNextToken()

	if err != nil {
		t.Error(err)
	}

	if token.Type != "NUMBER" || token.Value != "24" {
		t.Errorf("Invalid token: %s", token.Type)
	}

	token, err = tokenizer.GetNextToken()

	if err != nil {
		t.Error(err)
	}

	if token.Type != "STRING" || token.Value != "test" {
		t.Errorf("Invalid token: %s, value: %s", token.Type, token.Value)
	}

	token, err = tokenizer.GetNextToken()

	if err != nil {
		t.Error(err)
	}

	if token.Type != "STRING" || token.Value != "OUI" {
		t.Errorf("Invalid token: %s, value: %s", token.Type, token.Value)
	}

	token, err = tokenizer.GetNextToken()

	if err != nil {
		t.Error(err)
	}

	if token.Type != "NUMBER" || token.Value != "12" {
		t.Errorf("Invalid token: %s, value: %s", token.Type, token.Value)
	}
}
