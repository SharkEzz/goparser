package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/SharkEzz/goparser/parser/types/token"
)

var specs = [][2]string{
	{`^\d+`, "NUMBER"},
	{`^"([^"]*)"`, "STRING"},
	{`^'([^']*)'`, "STRING"},
	{`^\s+`, ""},
	{`^\n+`, ""},
	{`^{`, "{"},
	{`^}`, "}"},
}

type Tokenizer struct {
	input  string
	cursor int
}

func (t *Tokenizer) hasMoreTokens() bool {
	return t.cursor < len(t.input)
}

func (t *Tokenizer) Init(input string) {
	t.input = input
	t.cursor = 0
}

func (t *Tokenizer) match(regexStr, str string) string {
	regex := regexp.MustCompile(regexStr)

	isMatched := regex.MatchString(str)

	if !isMatched {
		return ""
	}

	matched := regex.FindString(str)
	submatches := regex.FindStringSubmatch(str)

	t.cursor += len(matched)

	if len(submatches) > 1 {
		// return the value in itself
		return submatches[1]
	}

	return matched
}

func (t *Tokenizer) GetNextToken() (*token.Token, error) {
	if !t.hasMoreTokens() {
		return nil, fmt.Errorf("No more tokens to process")
	}

	strSlice := strings.Split(t.input, "")
	str := strings.Join(strSlice[t.cursor:], "")

	for _, spec := range specs {
		tokenValue := t.match(spec[0], str)

		if tokenValue == "" {
			continue
		}

		if spec[1] == "" {
			return t.GetNextToken()
		}

		return &token.Token{
			Type:  spec[1],
			Value: tokenValue,
		}, nil
	}

	return nil, fmt.Errorf(`Syntax error: unexpected "%s"`, str)
}
