package main

import "fmt"

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func main() {
	var kelvin float64 = 0
	f := kelvinToFahrenheit(kelvin)
	fmt.Print(kelvin, "ºK is", f, "ºF")
}

// celcisusToFahrenheit converts ºC to ºF
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

// kelvinToFahrenheit converts ºK to ºF
func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	f := celsiusToFahrenheit(c)
	return f
}
