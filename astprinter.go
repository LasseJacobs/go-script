package main

/*
func main() {
	var expr Expression = BinaryExpression{
		Left: UnaryExpression{
			Operator: Token{TT_MINUS, "-", nil, 1},
			Right:    LiteralExpression{Value: 123},
		},
		Operator: Token{
			TokenType: TT_STAR,
			Lexeme:    "*",
			Literal:   nil,
			Line:      1,
		},
		Right: GroupingExpression{Expression: &LiteralExpression{Value: 54.56}},
	}
	var printer = AstPrinter{}
	fmt.Println(printer.Print(expr))
}
*/

/*
type AstPrinter struct {
}

func (a *AstPrinter) Print(expr Expression) string {
	return a.stringify(expr.Accept(a))
}

func (a *AstPrinter) visitBinaryExpr(expr BinaryExpression) Any {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (a *AstPrinter) visitGroupingExpr(expr GroupingExpression) Any {
	return a.parenthesize("group", expr.Expression)
}

func (a *AstPrinter) visitLiteralExpr(expr LiteralExpression) Any {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%s", expr.Value)
}

func (a *AstPrinter) visitUnaryExpr(expr UnaryExpression) Any {
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (a *AstPrinter) parenthesize(name string, exprs ...Expression) string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(name)
	for _, e := range exprs {
		sb.WriteString(" ")
		sb.WriteString(a.stringify(e.Accept(a)))
	}
	sb.WriteString(")")
	return sb.String()
}

func (a *AstPrinter) stringify(object Any) string {
	return fmt.Sprintf("%s", object)
}
*/
