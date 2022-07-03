package main

type Expression interface {
	Accept(visitor ExpressionVisitor) Any
}

type ExpressionVisitor interface {
	visitBinaryExpr(expr BinaryExpression) Any
	visitGroupingExpr(expr GroupingExpression) Any
	visitLiteralExpr(expr LiteralExpression) Any
	visitUnaryExpr(expr UnaryExpression) Any
	visitVarExpr(expr VariableExpression) Any
	visitAssignExpr(expr AssignExpression) Any
	visitCallExpr(expr CallExpression) Any
}

type BinaryExpression struct {
	Left     Expression
	Operator Token
	Right    Expression
}

func (b BinaryExpression) Accept(visitor ExpressionVisitor) Any {
	return visitor.visitBinaryExpr(b)
}

type GroupingExpression struct {
	Expression Expression
}

func (b GroupingExpression) Accept(visitor ExpressionVisitor) Any {
	return visitor.visitGroupingExpr(b)
}

type LiteralExpression struct {
	Value interface{}
}

func (b LiteralExpression) Accept(visitor ExpressionVisitor) Any {
	return visitor.visitLiteralExpr(b)
}

type UnaryExpression struct {
	Operator Token
	Right    Expression
}

func (b UnaryExpression) Accept(visitor ExpressionVisitor) Any {
	return visitor.visitUnaryExpr(b)
}

type VariableExpression struct {
	Name Token
}

func (b VariableExpression) Accept(visitor ExpressionVisitor) Any {
	return visitor.visitVarExpr(b)
}

type AssignExpression struct {
	Name  Token
	Value Expression
}

func (b AssignExpression) Accept(visitor ExpressionVisitor) Any {
	return visitor.visitAssignExpr(b)
}

type CallExpression struct {
	Callee    Expression
	Paren     Token
	Arguments []Expression
}

func (b CallExpression) Accept(visitor ExpressionVisitor) Any {
	return visitor.visitCallExpr(b)
}
