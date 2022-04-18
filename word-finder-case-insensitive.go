package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query_temp := os.Args[1:]

	var query [10]string

	for j := 0; j < len(query_temp); j++ {
		query[j] = strings.ToLower(query_temp[j])
	}

	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}
}
