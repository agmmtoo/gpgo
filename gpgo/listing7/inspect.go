package main

import (
	"fmt"
)

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)

	fmt.Printf("%T %T %T %T\n", "lmao", 1, 0xd4,  true)

	fmt.Printf("%x \n", 0xFF)

	var r, g, b uint8 = 0, 50, 213

	fmt.Printf("color: #%02x%02x%02x;", r, g, b)

	fmt.Printf("\n%05.3f\t%[1]T\n", 13/2)
}
