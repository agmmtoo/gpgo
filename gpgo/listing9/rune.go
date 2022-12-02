package main

import (
	"fmt"
)

func main() {
	var pi, alpha, omega rune = 960, 940, 969
	var bang byte = 33

	fmt.Printf("%v %[1]c: %[1]T, %v %[2]c: %[2]T, %v %[3]c: %[3]T, %v %[4]c: %[4]T\n", pi, alpha, omega, bang)
}
