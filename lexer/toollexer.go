package lexer

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"toolip-go/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	lineNum      int
	linePosition int
	prevToken    token.Token
	ch           byte
	prevCh       byte
}

var numArgs int = len(os.Args)

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.linePosition = -1
	l.lineNum = 1
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	l.prevCh = l.ch
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.linePosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	if l.prevToken.Type == token.NEWLINE {
		l.lineNum++
		l.linePosition = -1
	}

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQUALS, Value: value}
			break
		}
		tok = newToken(token.ASSIGN, l.ch)
		break
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		break
	case '(':
		tok = newToken(token.LPAREN, l.ch)
		break
	case ')':
		tok = newToken(token.RPAREN, l.ch)
		break
	case '{':
		tok = newToken(token.LBRACE, l.ch)
		break
	case '}':
		tok = newToken(token.RBRACE, l.ch)
		break
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
		break
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
		break
	case '#':
		if l.peekChar() == '#' {
			l.eatLineComment()
			break
		} else if l.peekChar() == '[' {
			l.eatBlockComment()
			break
		}
		tok = newToken(token.HASH, l.ch)
		break
	case '+':
		if l.peekChar() == '+' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.INCREMENT, Value: value}
			break
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.PLUSEQ, Value: value}
			break
		}
		tok = newToken(token.PLUS, l.ch)
		break
	case '-':
		if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.DECREMENT, Value: value}
			break
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.MINUSEQ, Value: value}
			break
		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.ARROW, Value: value}
			break
		}
		tok = newToken(token.MINUS, l.ch)
		break
	case '*':
		if l.peekChar() == '*' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			if l.peekChar() == '=' {
				l.readChar()
				value += string(l.ch)
				tok = token.Token{Type: token.EXPOEQ, Value: value}
				break
			}
			tok = token.Token{Type: token.EXPO, Value: value}
			break
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.MULTEQ, Value: value}
			break
		}
		tok = newToken(token.MULT, l.ch)
		break
	case '/':
		if l.peekChar() == '/' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			if l.peekChar() == '=' {
				l.readChar()
				value += string(l.ch)
				tok = token.Token{Type: token.FDIVEQ, Value: value}
				break
			}
			tok = token.Token{Type: token.FDIV, Value: value}
			break
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.DIVEQ, Value: value}
			break
		}
		tok = newToken(token.DIV, l.ch)
	case '%':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.MODEQ, Value: value}
			break
		}
		tok = newToken(token.MOD, l.ch)
		break
	case '<':
		if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			if l.peekChar() == '=' {
				l.readChar()
				value += string(l.ch)
				tok = token.Token{Type: token.LSHIFTEQ, Value: value}
				break
			}
			tok = token.Token{Type: token.LSHIFT, Value: value}
			break
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LTEQ, Value: value}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			if l.peekChar() == '=' {
				l.readChar()
				value += string(l.ch)
				tok = token.Token{Type: token.RSHIFTEQ, Value: value}
				break
			}
			tok = token.Token{Type: token.RSHIFT, Value: value}
			break
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GTEQ, Value: value}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '?':
		if l.peekChar() == '?' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			if l.peekChar() == '=' {
				l.readChar()
				value += string(l.ch)
				tok = token.Token{Type: token.COALESCEEQ, Value: value}
				break
			}
			tok = token.Token{Type: token.COALESCE, Value: value}
			break
		} else {
			tok = newToken(token.TERNARY, l.ch)
			break
		}
	case '&':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.BITANDEQ, Value: value}
			break
		}
		tok = newToken(token.BITAND, l.ch)
		break
	case '|':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.BITOREQ, Value: value}
			break
		}
		tok = newToken(token.BITOR, l.ch)
		break
	case '~':
		tok = newToken(token.BITNOT, l.ch)
		break
	case '^':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.BITXOREQ, Value: value}
			break
		}
		tok = newToken(token.BITXOR, l.ch)
		break
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			tok = token.Token{Type: token.BOOLNOTEQ, Value: value}
			break
		}
		tok = newToken(token.BOOLNOT, l.ch)
		break
	case ',':
		tok = newToken(token.COMMA, l.ch)
		break
	case ':':
		tok = newToken(token.COLON, l.ch)
		break
	case '"':
		tok.Type = token.STRVAL
		tok.Value = l.readSLString()
		break
	case '`':
		tok.Type = token.STRVAL
		tok.Value = l.readMLString()
		break
	case '\'':
		tok.Type = token.CHARVAL
		tok.Value = l.readCharString()
	case '.':
		if l.peekChar() == '.' {
			ch := l.ch
			l.readChar()
			value := string(ch) + string(l.ch)
			if l.peekChar() == '=' {
				l.readChar()
				value += string(l.ch)
				tok = token.Token{Type: token.CONCATEQ, Value: value}
				break
			}
			tok = token.Token{Type: token.CONCAT, Value: value}
			break
		} else {
			tok = newToken(token.DOT, l.ch)
			break
		}
	case '\n':
		tok = newToken(token.NEWLINE, l.ch)
		break
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
		break
	default:
		if isLetter(l.ch) {
			tok.Value = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Value)
			return tok
		} else if isDigit(l.ch) {
			tok = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}
	l.prevToken = tok
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

