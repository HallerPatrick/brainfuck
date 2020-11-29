package interpreter

import (
	"fmt"
	"github.com/HallerPatrick/brainfuck/src/parser"
)

type State struct {
	// All cells initialized to zero
	array [30]byte

	// Current pointer position in program
	pointerPosition int

}

func (state *State) getCurrentCell() byte {
	return state.array[state.pointerPosition]
}

func New(tree []parser.ExprNode) Interpreter {
	return Interpreter {
		State:  State{
			pointerPosition: 0,
		},
		exprStack: parser.NewExprStack(tree),
	}
}

type Interpreter struct {
	State State
	exprStack parser.Stack
}

func (interpreter *Interpreter) PrintArray() {
	fmt.Println(interpreter.State.array)
}

func (interpreter *Interpreter) incrementData() {
	interpreter.State.array[interpreter.State.pointerPosition] = interpreter.State.array[interpreter.State.pointerPosition] + 1
}

func (interpreter *Interpreter) decrementData() {
	interpreter.State.array[interpreter.State.pointerPosition]--
}

func (interpreter *Interpreter) incrementPoiner() {
	interpreter.State.pointerPosition++
}

func (interpreter *Interpreter) decrementPointer() {
	interpreter.State.pointerPosition--
}

func (interpreter *Interpreter) printCell() {
	fmt.Printf("%c", rune(interpreter.State.getCurrentCell()))
}

func (interpreter *Interpreter) executeLoop(exprNode parser.ExprNode) {
	for interpreter.State.getCurrentCell() != 0 {
		for _, exprNode := range exprNode.Exprs {
			interpreter.executeCommand(exprNode)
		}
	}
}

func (interpreter *Interpreter) executeCommand(exprNode parser.ExprNode) {
	switch exprNode.Expr {
	case parser.IncrementDataExpr:
		interpreter.incrementData()

	case parser.DecrementDataExpr:
		interpreter.decrementData()

	case parser.IncrementPointerExpr:
		interpreter.incrementPoiner()

	case parser.DecrementPointerExpr:
		interpreter.decrementPointer()

	case parser.PrintExpr:
		interpreter.printCell()

	case parser.LoopExpr:
		interpreter.executeLoop(exprNode)
	}
}

func (interpreter *Interpreter) Interpret() {
	for !interpreter.exprStack.IsEmpty() {
		exprNode := interpreter.exprStack.Pop().(parser.ExprNode)
		interpreter.executeCommand(exprNode)
	}
}



