package main

import (
	"fmt"
	"strings"
)

func hyperspace(worlds []string) {
	for i, world := range worlds {
		worlds[i] = strings.TrimSpace(world)
	}
}

func main() {
	planets := []string{" Venus	", "Earch  ", " Mars"}
	fmt.Println(planets)

	hyperspace(planets)
	fmt.Println(planets)
	fmt.Println(strings.Join(planets, ""))
}
