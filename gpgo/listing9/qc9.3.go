package main

import "fmt"

func main() {
	s := "shalom"

	for _,c := range(s) {
		fmt.Printf("%c %[1]v\n", c)
	}
}
