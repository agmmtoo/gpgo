package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var k kelvin = 294
	c := kelvinToCelsius(k)
	fmt.Println(k, "ºK is ", c, "ºC")

	const SUNLIT_MOON_TEMPT celsius = 127 // ºC
	fmt.Printf("The surface temperature of the sunlit moon is %vºK", celsiusToKelvin(SUNLIT_MOON_TEMPT))
}
