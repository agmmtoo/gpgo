package main

import (
	"fmt"
	//"math/big"
)

func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792 // km/s
	const year = 1 * 60 * 60 * 24 * 365

	lightyear := distance / lightSpeed / year

	fmt.Println("It takes", lightyear, "lightyears to get to Canis Major Dwarf");
}
