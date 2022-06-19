package main

type Expression interface {
	Accept(visitor ExpressionVisitor) string
}

type ExpressionVisitor interface {
	visitBinaryExpr(expr BinaryExpression) string
	visitGroupingExpr(expr GroupingExpression) string
	visitLiteralExpr(expr LiteralExpression) string
	visitUnaryExpr(expr UnaryExpression) string
}

type BinaryExpression struct {
	Left     Expression
	Operator Token
	Right    Expression
}

func (b BinaryExpression) Accept(visitor ExpressionVisitor) string {
	return visitor.visitBinaryExpr(b)
}

type GroupingExpression struct {
	Expression Expression
}

func (b GroupingExpression) Accept(visitor ExpressionVisitor) string {
	return visitor.visitGroupingExpr(b)
}

type LiteralExpression struct {
	Value interface{}
}

func (b LiteralExpression) Accept(visitor ExpressionVisitor) string {
	return visitor.visitLiteralExpr(b)
}

type UnaryExpression struct {
	Operator Token
	Right    Expression
}

func (b UnaryExpression) Accept(visitor ExpressionVisitor) string {
	return visitor.visitUnaryExpr(b)
}
