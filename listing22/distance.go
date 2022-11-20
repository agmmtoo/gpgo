package main

import (
	"fmt"
	"math"
)

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

	return sign * (c.d + c.d/60 + c.s/3600)
}

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	spirit := newLocation(
		coordinate{14, 34, 6.2, 'S'},
		coordinate{175, 28, 21.5, 'E'},
	)
	opportunity := newLocation(
		coordinate{1, 56, 46.3, 'S'},
		coordinate{354, 28, 24.2, 'E'},
	)
	curiosity := newLocation(
		coordinate{4, 35, 22.2, 'S'},
		coordinate{137, 26, 30.1, 'E'},
	)
	insight := newLocation(
		coordinate{4, 30, 0.0, 'N'},
		coordinate{135, 54, 0, 'e'},
	)

	fmt.Println(spirit, opportunity, curiosity, insight)

	mars := world{3389.5}

	locations := map[string]location{
		"spirit":      spirit,
		"opportunity": opportunity,
		"curiosity":   curiosity,
		"insight":     insight,
	}

	for name1, loc1 := range locations {
		for name2, loc2 := range locations {
			if !(name1 == name2) {
				dist := mars.distance(loc1, loc2)
				fmt.Printf("distance between %s %s: %f km\n", name1, name2, dist)
			}
		}
	}

	earth := world{6371.0}

	london := newLocation(
		coordinate{
			d: 51,
			m: 30,
			h: 'N',
		},
		coordinate{
			m: 8,
			h: 'W',
		},
	)
	paris := newLocation(
		coordinate{
			d: 48,
			m: 51,
			h: 'N',
		},
		coordinate{
			d: 2,
			m: 21,
			h: 'E',
		},
	)

	fmt.Printf("Landon and Paris is %f km aparts.\n", earth.distance(london, paris))

	mountSharp := newLocation(coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
	olympusMons := newLocation(coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})

	fmt.Printf("Mount Sharp & Olympus Mons: %f km\n", mars.distance(mountSharp, olympusMons))
}
