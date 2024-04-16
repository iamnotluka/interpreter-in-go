package repl

import (
	"bufio"
	"example/lexer/lexer"
	"example/lexer/token"
	"fmt"
	"io"
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

		newLexer := lexer.NewLexer(line)

		for currentToken := newLexer.ParseToken(); currentToken.Type != token.EOF; currentToken = newLexer.ParseToken() {
			fmt.Printf("%+v\n", currentToken)
		}
		}
}