package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars": -65,
		"Moon": 0,
	}

	if temp, exist := temperature["Moon"]; exist {
		fmt.Printf("On average the Earth is %vºC\n", temp)
	} else {
		fmt.Println("No known record for this.")
	}
}
