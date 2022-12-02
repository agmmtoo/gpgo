package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	fmt.Println(planets[:])

	slice1 := planets[:3]
	fmt.Println(slice1)
	planets[2] = "???"
	fmt.Println(planets)
	fmt.Println(slice1)
}
