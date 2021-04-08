package lexer

import (
	"testing"

	"bo/token"
)

func TestNextToken(t *testing.T) {
	input := `=;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// {token.LET, "let"},
		{token.ASSIGN, "="},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		t.Log(tok)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
