package main

import "fmt"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) tokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token Token
	Value string
}

type Widget struct {
	Name       string
}

func main() {
	tok := Identifier{Token{"IDENT", "let"}, "let"}
	stmt := LetStatement{}

	widget := Widget{"wicket"}
	id := widget.(*Name)

	prog := Program{}

	statements := append(prog.Statements, stmt)

	fmt.Printf("%s", statements)

}
