package main

import (
	"fmt"
)

func main() {
	var red uint8 = 64

	red++
	fmt.Printf("%08b\n", red)
}
