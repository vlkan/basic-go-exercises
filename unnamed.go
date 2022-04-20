package main

import "fmt"

func main() {

	type (
		bookcase [5]int
		cabinet  [5]int
	)

	blue := bookcase{6, 9, 3, 2, 1}
	red := cabinet{6, 9, 3, 2, 1}

	fmt.Print("Are they equal? ")

	if cabinet(blue) == red {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
