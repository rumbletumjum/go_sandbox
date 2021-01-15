package main

import (
	"fmt"
	"regexp"
	"awesomeProject/reader"
	"strings"
)

type Atom interface{}
type List []Atom
type LispType interface{}

func tokenize(input string) []string {
	var res string

	//left := "\\s*\\(\\s*"
	left := `\(`
	//right := "\\s*\\)\\s*"
	right := `\)`

	left_re := regexp.MustCompile(left)
	right_re := regexp.MustCompile(right)

	res = left_re.ReplaceAllString(input, "( ")
	res = right_re.ReplaceAllString(res, " )")
	//res = strings.Trim(res, " ")
	tokens := strings.Split(res, " ")
	return tokens
}

func readTokens(rdr *reader.TokenReader) interface{} {
	token := rdr.Pop()
	if *token == "(" {
		var L []interface{}
		for *rdr.Peek() != ")" {
			L = append(L, readTokens(rdr))
		}
		rdr.Pop()
		return L
	} else {
		return *token
	}
}

func readFromTokens(rdr *reader.TokenReader) []LispType {
	var outerList []LispType
	token := rdr.Next()
	if *token == "(" {
		var list []LispType
		token = rdr.Next()
		for; true; token = rdr.Next() {
			if *token == ")" {
				break
			}
			list = append(list, *token)
		}
		outerList = append(outerList, list)
	}
	outerList = append(outerList, *token)
	return outerList
}

func main() {
	//src := "(* (+ 1 2) 6)"
	program := "(begin (define r 10) (* pi (* r r)))"
	tokens := tokenize(program)
	rdr := &reader.TokenReader{
		Tokens:   tokens,
		Position: 0,
	}
	readTokens(rdr)
	var whatever []interface{}
	var subslice []interface{}
	whatever = append(whatever, 7)
	subslice = append(subslice, 42, 24)
	whatever = append(whatever, subslice)

	fmt.Println("break")

}
