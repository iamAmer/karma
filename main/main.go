package main

import (
	"fmt"
	"os"
	"os/user"
	"karma/repl"
	"karma/logo"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", logo.KARMA)
	fmt.Printf("Hello %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}