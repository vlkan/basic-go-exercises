package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	// var books [yearly]string // when create empty array
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	// "Kafka's Revenge"
	// "Stay Golden"
	// "Everythingship"
	//  books[0] + " 2nd Edition"

	// fmt.Printf("books	: %T\n", books)
	// fmt.Println("books	:", books)
	// fmt.Printf("books	: %q\n", books)  //q or s meant string arrays
	fmt.Printf("books	: %#v\n", books) //types and elements together

	// var (
	// 	wBooks [winter]string
	// 	sBooks [summer]string
	// )
	// wBooks[0] = books[0]
	// //sBooks[0] = books[1]
	// //sBooks[0] = books[2]
	// //sBooks[0] = books[3]

	// for i := range sBooks {
	// 	sBooks[i] = books[i+1]
	// }
	// //range copy the array, when change the data, its only change the copy version not original one

	// fmt.Printf("\nwinter : %#v\n", wBooks)
	// fmt.Printf("\nwinter : %#v\n", sBooks)

	// var published [len(books)]bool //len also working constant

	// published[0] = true
	// published[len(books)-1] = true

	// fmt.Println("\nPubliched Books:")
	// for i, ok := range published {
	// 	if ok {
	// 		fmt.Printf("+ %s\n", books[i])
	// 	}
	// }

}