func (l *Lexer) readSLString() string {
	b := &strings.Builder{}
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == 0 {
			fmt.Printf("Toolip:%d:%d. End of file reached before single-line string was closed.\n", l.lineNum, l.linePosition)
			if numArgs > 1 {
				os.Exit(1)
			}
			break
		}
		if l.ch == '\\' {
			l.readEscapeSequence(b)
			// Skip over the '\\' and the matched single escape char
			l.readChar()
			continue
		}
		if l.ch == '\n' {
			fmt.Printf("Toolip:%d:%d.: Multiple lines in a single-line string.\n", l.lineNum, l.linePosition)
			if numArgs > 1 {
				os.Exit(1)
			}
		}
		if l.ch == '"' {
			break
		}
	}
	return l.input[position:l.position]
}
func (l *Lexer) readMLString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == 0 {
			fmt.Printf("Toolip:%d:%d.: End of file reached before multi-line string was closed.\n", l.lineNum, l.linePosition)
			if numArgs > 1 {
				os.Exit(1)
			}
			break
		}
		if l.ch == '`' {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readCharString() string {
	b := &strings.Builder{}
	position := l.position + 1
	charCount := 0
	for {
		l.readChar()
		if l.ch == 0 {
			fmt.Printf("Toolip:%d:%d.: End of file reached before character string was closed.\n", l.lineNum, l.linePosition)
			if numArgs > 1 {
				os.Exit(1)
			}
			break
		}
		if l.ch == '\n' {
			fmt.Printf("Toolip:%d:%d.: Multiple lines in a character string.\n", l.lineNum, l.linePosition)
			if numArgs > 1 {
				os.Exit(1)
			}
		}
		if l.ch == '\'' {
			if charCount > 1 {
				fmt.Printf("Toolip:%d:%d.: More than 1 byte stored in a character string.\n", l.lineNum, l.linePosition)
				if numArgs > 1 {
					os.Exit(1)
				}
			}
			break
		}
		charCount++
		if l.ch == '\\' {
			l.readEscapeSequence(b)
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) eatLineComment() {
	for {
		l.readChar()
		if l.ch == 0 {
			fmt.Printf("Toolip:%d:%d.: End of file reached before string was closed.\n", l.lineNum, l.linePosition)
			if numArgs > 1 {
				os.Exit(1)
			}
			break
		}
		if l.ch == '\n' {
			break
		}
	}
}

func (l *Lexer) eatBlockComment() {
	for {
		l.readChar()
		if l.ch == 0 {
			fmt.Printf("Toolip:%d:%d.: End of file reached before block comment was closed.\n", l.lineNum, l.linePosition)
			if numArgs > 1 {
				os.Exit(1)
			}
			break
		}
		if l.ch == ']' && l.peekChar() == '#' {
			l.readChar()
			break
		}
	}
}

func (l *Lexer) readEscapeSequence(b *strings.Builder) {
	switch l.peekChar() {
	case '"':
		b.WriteByte('"')
	case '\'':
		b.WriteByte('\'')
	case 'n':
		b.WriteByte('\n')
	case 'r':
		b.WriteByte('\r')
	case 't':
		b.WriteByte('\t')
	case '\\':
		b.WriteByte('\\')
	case 'x':
		l.readChar()
		l.readChar()
		l.readChar()
		src := string([]byte{l.prevCh, l.ch})
		dst, err := hex.DecodeString(src)
		if err != nil {
			fmt.Printf("Toolip:%d:%d. Could not properly decode Hex Escape Sequence! :c\n", l.lineNum, l.linePosition)
			if numArgs > 1 {
				os.Exit(1)
			}
		}
		b.Write(dst)
	}
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

func (l *Lexer) readNumber() token.Token {
	position := l.position
	var numType token.TokenType
	l.readChar()
	for isDigit(l.ch) {
		l.readChar()
	}
	if l.ch == '.' {
		l.readChar()
		numsPastFloat := 0
		for isDigit(l.ch) {
			l.readChar()
			numsPastFloat++
		}
		if numsPastFloat == 0 {
			fmt.Printf("Toolip:%d:%d.: Attempted to pass float without digits following the floating point.\n", l.lineNum, l.linePosition)
			if numArgs > 1 {
				os.Exit(1)
			}
		}
		numType = token.FLOATVAL
	} else {
		numType = token.INTVAL
	}
	return token.Token{Type: numType, Value: l.input[position:l.position]}
}
