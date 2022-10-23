package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const nickels = 0.05
	const dimes = 0.10
	const quarters = 0.25

	var deposit string
	var piggy float32

	for piggy < 20 {
		if n := rand.Intn(3); n < 1 {
			piggy += nickels
			deposit = "nickels"
		} else if n < 2 {
			piggy += dimes
			deposit = "dimes"
		} else {
			piggy += quarters
			deposit = "quarters"
		}
		fmt.Printf("put a %s: %4.2f\n", deposit, piggy)
	}

	fmt.Printf("finally %4.2f\n", piggy)
}
