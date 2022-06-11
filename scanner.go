package main

type Scanner struct {
	source []rune
	tokens []Token

	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	//TODO: rune conversion should be done at advance time. I see no reason to store a rune array
	return &Scanner{source: []rune(source), start: 0, current: 0, line: 1}
}

func (s *Scanner) ScanTokens() {
	for !s.isAtEnd() {

	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() rune {
	var c = s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) scanToken() {
	var c rune = s.advance()
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
		fault(s.line, "Unexpected character")
	}

}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		//null character
		return '\000'
	}
	return s.source[s.current]
}

func (s *Scanner) match(r rune) bool {
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
	var text []rune = s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{
		TokenType: tokenType,
		Lexeme:    string(text),
		Literal:   literal,
		Line:      s.line,
	})
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
	s.addToken(TT_STRING, string(value))
}
