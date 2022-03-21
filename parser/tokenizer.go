package parser

import (
	"fmt"
	"regexp"

	"github.com/SharkEzz/goparser/parser/types/toknizer"
)

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
		// return the value in itself (omit quotes and so)
		return submatches[1]
	}

	return matched
}

func (t *Tokenizer) GetNextToken() (*toknizer.Token, error) {
	if !t.hasMoreTokens() {
		return nil, fmt.Errorf("No more tokens to process")
	}

	str := string(t.input[t.cursor:])

	for _, spec := range toknizer.Tokens {
		tokenValue := t.match(spec.RegexStr, str)

		if tokenValue == "" {
			continue
		}

		if spec.Name == "" {
			return t.GetNextToken()
		}

		return &toknizer.Token{
			Type:  spec.Name,
			Value: tokenValue,
		}, nil
	}

	return nil, fmt.Errorf(`Syntax error: unexpected "%s" at position %d`, string(str[0]), t.cursor)
}
