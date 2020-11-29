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

func initStack(source string) Stack {
	return NewSourceStack(strings.Split(source, ""))
}

func ScanFromFile(filename string) []CommandToken {
	source, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return Scan(initStack(string(source)))
}

func isCommand(char string) bool {
	for _, command := range CommandLiterals {
		if command == char {
			return true
		}
	}
	return false
}

func Scan(source Stack) []CommandToken {
	var tokens []CommandToken

	for !source.IsEmpty() {
		elem := source.Pop().(string)
		if isCommand(elem) {
			token := Commands[elem]
			tokens = append(tokens, token)
		}
	}

	return tokens
}
