// How long does it take to get to Mars?
package main

import "fmt"

func main() {
	const SPACEXSPEED = 100800 // km/h
	const HOURSPERDAY = 24
	var distance = 96300000 // km

	fmt.Println(distance/SPACEXSPEED/HOURSPERDAY, "days")
}
