package main

import "fmt"

var nowhere *int

func main() {
	// a := 100
	// nowhere = &a
	fmt.Println(nowhere)

	fmt.Println(*nowhere)
}
