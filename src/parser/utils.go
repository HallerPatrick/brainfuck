package utils

import "github.com/HallerPatrick/brainfuck/src/parser"

type Stack struct {
	stack []interface{}
}

func NewSourceStack(list []string) Stack {
	stackList := make([]interface{}, 0)
	for _, elem := range list {
		stackList = append(stackList, elem)
	}
	return Stack{stack: stackList}
}

func NewTokenStack(list []parser.CommandToken) Stack {
	commandList := make([]interface{}, 0)
	for _, elem := range list {
		commandList = append(commandList, elem)
	}
	return Stack{stack: commandList}
}


func NewExprStack(list []parser.ExprNode) Stack {
	exprList := make([]interface{}, 0)
	for _, elem := range list {
		exprList = append(exprList, elem)
	}
	return Stack{stack: exprList}
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.stack) == 0
}

func (stack *Stack) Pop() interface{} {
	if len(stack.stack) == 0 {
		panic("Reach EOF while parsing")
	}
	elem := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return elem
}
