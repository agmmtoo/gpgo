package main

import (
	"fmt"
)

func main() {
	age := 41
	marsAge := float64(age)
	marsDay := 687.0
	earthDay := 365.2425
	marsAge = marsAge * earthDay / marsDay
	fmt.Println("I am", marsAge, "years old on Mars.")
}
