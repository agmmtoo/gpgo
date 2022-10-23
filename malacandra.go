package main

import "fmt"

func main() {
	const DISTANCE = 56000000 // km
	const DAY = 28
	const HOURINDAY = 24

	var ans = DISTANCE / (HOURINDAY * DAY)
	
	fmt.Printf("The ship needs to travel %v km/h", ans)
}
