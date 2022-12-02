package main

import (
	"fmt"
)

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
	// count is no longer in scope
}
