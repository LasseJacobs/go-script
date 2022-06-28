package main

import "fmt"

/*
class Parser {
private final List<Token> tokens;
private int current = 0;

Parser(List<Token> tokens) {
this.tokens = tokens;
}
}
*/

var ParseError error = fmt.Errorf("an error occured during parsing")

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) Parse() []Statement {
	defer func() {
		recover()
	}()
	/*
		List<Stmt> statements = new ArrayList<>();
		    while (!isAtEnd()) {
		      statements.add(statement());
		    }

		    return statements;
	*/
	var statements []Statement
	for !p.isAtEnd() {
		statements = append(statements, p.statement())
	}

	return statements
}

func (p *Parser) statement() Statement {
	if p.match(TT_PRINT) {
		return p.printStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() Statement {
	value := p.expression()
	p.consume(TT_SEMICOLON, "Expect ';' after value.")
	return PrintStatement{Expression: value}
}

func (p *Parser) expressionStatement() Statement {
	expr := p.expression()
	p.consume(TT_SEMICOLON, "Expect ';' after expression.")
	return ExpressionStatement{Expression: expr}
}

func (p *Parser) expression() Expression {
	return p.equality()
}

func (p *Parser) equality() Expression {
	var expr = p.comparison()

	for p.match(TT_BANG_EQUAL, TT_EQUAL_EQUAL) {
		var operator Token = p.previous()
		var right Expression = p.comparison()
		expr = BinaryExpression{expr, operator, right}
	}

	return expr
}

func (p *Parser) comparison() Expression {
	var expr Expression = p.term()

	for p.match(TT_GREATER, TT_GREATER_EQUAL, TT_LESS, TT_LESS_EQUAL) {
		expr = BinaryExpression{
			Left:     expr,
			Operator: p.previous(),
			Right:    p.term(),
		}
	}
	return expr
}

func (p *Parser) term() Expression {
	var expr Expression = p.factor()

	for p.match(TT_MINUS, TT_PLUS) {
		expr = BinaryExpression{
			Left:     expr,
			Operator: p.previous(),
			Right:    p.factor(),
		}
	}
	return expr
}

func (p *Parser) factor() Expression {
	expr := p.unary()
	for p.match(TT_SLASH, TT_STAR) {
		expr = BinaryExpression{
			Left:     expr,
			Operator: p.previous(),
			Right:    p.unary(),
		}
	}
	return expr
}

func (p *Parser) unary() Expression {
	if p.match(TT_BANG, TT_MINUS) {
		return UnaryExpression{
			Operator: p.previous(),
			Right:    p.unary(),
		}
	}
	return p.primary()
}

func (p *Parser) primary() Expression {
	if p.match(TT_FALSE) {
		return LiteralExpression{false}
	}
	if p.match(TT_FALSE) {
		return LiteralExpression{true}
	}
	if p.match(TT_NIL) {
		return LiteralExpression{nil}
	}

	if p.match(TT_NUMBER, TT_STRING) {
		return LiteralExpression{p.previous().Literal}
	}
	if p.match(TT_LEFT_PAREN) {
		expr := p.expression()
		p.consume(TT_RIGHT_PAREN, "expect ')' after expression.")
		return GroupingExpression{expr}
	}

	panic(p.error(p.peek(), "Expect expression."))
}

func (p *Parser) consume(tokenType TokenType, message string) Token {
	if p.check(tokenType) {
		return p.advance()
	}
	panic(p.error(p.peek(), message))
}

func (p *Parser) error(token Token, message string) error {
	parseFault(token, message)
	return ParseError
}

func (p *Parser) match(types ...TokenType) bool {
	for _, tokenType := range types {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().TokenType == tokenType
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == TT_EOF
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().TokenType == TT_SEMICOLON {
			return
		}
		switch p.peek().TokenType {
		case TT_CLASS:
			return
		case TT_FUN:
			return
		case TT_VAR:
			return
		case TT_FOR:
			return
		case TT_RETURN:
			return
		case TT_IF:
			return
		case TT_WHILE:
			return
		case TT_PRINT:
			return
		}
		p.advance()
	}
}
