package toknizer

type Token struct {
	Type  string
	Value string
}

type tokenSpec struct {
	RegexStr   string
	Name       string
	Identifier string
}

var (
	Number = tokenSpec{
		RegexStr:   `^\d+`,
		Name:       "NUMBER",
		Identifier: "number",
	}
	DoubleQuoteString = tokenSpec{
		RegexStr:   `^"([^"]*)"`,
		Name:       "STRING",
		Identifier: "string",
	}
	SingleQuoteString = tokenSpec{
		RegexStr:   `^'([^']*)'`,
		Name:       "STRING",
		Identifier: "string",
	}

	// Skipped tokens
	WhiteSpace = tokenSpec{
		RegexStr:   `^\s+`,
		Name:       "",
		Identifier: "white_space",
	}
	LineBreak = tokenSpec{
		RegexStr:   `^\n+`,
		Name:       "",
		Identifier: "line_break",
	}

	LeftBracket = tokenSpec{
		RegexStr:   `^{`,
		Name:       "{",
		Identifier: "left_bracket",
	}
	RightBracket = tokenSpec{
		RegexStr:   `^}`,
		Name:       "}",
		Identifier: "right_bracket",
	}

	Tokens []tokenSpec = []tokenSpec{
		DoubleQuoteString,
		SingleQuoteString,
		Number,
		WhiteSpace,
		LineBreak,
		LeftBracket,
		RightBracket,
	}
)
