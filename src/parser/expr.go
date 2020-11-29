package parser

// ExprNode is either a command or a loop closure with a list of ExprNode
type ExprNode struct {
	Expr      Expr
	Exprs []ExprNode
}
