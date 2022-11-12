package main

import (
	"fmt"
)

func main() {
	launch := true
	//launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launch)
	
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)
}
