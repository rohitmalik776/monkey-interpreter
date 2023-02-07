package main

import (
	"fmt"
	"monkey-lang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!, this is Monkey REPL\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
