package main

import (
	"fmt"
	"math"
)

type deg float64

func (d deg) rad() float64 {
	return float64(d) * math.Pi / 180
}

type location struct {
	lat, long deg
	name      string
}

func (l location) description() string {
	return fmt.Sprintf("%v (%.1fº, %.1fº)", l.name, l.lat, l.long)
}

type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(p1.lat.rad())
	s2, c2 := math.Sincos(p2.lat.rad())
	clong := math.Cos((p1.long - p2.long).rad())

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

type gps struct {
	current, destination location
	world                world
}

func (g gps) distance() float64 {
	d := g.world.distance(g.current, g.destination)
	if d < 0 {
		d = -d
	}
	return d
}

func (g gps) message() string {
	return fmt.Sprintf("%f km to %v!", g.distance(), g.destination.description())
}

type rover struct {
	gps
}

func main() {
	mars := world{3389.5}
	bradbury := location{
		deg(-4.5895),
		deg(137.4417),
		"Bradbury Landing",
	}
	elysium := location{
		deg(4.5),
		deg(135.9),
		"Elysium Planitia",
	}
	marsGPS := gps{
		bradbury,
		elysium,
		mars,
	}

	curiosity := rover{
		marsGPS,
	}

	fmt.Println(curiosity.message())
}
