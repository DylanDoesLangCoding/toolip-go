package lexer

import (
	"fmt"
	"testing"

	"toolip-go/token"
)

func TestNextToken(t *testing.T) {
	input := `string start = "Welcome to "
string langName = "Toolip"
func string concat(string a, string b)
	return a .. b
end
string message = concat(start, langName)`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.STRTYPE, "string"},
		{token.IDENTIFIER, "start"},
		{token.ASSIGN, "="},
		{token.STRVAL, "\"Welcome to \""},
		{token.NEWLINE, "\n"},
		{token.STRTYPE, "string"},
		{token.IDENTIFIER, "langName"},
		{token.ASSIGN, "="},
		{token.STRVAL, "\"Toolip\""},
		{token.NEWLINE, "\n"},
		{token.FUNC, "func"},
		{token.STRTYPE, "string"},
		{token.IDENTIFIER, "concat"},
		{token.LPAREN, "("},
		{token.STRTYPE, "string"},
		{token.IDENTIFIER, "a"},
		{token.COMMA, ","},
		{token.STRTYPE, "string"},
		{token.IDENTIFIER, "b"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "a"},
		{token.CONCAT, ".."},
		{token.IDENTIFIER, "b"},
		{token.NEWLINE, "\n"},
		{token.END, "end"},
		{token.NEWLINE, "\n"},
		{token.STRTYPE, "string"},
		{token.IDENTIFIER, "message"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "concat"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "start"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "langName"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	var lineNo int = 1
	for i, tt := range tests {
		i++
		tok := l.NextToken()
		if tok.Type == token.NEWLINE {
			lineNo++
		}
		if tok.Type != tt.expectedType {
			t.Fatalf("lexer/toollexer_test.go:%d. expected=(%q: %q), got=(%q: %q).", lineNo, tt.expectedType, tt.expectedValue, tok.Type, tok.Value)
		}
		fmt.Printf("Token( %q, %q)\n", tok.Type, tok.Value)
	}
}
