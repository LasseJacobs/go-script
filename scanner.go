package main

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	source string
	tokens []Token

	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source, start: 0, current: 0, line: 1}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.addToken(TT_EOF, nil)
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	var c = s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) scanToken() {
	var c byte = s.advance()
	switch c {
	case '(':
		s.addToken(TT_LEFT_PAREN, nil)
	case ')':
		s.addToken(TT_RIGHT_PAREN, nil)
	case '{':
		s.addToken(TT_LEFT_BRACE, nil)
	case '}':
		s.addToken(TT_RIGHT_BRACE, nil)
	case ',':
		s.addToken(TT_COMMA, nil)
	case '-':
		s.addToken(TT_MINUS, nil)
	case '+':
		s.addToken(TT_PLUS, nil)
	case ';':
		s.addToken(TT_SEMICOLON, nil)
	case '*':
		s.addToken(TT_STAR, nil)
	case '!':
		if s.match('=') {
			s.addToken(TT_BANG_EQUAL, nil)
		} else {
			s.addToken(TT_BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(TT_EQUAL_EQUAL, nil)
		} else {
			s.addToken(TT_EQUAL, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(TT_LESS_EQUAL, nil)
		} else {
			s.addToken(TT_LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(TT_GREATER_EQUAL, nil)
		} else {
			s.addToken(TT_GREATER, nil)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(TT_SLASH, nil)
		}
	case ' ':
		//explicit ignore
	case '\r':
		//explicit ignore
	case '\t':
		//explicit ignore
	case '\n':
		s.line++
	case '"':
		s.scanString()
	default:
		if s.isDigit(c) {
			s.scanNumber()
		} else if s.isAlpha(c) {
			s.scanIdentifier()
		} else {
			fault(s.line, "Unexpected character")
		}
	}
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		//null character
		return '\000'
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		//null character
		return '\000'
	}
	return s.source[s.current+1]
}

func (s *Scanner) match(r byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != r {
		return false
	}
	_ = s.advance()
	return true
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	var text string = s.source[s.start:s.current]
	if tokenType == TT_EOF {
		text = ""
	}
	s.tokens = append(s.tokens, Token{
		TokenType: tokenType,
		Lexeme:    string(text),
		Literal:   literal,
		Line:      s.line,
	})
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func (s *Scanner) isAlphaNumeric(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) scanString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		_ = s.advance()
	}
	if s.isAtEnd() {
		fault(s.line, "Unterminated string.")
	}
	//closing "
	_ = s.advance()
	var value = s.source[s.start+1 : s.current-1]
	s.addToken(TT_STRING, value)
}

func (s *Scanner) scanNumber() {
	for s.isDigit(s.peek()) {
		_ = s.advance()
	}
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		//consume '.'
		_ = s.advance()
		//parse fractional
		for s.isDigit(s.peek()) {
			_ = s.advance()
		}
	}
	var value, err = strconv.ParseFloat(s.source[s.start:s.current], 32)
	if err != nil {
		//this should never happen
		panic(fmt.Errorf("failed to scan number: %s", err))
	}
	s.addToken(TT_NUMBER, value)
}

func (s *Scanner) scanIdentifier() {
	// by having the scan path isAlpha but the loop isAlphaNumeric,
	// identifiers are restricted to starting with an alpha character
	for s.isAlphaNumeric(s.peek()) {
		_ = s.advance()
	}

	var text = s.source[s.start:s.current]
	var t_type = keywords[text]
	if t_type == TT_NO_TOKEN {
		t_type = TT_IDENTIFIER
	}
	s.addToken(t_type, nil)
}
