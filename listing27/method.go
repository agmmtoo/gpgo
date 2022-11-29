package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	//p.age++
}

func main() {
	var p *person = &person{}
	p.birthday()
	fmt.Println(p)
}
