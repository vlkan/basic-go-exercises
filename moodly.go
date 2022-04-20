package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[your name]")
		return
	}
	name := os.Args[1]

	moodly := [...]string{" happy", " awesome",
		" cool", " sad", " bad",
		" terrible",
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moodly))

	fmt.Printf("%s feels%s\n", name, moodly[n])
}
