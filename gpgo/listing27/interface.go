package main

import "fmt"

func main() {
	var v interface{}
	var p *int
	v = p
	fmt.Printf("%T %[1]v %v %#[1]v", v, v == nil)
}
