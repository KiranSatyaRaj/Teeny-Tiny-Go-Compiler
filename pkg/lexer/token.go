package lexer

import (
	"fmt"
)

type Token[T any] struct{
	text T
	kind TokenKind
}

type TokenKind int

type tokenMetadata struct {
	name string
}

var tokenMap = map[TokenKind]string{
	0: "EOF", 1: "NEWLINE", 2: "NUMBER", 3: "IDENT", 4: "STRING",
	// Keywords
	5: "LABEL", 6: "GOTO", 7: "PRINT", 8: "INPUT", 9: "LET", 10: "IF", 11: "THEN", 12: "ENDIF", 13: "WHILE", 14: "REPEAT", 15: "ENDWHILE",
	// Operators
	16: "EQ", 17: "PLUS", 18: "MINUS", 19: "ASTERISK", 20: "SLASH", 21: "EQEQ", 22: "NOTEQ", 23: "LT", 24: "LTEQ", 25: "GT", 26: "GTEQ",
}

const (
	EOF TokenKind = iota
	NEWLINE
	NUMBER
	IDENT
	STRING
	// Keywords
	LABEL 
	GOTO
	PRINT
	INPUT
	LET
	IF
	THEN
	ENDIF
	WHILE
	REPEAT
	ENDWHILE
	// Operators
	EQ
	PLUS
	MINUS
	ASTERISK
	SLASH
	EQEQ
	NOTEQ
	LT
	LTEQ
	GT
	GTEQ
)

func (t TokenKind) String() string {
	return fmt.Sprintf("%v", tokenMap[t])
}

func (token *Token[any]) GetKind() (TokenKind) {
	return token.kind
}

func (token *Token[any]) checkIfKeyword(tokText string) TokenKind {
	for k, v := range tokenMap {
		if v == tokText && k >= 5 && k < 16 {
			return k
		}
	}
	return TokenKind(-1)
}
