package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "မင်္ဂလာပါ"

	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size)
}
