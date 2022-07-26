package main

type Statement interface {
	Accept(visitor StatementVisitor) Any
}

type StatementVisitor interface {
	visitPrintStmt(stmt PrintStatement) Any
	visitExprStmt(stmt ExpressionStatement) Any
	visitVarStmt(stmt VarStatement) Any
	visitBlockStmt(stmt BlockStatement) Any
	visitIfStmt(stmt IfStatement) Any
	visitWhileStmt(stmt WhileStatement) Any
	visitFunctionStmt(stmt FunctionStatement) Any
	visitReturnStmt(stmt ReturnStatement) Any
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

type VarStatement struct {
	Name        Token
	Initializer Expression
}

func (e VarStatement) Accept(visitor StatementVisitor) Any {
	return visitor.visitVarStmt(e)
}

type BlockStatement struct {
	Statements []Statement
}

func (e BlockStatement) Accept(visitor StatementVisitor) Any {
	return visitor.visitBlockStmt(e)
}

type IfStatement struct {
	Condition Expression
	ThenBlock Statement
	ElseBlock Statement
}

func (b IfStatement) Accept(visitor StatementVisitor) Any {
	return visitor.visitIfStmt(b)
}

type WhileStatement struct {
	Condition Expression
	Body      Statement
}

func (b WhileStatement) Accept(visitor StatementVisitor) Any {
	return visitor.visitWhileStmt(b)
}

type FunctionStatement struct {
	Name   Token
	Params []Token
	Body   []Statement
}

func (b FunctionStatement) Accept(visitor StatementVisitor) Any {
	return visitor.visitFunctionStmt(b)
}

type ReturnStatement struct {
	Keyword Token
	Value   Expression
}

func (b ReturnStatement) Accept(visitor StatementVisitor) Any {
	return visitor.visitReturnStmt(b)
}
