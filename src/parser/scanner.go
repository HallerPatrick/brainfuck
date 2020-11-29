package parser

import (
	"io/ioutil"
	"strings"
)

const (
	IncrementData = "+"
	DecrementData = "-"
	IncrementPointer = ">"
	DecrementPointer = "<"
	LeftBracket = "["
	RightBracket = "]"
	Print = "."
)

type CommandToken int

const (
	IncrementDataToken CommandToken = iota
	DecrementDataToken
	IncrementPointerToken
	DecrementPointerToken
	LeftBracketToken
	RightBracketToken
	PrintToken
)


var CommandLiterals = [7]string {
	IncrementData,
	DecrementData,
	IncrementPointer,
	DecrementPointer,
	LeftBracket,
	RightBracket,
	Print,
}

var Commands = map[string]CommandToken {
	IncrementData: IncrementDataToken,
	DecrementData: DecrementDataToken,
	IncrementPointer: IncrementPointerToken,
	DecrementPointer: DecrementPointerToken,
	LeftBracket: LeftBracketToken,
	RightBracket: RightBracketToken,
	Print: PrintToken,
}

type Stack struct {
	source []string
}

func initStack(source string) Stack {
	return Stack{strings.Split(source, "")}
}

func (stack *Stack) pop() string {
	elem := stack.source[len(stack.source)-1]
	stack.source = stack.source[:len(stack.source)-1]
	return elem
}

func (stack *Stack) isEmpty() bool {
	return len(stack.source) == 0
}

func ParseFromFile(filename string) []CommandToken {
	source, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return Parse(initStack(string(source)))
}

func isCommand(char string) bool {
	for _, command := range CommandLiterals {
		if command == char {
			return true
		}
	}
	return false
}

func Parse(source Stack) []CommandToken {
	var tokens []CommandToken

	for !source.isEmpty() {
		elem := source.pop()
		if isCommand(elem) {
			token := Commands[elem]
			tokens = append(tokens, token)
		}
	}

	return tokens
}
