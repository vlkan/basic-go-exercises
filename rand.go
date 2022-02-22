package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guess := 10

	rand.Seed(time.Now().UnixNano())

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
