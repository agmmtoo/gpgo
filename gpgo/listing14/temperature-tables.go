package main

import (
	"fmt"
)

type convFunc func(t float64) float64

func main() {
	drawTable(cToF, true)
	drawTable(fToC, false)
}

func drawTable(convFunc convFunc, cToF bool) {
	var start, end, step = -40, 100, 5

	printRow("=", 23)

	if (cToF) {
		fmt.Println("| ºC | ºF |")
	} else {
		fmt.Println("| ºF | ºC |")
	}

	printRow("=", 23)

	for i := start; i < end; i += step {
		from := float64(i)
		to := convFunc(from)
		fmt.Printf("|%4.1f|%4.1f|\n", from, to)
	}

	printRow("=", 23)
}

func cToF(c float64) float64 {
	return c * 9.0/5.0 + 32.0
}

func fToC(f float64) float64 {
	return (f - 32.0) * 5.0/9.0
}

func printRow(char string, count int) {
	// fmt.Println(char * count)
	for i := 0; i < count; i++ {
		fmt.Print(char)
	}
	fmt.Print("\n")
}
