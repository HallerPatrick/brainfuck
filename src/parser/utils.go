package parser

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

func NewTokenStack(list []CommandToken) Stack {
	commandList := make([]interface{}, 0)
	for _, elem := range list {
		commandList = append(commandList, elem)
	}
	return Stack{stack: commandList}
}


func NewExprStack(list []ExprNode) Stack {
	exprList := make([]interface{}, 0)
	for _, elem := range list {
		exprList = append(exprList, elem)
	}
	return Stack{stack: exprList}
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.stack) == 0
}

// FIFO
func (stack *Stack) Pop() interface{} {
	if len(stack.stack) == 0 {
		panic("Reach EOF while parsing")
	}
	elem := stack.stack[0]
	stack.stack = stack.stack[1:]
	return elem
}
