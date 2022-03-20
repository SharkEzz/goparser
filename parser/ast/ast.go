package ast

type Literal struct {
	Type  string
	Value string
}

type Expression struct {
	Type       string
	Expression any
}

type Node struct {
	Type  string
	Value any
}

type Program struct {
	Type string
	Body []Expression
}
