package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]

	feet, err := strconv.ParseFloat(args, 64)
	if err != nil {
		fmt.Printf("ERROR: %q is not a number.\n", args)
		return
	}
	meters := feet * 0.3048
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
