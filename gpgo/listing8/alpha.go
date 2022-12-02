package main

import (
	"fmt"
)

func main() {
	const c = 299792               // km/s
	const sPerD = 1 * 60 * 60 * 24 // s

	var d int64  = 41.3e12 // km
	fmt.Println("Alpha Century is", d, "km away.")

	days := d / c / sPerD
	fmt.Println("Takes", days, "days to travel at light speed.")
}
