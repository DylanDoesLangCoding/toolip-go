package main

import (
	"fmt"
	"os"
	"os/user"
	"toolip-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 1 {
		fmt.Printf("Hello %s! Welcome to the Toolip Programming Language!\nEnter some commands below to get started.\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
	}
}
