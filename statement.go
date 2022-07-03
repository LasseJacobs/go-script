package main

type Statement interface {
	Accept(visitor StatementVisitor) Any
}

type StatementVisitor interface {
	visitPrintStmt(expr PrintStatement) Any
	visitExprStmt(expr ExpressionStatement) Any
	visitVarStmt(expr VarStatement) Any
	visitBlockStmt(expr BlockStatement) Any
	visitIfStmt(expr IfStatement) Any
	visitWhileStmt(expr WhileStatement) Any
	visitFunctionStmt(expr FunctionStatement) Any
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

/*
"Function   : Token name, List<Token> params," +
                  " List<Stmt> body"
*/

type FunctionStatement struct {
	Name   Token
	Params []Token
	Body   []Statement
}

func (b FunctionStatement) Accept(visitor StatementVisitor) Any {
	return visitor.visitFunctionStmt(b)
}
