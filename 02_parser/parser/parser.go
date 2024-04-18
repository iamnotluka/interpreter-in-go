package parser

import (
	"lexer/lexer"
	"lexer/token"
	"parser/ast"
)

type Parser struct {
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.ParseToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}