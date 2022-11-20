package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	lat, long float64
}

func nnewLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	// Bradbury Landing: 4º35'22.2" S, 137º26'30.1" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())

	curiosity := nnewLocation(lat, long)
	fmt.Println(curiosity)
}
