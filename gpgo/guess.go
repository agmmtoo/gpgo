package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const luckyNum = 7 
	var i = 0
	for i < 10 {
		if rand.Intn(100) == luckyNum {
			fmt.Printf("Correct %10v\n", luckyNum)
			break
		}
		i++
	}

	if i == 10 {
		fmt.Println("Good luck next time!")
	}
}
