package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const D = 62100000
	var tripType = "One-way"
	var spaceline string

	fmt.Printf("%-20v %4v %-10v %5v\n", "Spaceline", "Days", "Trip type", "Price")
	for i := 0; i < 42; i++ {
		fmt.Print("=")
	}

	for i := 0; i < 10; i++ {
			km := rand.Intn(30) + 16
			duration := (D / km) / 60 / 60 / 24 // day
			price := rand.Intn(50) + 36 - duration
			if price % 2 == 0 {
				price *= 2
				tripType = "Round-trip"
			} else {
				tripType = "One-way"
			}

			if r := rand.Intn(9); r < 3 {
				spaceline = "Space Adventures"
			} else if r < 6 {
				spaceline = "SpaceX"
			} else {
				spaceline = "Virgin Galactic"
			}
			fmt.Printf("\n%-20v %4v %-10v $%-5v", spaceline, duration, tripType, price)
	}
}
