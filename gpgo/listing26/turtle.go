package main

import (
	"fmt"
)

type turtle struct {
	x, y int
}

func (t *turtle) move(axis rune, step int) {
	switch axis {
	case 'x', 'X':
		t.x += step
	case 'y', 'Y':
		t.y += step
	}
}

func (t *turtle) up() {
	t.move('y', 1)
}
func (t *turtle) right() {
	t.move('x', 1)
}
func (t *turtle) down() {
	t.move('Y', -1)
}
func (t *turtle) left() {
	t.move('X', -1)
}

func main() {
	t := turtle{10, 10}
	t.right()
	t.right()
	t.down()
	t.left()
	t.up()
	fmt.Println(t)
}
