package main

type Statement interface {
	Accept(visitor StatementVisitor) Any
}

type StatementVisitor interface {
	visitPrintStmt(expr PrintStatement) Any
	visitExprStmt(expr ExpressionStatement) Any
}

type PrintStatement struct {
	Expression Expression
}

func (p PrintStatement) Accept(visitor StatementVisitor) Any {
	return visitor.visitPrintStmt(p)
}

type ExpressionStatement struct {
	Expression Expression
}

func (e ExpressionStatement) Accept(visitor StatementVisitor) Any {
	return visitor.visitExprStmt(e)
}
