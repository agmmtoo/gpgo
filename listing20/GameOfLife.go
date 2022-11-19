package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 100
	height = 25
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)

	for r := range u {
		u[r] = make([]bool, width)
	}

	return u
}

func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		randX, randY := rand.Intn(width), rand.Intn(height)
		u.Set(randX, randY, true)
	}
}

func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	n := 0

	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

func (u Universe) String() string {
	var s string

	for _, h := range u {
		for i := range h {
			cell := " "
			if h[i] {
				cell = "*"
			}
			s += cell
		}
		s += "\n"
	}
	return s
}

func (u Universe) Show() {
	fmt.Print("\033[H", u.String())
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 1000; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second/10)
		a, b = b, a
	}
}
