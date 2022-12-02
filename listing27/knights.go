package main

import (
	"fmt"
)

type item string

type character struct {
	name      string
	leftHand  *item
	rightHand *item
}

func (c character) String() string {
	if c.leftHand == nil || c.rightHand == nil {
		return fmt.Sprintf("%v with bare hand", c.name)
	}
	return fmt.Sprintf("%v (%v, %v)", c.name, c.leftHand, c.rightHand)
}

func (c character) pickup(i *item) {
	c.rightHand = i
	fmt.Println(c, "picked up an", i)
}

func (c *character) give(to *character) {
	to.leftHand, c.rightHand = c.rightHand, to.leftHand
	fmt.Println(c, "gave to", to)
}

func main() {
	arthurItem := item("Arthur's Sword")
	knightItem := item("Knight's Sword")

	arthur := &character{name: "Arthur", leftHand: &arthurItem}
	knight := &character{name: "Knight"}

	fmt.Println(arthur, knight)
	arthur.pickup(&knightItem)
	arthur.give(knight)

	fmt.Println(arthur, knight)
}
