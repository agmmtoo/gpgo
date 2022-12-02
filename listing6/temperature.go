package main

import "fmt"

func main() {
	celsius := 21.0
	fmt.Print((celsius * 9.0 / 5.0) + 32.0, "° F\n")
	fmt.Print((9 * celsius / 5.0) + 32.0, "° F\n")
}
