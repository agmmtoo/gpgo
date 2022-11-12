package main

import (
	"fmt"
)

type Planets []string

func (planets Planets) terraform(){
	terraform(planets)
}

func terraform(planets []string) {
	for i, name := range planets {
		planets[i] = "New " + name
	}
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()

	fmt.Println(planets)
}
