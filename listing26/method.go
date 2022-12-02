package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func (p *person) birthday() {
	p.age++
}

func main() {
	p := person{"ammo", 22}
	p.birthday()
	fmt.Println(p)
}
