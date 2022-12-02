package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)

	for i := range u {
		u[i] = make([]bool, width)
	}

	return u
}

// Universe.Show
func (u Universe) Show() {
	fmt.Print(" \033[2J")
	for _, r := range u {
		for _, c := range r {
			cell := " "
			if c {
				cell = "*"
			}
			fmt.Print(cell)
		}
		fmt.Print("\n")
	}
}

// Universe.Seed
func (u Universe) Seed() {
	for _, r := range u {
		for idx := range r {
			i := rand.Intn(100)
			if i < 25 {
				r[idx] = true
			}
		}
	}
}

// Universe.Alive
func (u Universe) Alive(x, y int) bool {
	// wrap width
	if x < 0 {
		x += width
	} else if x >= width {
		x %= width
	}

	// wrap height
	if y < 0 {
		y += height
	} else if y >= height {
		y %= height
	}
	return u[y][x]
}

// Universe.Neighbors
func (u Universe) Neighbors(x, y int) int {
	// cell := u[y][x]
	n := 0
	cord := [][]int{
		[]int{-1, -1}, []int{-1, 0}, []int{-1, 1},
		[]int{0, -1}, []int{0, 1},
		[]int{1, -1}, []int{1, 0}, []int{1, 1},
	}
	for _, c := range cord {
		nx, ny := x + c[0], y + c[1]
		if (u.Alive(nx, ny)) {
			n++
		}
	}

	return n
}

// Universe.Next
func (u Universe) Next(x, y int) bool {
	if u.Alive(x, y) {
		if u.Neighbors(x, y) >= 2 || u.Neighbors(x, y) <= 3 {
			return true
		} else {
			return false
		}
	} else if u.Neighbors(x, y) == 3 {
		return true
	} else {
		return false
	}
}

// Universe.Step
func Step(a, b Universe) {
	for ri := range a {
		for ci := range a[ri] {
			b[ri][ci] = a.Next(ri, ci)
		}
	}
}


func main() {
	u := NewUniverse()
	u.Seed()
	u.Show()
	//c := u(80, 15)
	fmt.Print(u.Next(80, 15))
	fmt.Print(u.Neighbors(80, 15))
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	for i := 0; i < 20; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second)
		a,b = b, a
	}

}
