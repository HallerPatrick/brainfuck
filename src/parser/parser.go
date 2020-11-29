package parser

import (
	"fmt"
)

type Expr int

const (
	IncrementDataExpr Expr = iota
	DecrementDataExpr
	IncrementPointerExpr
	DecrementPointerExpr
	LoopExpr
	PrintExpr
)


type Parser struct {
	token Stack
	Ast   []ExprNode
}

func (parser *Parser) PrintTree() {
	for _, expr := range parser.Ast {
		fmt.Println(expr)
	}
}


func (parser *Parser) GetTree() []ExprNode {
	return parser.Ast
}

func (parser *Parser) Parse() {
	for !parser.token.IsEmpty() {
		token := parser.token.Pop().(CommandToken);
		exprNode := parser.parseToken(token)
		parser.Ast = append(parser.Ast, exprNode)
	}
}

func (parser *Parser) parseToken(token CommandToken) ExprNode {
	switch  token {
	case LeftBracketToken:
		return parser.eatLoop()
	case IncrementPointerToken:
		return ExprNode{Expr: IncrementPointerExpr,}
	case DecrementPointerToken:
		return ExprNode{
			Expr: DecrementPointerExpr,
		}
	case IncrementDataToken:
		return ExprNode{
			Expr: IncrementDataExpr,
		}
	case DecrementDataToken:
		return ExprNode{
			Expr: DecrementDataExpr,
		}
	case PrintToken:
		return ExprNode{
			Expr: PrintExpr,
		}
	default:
		panic("Unkown Token Found!")
	}
}

func (parser *Parser) eatLoop() ExprNode {
	currentToken := parser.token.Pop().(CommandToken)
	loopExprs := make([]ExprNode, 0)
	for currentToken != RightBracketToken {
		exprNode := parser.parseToken(currentToken)
		loopExprs = append(loopExprs, exprNode)
		currentToken = parser.token.Pop().(CommandToken)
	}
	return ExprNode{
		Expr:  LoopExpr,
		Exprs: loopExprs,
	}
}

func (parser *Parser) addExpr(exprNode ExprNode) {
	parser.Ast = append(parser.Ast, exprNode)
}


func New(tokens []CommandToken) Parser {
	return Parser{
		token: NewTokenStack(tokens),
	}
}

