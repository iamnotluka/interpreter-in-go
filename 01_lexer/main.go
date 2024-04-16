package main

import (
	"fmt"
	"lexer/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Luka's programming language called fluxo. Welcome!\n", user.Username)
	fmt.Println("Feel free to start typing in commands.")

	repl.Start(os.Stdin, os.Stdout)
}