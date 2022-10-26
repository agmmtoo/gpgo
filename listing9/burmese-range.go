package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "မင်္ဂလာပါ"

	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	for _, c := range question {
		fmt.Printf("%c\n", c)
	}
}
