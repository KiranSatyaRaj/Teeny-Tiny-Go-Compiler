package parser

import (
	"os"
	"github.com/KiranSatyaRaj/TeenyTinyGoCompiler/pkg/lexer"
)

var parser Parser

func Init(lexer lexer.Lexer) {
	parser = &Parser{}
	parser.lexer = lexer
	parser.nextToken()
	parser.nextToken() //Call this twice to initialize current and peek.
}

// Parser object keeps track of current token and checks if the code matches the grammar.
type Parser struct{
	lexer lexer.Lexer
	curToken lexer.Token
	peekToken lexer.Token
}

// Return true if the current token matches.
func (p *Parser) checkToken(kind lexer.Tokenkind) bool {
	return kind == p.curToken.kind.GetKind()
}

// Return true if the next token matches.
func (p *Parser) checkPeek(kind lexer.TokenKind) bool {
	return kind == p.peekToken.kind.GetKind()
}

// Try to match current token. If not, error. Advances the current token.
func (p *Parser) match(kind lexer.TokenKind) {
	if !p.checkToken(kind) {
		p.abort(fmt.Sprintf("Expected %s, got %s", kind, p.curToken.kind.GetKind()))
	}
	p.nextToken()
}

// Advances the current token.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.GetToken()
}

func (p *Parser) abort(msg string) {
	fmt.Println("Parsing error: ", msg)
	os.Exit(1)
}

// Production rules

// program ::= {statement}, it means that a program is made up of zero or more statements
func (p *Parser) program() {
	fmt.Println("PROGRAM")

	// Parse all the statements in the program.
	for p.checkToken(lexer.EOF) {
		p.statement()
	}
}
