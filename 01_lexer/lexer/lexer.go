package lexer

import (
	"lexer/token"
)

type Lexer struct {
	input			string
	currentPosition	int
	nextPosition	int
	currentChar			byte
}

/** 
Constructor which returns instantiated lexer given the input.
*/
func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readNextChar()
	return lexer
}

func (lexer *Lexer) ParseToken() token.Token {
	var newToken token.Token

	lexer.skipWhitespace()

	switch lexer.currentChar {
	case '=':
		if (lexer.peekChar() == '=') {
			char := lexer.currentChar
			lexer.readNextChar()
			newToken = token.Token{Type: token.EQUAL, Literal: string(char) + string(lexer.currentChar)}
		} else {
			newToken = createNewToken(token.ASSIGN, lexer.currentChar)
		}
	case ';':
		newToken = createNewToken(token.SEMICOLON, lexer.currentChar)
	case '(':
		newToken = createNewToken(token.LPAREN, lexer.currentChar)
	case ')':
		newToken = createNewToken(token.RPAREN, lexer.currentChar)
	case ',':
		newToken = createNewToken(token.COMMA, lexer.currentChar)
	case '+':
		newToken = createNewToken(token.PLUS, lexer.currentChar)
	case '{':
		newToken = createNewToken(token.LBRACE, lexer.currentChar)
	case '}':
		newToken = createNewToken(token.RBRACE, lexer.currentChar)
	case '-':
		newToken = createNewToken(token.MINUS, lexer.currentChar)
	case '!':
		if (lexer.peekChar() == '=') {
			char := lexer.currentChar
			lexer.readNextChar()
			newToken = token.Token{Type: token.NOT_EQUAL, Literal: string(char) + string(lexer.currentChar)}
		} else {
			newToken = createNewToken(token.BANG, lexer.currentChar)
		}
	case '*':
		newToken = createNewToken(token.ASTERISK, lexer.currentChar)
	case '/':
		newToken = createNewToken(token.SLASH, lexer.currentChar)
	case '<':
		newToken = createNewToken(token.LT, lexer.currentChar)
	case '>':
		newToken = createNewToken(token.GT, lexer.currentChar)
	case 0:
		newToken.Literal = ""
		newToken.Type = token.EOF
	default:
		if isCharLetter(lexer.currentChar) {
			newToken.Literal = lexer.readIdentifier()
			newToken.Type = token.GetTokenTypeFromgetTokenTypeFromLiteralLiteral(newToken.Literal)
			return newToken
		} else if isCharDigit(lexer.currentChar) {
			newToken.Type = token.INT
			newToken.Literal = lexer.readNumber()
			return newToken
		} else {
			newToken = createNewToken(token.ILLEGAL, lexer.currentChar)
		}
	}

	lexer.readNextChar()
	return newToken
}

func (lexer *Lexer) peekChar() byte {
	if lexer.nextPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.nextPosition]
	}
}

/*
Function which reads the next character in the input.

We first check if we have reached the end of the input, and if we did so we set the char to 0. We use this value to terminate the reading.

Otherwise we just move to the next character in the sequence and update the state of the lexer using ch, position and readPosition values.
*/
func (lexer *Lexer) readNextChar() {
	if lexer.nextPosition >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.input[lexer.nextPosition]
	}
	lexer.currentPosition = lexer.nextPosition
	lexer.nextPosition++
}

/** 
In our language whitespace basically is just used to improve readability, which means we can just skip it.

This function iterates to the next character if the whitespace character is encountered.
*/
func (lexer *Lexer) skipWhitespace() {
	for lexer.currentChar == ' ' || lexer.currentChar == '\t' || lexer.currentChar == '\n' || lexer.currentChar == '\r' {
		lexer.readNextChar()
	}
}

func isCharLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isCharDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

/**
readidentified and readNumber functions:

Remember the beggining of the identified in the lexer input and iterate until you find the end of it.

Return the whole identifier.
*/
func (lexer *Lexer) readIdentifier() string {
	position := lexer.currentPosition
	for isCharLetter(lexer.currentChar) {
		lexer.readNextChar()
	}
	return lexer.input[position:lexer.currentPosition]
}
func (l *Lexer) readNumber() string {
	currentPosition := l.currentPosition
	for isCharDigit(l.currentChar) {
		l.readNextChar()
	}

	return l.input[currentPosition:l.currentPosition]
}

/**
Helper function which returns the token representation of the data we have.
*/
func createNewToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}