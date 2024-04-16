package repl

import (
	"bufio"
	"fmt"
	"io"
	"lexer/lexer"
	"lexer/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "quit" {
			break
		}

		newLexer := lexer.NewLexer(line)

		for currentToken := newLexer.ParseToken(); currentToken.Type != token.EOF; currentToken = newLexer.ParseToken() {
			fmt.Printf("%+v\n", currentToken)
		}
		}
}