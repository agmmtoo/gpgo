package main

import "fmt"

func main() {
	elements := make([]string, 10)
	s := "why the fuck os keep crashing?"

	for i := 0; i < len(s); i++ {
		elements = append(elements, string(s[i]))
		fmt.Printf("%v, cap: %v\n", elements, cap(elements))
	}
}
