package main

import (
	"fmt"
)

func main() {
	var red uint8 = 255

	fmt.Println(red + 10)
	var number int8 = 127
	number += 127
	fmt.Println(number)

	var green int8 = -127
	green -= 1
	fmt.Println(green)
}
