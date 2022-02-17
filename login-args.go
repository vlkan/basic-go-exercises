package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	u, p := args[1], args[2]

	if u != "volkan" {
		fmt.Printf("Access denied for %q.\n", u)
	} else if p != "1234" {
		fmt.Printf("wrong password for %q.\n", u)
	} else {
		fmt.Printf("Access successful to %q.\n", u)
	}
}
