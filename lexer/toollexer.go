package lexer

import (
	"log"
	"toolip-go/token"
)

var a int = 1

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '"':
		tok.Type = token.STRVAL
		tok.Value = l.readString()
	case '.':
		if l.peekChar() == '.' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.CONCAT, Value: value}
		} else {
			tok = newToken(token.DOT, l.ch)
		}
	case '\n':
		tok = newToken(token.NEWLINE, l.ch)
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Value = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Value)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Value: string(ch)}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == 0 {
			log.Fatalln("ERROR: End of file reached before string was closed")
			break
		}
		if l.ch == '\\' {
			if isEscapeChar(l.peekChar()) {
				l.readChar()
			} else {
				log.Fatalln("ERROR: Unrecognized escape character" + string(l.peekChar()) + ".")
			}
		}
		if l.ch == '"' {
			break
		}
	}
	return l.input[position:l.position]
}

func isEscapeChar(ch byte) bool {
	return ch == 'n' || ch == '\\' || ch == 't' || ch == 'r' || ch == '"'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	l.readChar()
	for isNextIdentChar(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isNextIdentChar(ch byte) bool {
	return isLetter(ch) || isDigit(ch)
}
