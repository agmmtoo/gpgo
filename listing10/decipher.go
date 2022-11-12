package main

import (
	"fmt"
	//"strings"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	for i, c := range cipherText {
		keywordIndex := i % len(keyword)
		keyChar := rune(keyword[keywordIndex])
		//fmt.Printf("%c", c - int32(keyChar))
		c -= 'A'
		k := keyChar - 'A'

		c = (c - k + 26) % 26 + 'A'
		fmt.Print(string(c))
	}

}
