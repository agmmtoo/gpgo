package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const nickels = 5
	const dimes = 10
	const quarters = 25

	var deposit string
	var piggy int16

	for piggy < 20 * 100 {
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
		dollars := piggy / 100
		cents := piggy % 100
		fmt.Printf("%s $%d.%02d\n", deposit, dollars, cents)
	}

//	fmt.Printf("finally %v\n", piggy/100)
}
