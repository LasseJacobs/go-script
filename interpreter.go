package main

import "fmt"

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

type Interpreter struct {
}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Interpret(expr Expression) {
	defer func() {
		if err := recover(); err != nil {
			runtimeFault(err.(RuntimeError))
		}
	}()
	var value Any = i.evaluate(expr)
	fmt.Println(i.stringify(value))
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
