package main

import (
	"github.com/HallerPatrick/brainfuck/src/interpreter"
	"github.com/HallerPatrick/brainfuck/src/parser"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("Please provide a source file.")
	}
	var sourceFile = os.Args[1]

	tokens := parser.ScanFromFile(sourceFile)
	tokenParser := parser.New(tokens)
	tokenParser.Parse()

	interpr := interpreter.New(tokenParser.Ast)
	interpr.Interpret()
}
