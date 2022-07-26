package main

type Resolver struct {
	interpreter *Interpreter
	scopes      *Stack
}

func NewResolver(interpreter *Interpreter) *Resolver {
	return &Resolver{interpreter: interpreter, scopes: NewStack()}
}

func (r *Resolver) Resolve(statements []Statement) Any {
	for _, statement := range statements {
		r.resolveStmt(statement)
	}
	return nil
}

func (r *Resolver) visitBinaryExpr(expr BinaryExpression) Any {
	r.resolveExpr(expr.Left)
	r.resolveExpr(expr.Right)
	return nil
}

func (r *Resolver) visitGroupingExpr(expr GroupingExpression) Any {
	r.resolveExpr(expr.Expression)
	return nil
}

func (r *Resolver) visitLiteralExpr(expr LiteralExpression) Any {
	return nil
}

func (r *Resolver) visitUnaryExpr(expr UnaryExpression) Any {
	r.resolveExpr(expr.Right)
	return nil
}

func (r *Resolver) visitVarExpr(expr VariableExpression) Any {
	if !r.scopes.IsEmpty() && r.scopes.Peek()[expr.Name.Lexeme] == false {
		parseFault(expr.Name, "Can't read local variable in its own initializer.")
	}
	r.resolveLocal(expr, expr.Name)
	return nil
}

func (r *Resolver) visitAssignExpr(expr AssignExpression) Any {
	r.resolveExpr(expr.Value)
	r.resolveLocal(expr, expr.Name)
	return nil
}

func (r *Resolver) visitCallExpr(expr CallExpression) Any {
	r.resolveExpr(expr.Callee)
	for _, arg := range expr.Arguments {
		r.resolveExpr(arg)
	}
	return nil
}

func (r *Resolver) resolveExpr(expr Expression) Any {
	expr.Accept(r)
	return nil
}

func (r *Resolver) visitPrintStmt(stmt PrintStatement) Any {
	r.resolveExpr(stmt.Expression)
	return nil
}

func (r *Resolver) visitExprStmt(stmt ExpressionStatement) Any {
	r.resolveExpr(stmt.Expression)
	return nil
}

func (r *Resolver) visitVarStmt(stmt VarStatement) Any {
	r.declare(stmt.Name)
	if stmt.Initializer != nil {
		r.resolveExpr(stmt.Initializer)
	}
	r.define(stmt.Name)
	return nil
}

func (r *Resolver) visitBlockStmt(stmt BlockStatement) Any {
	r.beginScope()
	for _, statement := range stmt.Statements {
		r.resolveStmt(statement)
	}
	r.endScope()
	return nil
}

func (r *Resolver) visitIfStmt(stmt IfStatement) Any {
	r.resolveExpr(stmt.Condition)
	r.resolveStmt(stmt.ThenBlock)
	if stmt.ElseBlock != nil {
		r.resolveStmt(stmt.ElseBlock)
	}
	return nil
}

func (r *Resolver) visitWhileStmt(stmt WhileStatement) Any {
	r.resolveExpr(stmt.Condition)
	r.resolveStmt(stmt.Body)
	return nil
}

func (r *Resolver) visitFunctionStmt(stmt FunctionStatement) Any {
	r.declare(stmt.Name)
	r.define(stmt.Name)

	r.resolveFunction(stmt)
	return nil
}

func (r *Resolver) visitReturnStmt(stmt ReturnStatement) Any {
	if stmt.Value != nil {
		r.resolveExpr(stmt.Value)
	}
	return nil
}

func (r *Resolver) resolveStmt(stmt Statement) Any {
	stmt.Accept(r)
	return nil
}

func (r *Resolver) resolveLocal(expr Expression, name Token) Any {
	for i := r.scopes.Len() - 1; i >= 0; i-- {
		if _, ok := r.scopes.Get(i)[name.Lexeme]; ok {
			r.interpreter.Resolve(expr, r.scopes.Len()-1-i)
			return nil
		}
	}
	return nil
}

func (r *Resolver) resolveFunction(function FunctionStatement) Any {
	r.beginScope()
	for _, param := range function.Params {
		r.declare(param)
		r.define(param)
	}
	for _, statement := range function.Body {
		r.resolveStmt(statement)
	}
	r.endScope()
	return nil
}

func (r *Resolver) beginScope() {
	r.scopes.Push(make(map[string]bool))
}

func (r *Resolver) endScope() {
	r.scopes.Pop()
}

func (r *Resolver) declare(name Token) {
	if r.scopes.IsEmpty() {
		return
	}
	scope := r.scopes.Peek()
	if _, ok := scope[name.Lexeme]; ok {
		parseFault(name, "Already a variable with this name in this scope.")
	}
	scope[name.Lexeme] = false
}

func (r *Resolver) define(name Token) {
	if r.scopes.IsEmpty() {
		return
	}
	scope := r.scopes.Peek()
	scope[name.Lexeme] = true
}
