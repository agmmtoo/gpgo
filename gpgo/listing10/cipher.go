package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"

	message := strings.ToUpper(strings.Replace(plainText, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	for i, c := range(message) {
		k := rune(keyword[i % len(keyword)])
		if c >= 'A' && c <= 'Z' {
			c -= 'A'
			k -= 'A'

			c = (c + k) % 26 + 'A'
			fmt.Printf("%c", c)
		}
	}

}

