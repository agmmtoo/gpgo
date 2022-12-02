package main

import (
	"fmt"
)

func main() {
	var s string = "yes"

	var boolean bool

	switch s {
	case "true", "yes", "1":
		boolean = true
	case "false", "no", "0":
		boolean = false
	default:
		fmt.Println("Error!", boolean)
	}
}
