package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	fmt.Printf("%v %[1]T\n", planets)

	sort.StringSlice(planets).Sort()

	fmt.Printf("%v %[1]T\n", planets)
}
