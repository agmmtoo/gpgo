package main

import "fmt"

type location struct {
	name string
	lat  float64
	long float64
}

func main() {
	locations := []location{
		{"Bradbury Landing", -4.5895, 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}

	fmt.Print(locations)
}
