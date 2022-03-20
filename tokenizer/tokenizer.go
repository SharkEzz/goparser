package tokenizer

import (
	"fmt"
	"regexp"
	"strings"
)

var specs = [][2]string{
	{`^\d+`, "NUMBER"},
	{`^"([^"]*)"`, "STRING"},
	{`^'([^']*)'`, "STRING"},
	{`^\s+`, ""},
	{`^;`, ";"},
}

type Tokenizer struct {
	input  string
	cursor int
}

func (t *Tokenizer) isEOF() bool {
	return t.cursor == len(t.input)
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

	matched := regex.FindAllString(str, -1)
	submatches := regex.FindStringSubmatch(str)

	t.cursor += len(matched[0])

	if len(submatches) > 1 {
		return submatches[1]
	}

	return matched[0]
}

func (t *Tokenizer) GetNextToken() (*Token, error) {
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

		return &Token{
			Type:  spec[1],
			Value: tokenValue,
		}, nil
	}

	return nil, fmt.Errorf(`Syntax error: unexpected "%s"`, str)
}
