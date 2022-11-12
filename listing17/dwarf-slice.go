package main

import "fmt"

func main() {
	dwarf := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("dwarf %T\ndwarfArray %T", dwarf, dwarfArray)
}

