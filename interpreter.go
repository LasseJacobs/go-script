package main

import (
	"fmt"
)

/*
class RuntimeError extends RuntimeException {
  final Token token;

  RuntimeError(Token token, String message) {
    super(message);
    this.token = token;
  }
}
*/

type RuntimeError struct {
	token   Token
	message string
}

func NewRuntimeError(token Token, message string) RuntimeError {
	return RuntimeError{token: token, message: message}
}

type Interpreter struct {
	globals *Environment
	env     *Environment
	locals  map[Expression]int
}

func NewInterpreter() *Interpreter {
	globals := NewEnvironment()
	globals.define("clock", clockFn{})
	return &Interpreter{globals: globals, env: globals, locals: make(map[Expression]int)}
}

func (i *Interpreter) Interpret(statements []Statement) {
	defer func() {
		if err := recover(); err != nil {
			runtimeFault(err.(RuntimeError))
		}
	}()
	for _, s := range statements {
		i.execute(s)
	}
}

func (i *Interpreter) execute(statement Statement) Any {
	return statement.Accept(i)
}

func (i *Interpreter) Resolve(expr Expression, depth int) Any {
	i.locals[expr] = depth
	return nil
}

func (i *Interpreter) executeBlock(statements []Statement, environment *Environment) Any {
	previous := i.env
	defer func() {
		i.env = previous
	}()
	i.env = environment
	for _, stmt := range statements {
		ret := i.execute(stmt)
		if ret != nil {
			return ret
		}
	}
	return nil
}

func (i *Interpreter) visitBinaryExpr(expr BinaryExpression) Any {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.TokenType {
	case TT_GREATER:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) > right.(float64)
	case TT_GREATER_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) >= right.(float64)
	case TT_LESS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) < right.(float64)
	case TT_LESS_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) <= right.(float64)
	case TT_MINUS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) - right.(float64)
	case TT_SLASH:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) / right.(float64)
	case TT_STAR:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) * right.(float64)
	case TT_PLUS:
		vf1, ok1 := left.(float64)
		vf2, ok2 := right.(float64)
		if ok1 && ok2 {
			return vf1 + vf2
		}
		vs1, ok1 := left.(string)
		vs2, ok2 := right.(string)
		if ok1 && ok2 {
			return vs1 + vs2
		}
		panic(RuntimeError{token: expr.Operator, message: "Operands must be a numbers or strings."})
	case TT_BANG_EQUAL:
		return !i.isEqual(left, right)
	case TT_EQUAL_EQUAL:
		return i.isEqual(left, right)
	}
	// unreachable
	return nil
}

func (i *Interpreter) visitGroupingExpr(expr GroupingExpression) Any {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) visitLiteralExpr(expr LiteralExpression) Any {
	return expr.Value
}

func (i *Interpreter) visitUnaryExpr(expr UnaryExpression) Any {
	right := i.evaluate(expr.Right)
	switch expr.Operator.TokenType {
	case TT_MINUS:
		i.checkNumberOperand(expr.Operator, right)
		return -(right.(float64))
	case TT_BANG:
		return !i.isTruthy(expr)
	}
	// should be unreachable
	return nil
}

func (i *Interpreter) visitVarExpr(expr VariableExpression) Any {
	return i.env.get(expr.Name)
}

func (i *Interpreter) visitAssignExpr(expr AssignExpression) Any {
	value := i.evaluate(expr.Value)
	//old way: i.env.assign(expr.Name, value)
	distance, ok := i.locals[expr]
	if ok {
		i.env.assignAt(distance, expr.Name, value)
	} else {
		i.globals.assign(expr.Name, value)
	}

	return value
}

func (i *Interpreter) visitCallExpr(expr CallExpression) Any {
	callee := i.evaluate(expr.Callee)
	var arguments []Any
	for _, arg := range expr.Arguments {
		arguments = append(arguments, i.evaluate(arg))
	}
	function, ok := callee.(Callable)
	if !ok {
		panic(NewRuntimeError(expr.Paren, "Can only call functions."))
	}
	if len(arguments) != function.Arity() {
		panic(NewRuntimeError(expr.Paren, fmt.Sprintf("Expected %d arguments but got %d.", function.Arity(), len(arguments))))
	}
	return function.Call(i, arguments)
}

/*
	Statement interface
*/
func (i *Interpreter) visitExprStmt(stmt ExpressionStatement) Any {
	i.evaluate(stmt.Expression)
	return nil
}

func (i *Interpreter) visitPrintStmt(stmt PrintStatement) Any {
	value := i.evaluate(stmt.Expression)
	fmt.Printf("%s\n", i.stringify(value))
	return nil
}

func (i *Interpreter) visitVarStmt(stmt VarStatement) Any {
	var value Any = nil
	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}
	i.env.define(stmt.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) visitBlockStmt(stmt BlockStatement) Any {
	return i.executeBlock(stmt.Statements, NewEnvironmentWithEnclosing(i.env))
}

func (i *Interpreter) visitIfStmt(stmt IfStatement) Any {
	if i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.ThenBlock)
	} else if stmt.ElseBlock != nil {
		i.execute(stmt.ElseBlock)
	}
	return nil
}

func (i *Interpreter) visitWhileStmt(stmt WhileStatement) Any {
	for i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.Body)
	}
	return nil
}

func (i *Interpreter) visitFunctionStmt(stmt FunctionStatement) Any {
	function := NewFunction(stmt, i.env)
	i.env.define(stmt.Name.Lexeme, function)
	return nil
}

func (i *Interpreter) visitReturnStmt(stmt ReturnStatement) Any {
	var value Any = nil
	if stmt.Value != nil {
		value = i.evaluate(stmt.Value)
	}
	return value
}

/*
	Helpers
*/
func (i *Interpreter) evaluate(expr Expression) Any {
	return expr.Accept(i)
}

func (i *Interpreter) isTruthy(obj Any) bool {
	if obj == nil {
		return false
	}
	if b, ok := obj.(bool); ok {
		return b
	}
	return true
}

func (i *Interpreter) isEqual(a Any, b Any) bool {
	// todo: the likely requires some work still to properly deal with pointers
	return a == b
}

func (i *Interpreter) stringify(object Any) string {
	if object == nil {
		return "nil"
	}
	if f, ok := object.(float64); ok {
		return fmt.Sprintf("%f", f)
	}
	return fmt.Sprintf("%s", object)
}

func (i *Interpreter) checkNumberOperand(operator Token, operand Any) {
	if _, ok := operand.(float64); ok {
		return
	}
	panic(RuntimeError{token: operator, message: "Operand must be a number."})
}

func (i *Interpreter) checkNumberOperands(operator Token, left Any, right Any) {
	_, leftOk := left.(float64)
	_, rightOk := right.(float64)
	if leftOk && rightOk {
		return
	}
	panic(RuntimeError{token: operator, message: "Operands must be a numbers."})
}

func (i *Interpreter) lookUpVariable(name Token, expr Expression) Any {
	distance, ok := i.locals[expr]
	if ok {
		return i.env.getAt(distance, name.Lexeme)
	} else {
		return i.globals.get(name)
	}
}
