package lexer

import (
	"fmt"
	"os"
	"unicode"
)

var Trex Lexer

func Init(source string) {
	Trex = Lexer{}
	Trex.init(source)
}

type Lexer struct {
	src string
	curChar rune
	curPos int
}

// Initialize Lexer
func (l *Lexer) init(src string) {
	l.src = src + "\n" // Source code to lex as a string. Append a newline to simplify lexing/parsing the last token/statement
	l.curChar = 0 // Current character in the string.
	l.curPos = -1 // Current position in the string.
	l.nextChar()
}

// Process the next character.
func (l *Lexer) nextChar() {
	l.curPos += 1
	if l.curPos >= len(l.src) {
		l.curChar = '\000'
	} else {
		l.curChar = rune(l.src[l.curPos])
	}
}

// Return the lookahead character.
func (l *Lexer) peek() rune {
	if l.curPos+1 >= len(l.src) {
		return '\000'
	}
	return rune(l.src[l.curPos+1])
}

// Invalid token found, print error message and exit.
func (l * Lexer) abort(msg string) {
	fmt.Println("Lexing error:", msg)
	os.Exit(1)
}

// skip whitespace except newlines, which we will use to indicate the end of a statement.
func (l *Lexer) skipWhitespace() {
	for l.curChar == ' ' || l.curChar == '\t' || l.curChar == '\r' {
		l.nextChar()
	}
}

// Skip comments in code.
func (l *Lexer) skipComments() {
	for l.curChar != '\n' {
		l.nextChar()
	}
}

// Return the next token.
func (l *Lexer) GetToken() (Token[any], error){
	l.skipWhitespace()
	var token Token[any]
	// Check the first character of this token to see if we can decide what it is.
	// If it is a multiple character operator (eg., !=), number, identifier, or keyword then we will process the rest.
	switch l.curChar {
	case '+': // Plus token
		token = Token[any]{l.curChar, PLUS}
	case '-': // Minus token
		token = Token[any]{l.curChar, MINUS}
	case '*': // Asterisk token
		token = Token[any]{l.curChar, ASTERISK}
	case '/': // Slash token
		if l.peek() == '/' {
			l.skipComments()
			token = Token[any]{l.curChar, NEWLINE}
		} else {
			token = Token[any]{l.curChar, SLASH}
		}
	case '\n': // newline token
		token = Token[any]{l.curChar, NEWLINE}
	case '\000': // EOF
		token = Token[any]{l.curChar, EOF}
	case '=':
		// Check whether this token is = or ==
		if l.peek() == '=' {
			lastChar := l.curChar
			l.nextChar()
			token = Token[any]{lastChar + l.curChar, EQEQ}
		} else {
			token = Token[any]{l.curChar, EQ}
		}
	case '>':
		// Check whether this token is > or >=
		if l.peek() == '=' {
			lastChar := l.curChar
			l.nextChar()
			token = Token[any]{lastChar + l.curChar, GTEQ}
		} else {
			token = Token[any]{l.curChar, GT}
		}
	case '<':
		// Check whether this token is < or <=
		if l.peek() == '=' {
			lastChar := l.curChar
			l.nextChar()
			token = Token[any]{lastChar + l.curChar, LTEQ}
		} else {
			token = Token[any]{l.curChar, LT}
		}
	case '!':
		// Check whether this token is ! or !=
		if l.peek() == '=' {
			lastChar := l.curChar
			l.nextChar()
			token = Token[any]{lastChar + l.curChar, NOTEQ}
		} else {
			l.abort("Expected !=, got !" + string(l.peek()))
		}
	case '"':
		// Get character between quotations.
		l.nextChar()
		startPos := l.curPos
		for l.curChar != '"' {
			// Don't allow special characters in the string. No escape characters, newlines, tabs, or %.
			// We will be using C's printf on this string.
			if l.curChar == '\r' || l.curChar == '\n' || l.curChar == '\t' || l.curChar == '\\' || l.curChar == '%' {
				l.abort("Illegal character in string.")
			}
			l.nextChar()
		}
		tokText := l.src[startPos:l.curPos] // Get the substring
		token = Token[any]{tokText, STRING}
	default:
		if unicode.IsDigit(l.curChar) {
			// Leading character is a digit, so this must be a number.
			// Get all consecutive digits and decimal if there is one.
			startPos := l.curPos
			for unicode.IsDigit(l.peek()) {
				l.nextChar()
			}
			if l.peek() == '.' { // Decimal
				l.nextChar()

				// Must have atleast one digit after decimal.
				if !unicode.IsDigit(l.peek()) {
					// Error!
					l.abort("Illegal character in number.")
				}
				for unicode.IsDigit(l.peek()) {
					l.nextChar()
				}
			}
			tokText := l.src[startPos:l.curPos+1]
			token = Token[any]{tokText, NUMBER}
		} else if unicode.IsLetter(l.curChar) {
			// Leading character is a letter, so this must be an identifier or a keyword.
			// Get all consecutive alpha numeric characters.
			startPos := l.curPos
			for IsAlphaNum(l.peek()) {
				l.nextChar()
			}

			// Check if token is in the list of keywords.
			tokText := l.src[startPos:l.curPos+1]
			keywordKind := token.checkIfKeyword(tokText)
			if keywordKind == -1 {
				token = Token[any]{tokText, IDENT}
			} else {
				token = Token[any]{tokText, keywordKind}
			}
		} else {
			// unknown token
			l.abort("Unknown Token")
		}
	}

	l.nextChar()
	return token, nil
}

func IsAlphaNum(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}

func GetCurChar() (rune) {
	return Trex.curChar
}
