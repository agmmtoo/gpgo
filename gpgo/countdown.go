package main

import (
	"fmt"
	"time"
)

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
		time.Sleep(time.Second)
	}
	fmt.Println("Liftoff!")
}
