// lexer/lexer_test.go

package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType 	token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, 		"="},
		{token.PLUS, 		"+"},
		{token.LPARENT, 	"("},
		{token.RPARENT, 	")"},
		{token.LBRACE, 		"{"},
		{token.RBRACE,		"}"},
		{token.COMMA,		","},
		{token.SEMICOLON,	";"},
		{token.EOF,			""},
	}
}

l := New(input)

for i, tt := range tests {
	tok := l.NextToken()

	if tok.Type != tt.expectedType {
		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			i, tt.expectedType, tok.Type)
	}

	if tok.Literal != tt.expectedLiteral {
		t.Fatalf("tests[%d] - leteral wrong. expected=%q, got=%q",
			i, tt.expectedLiteral, tok.Literal)
	}
}