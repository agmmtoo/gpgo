package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

// c.fahrenheit
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((9 / 5 * c) + 32)
}

// c.kelvin
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// f.celsius
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

// f.kelvin
func (f fahrenheit) kelvin() kelvin {
	c := f.celsius()
	return c.kelvin()
}

// k.celsius
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// k.fahrenheit
func (k kelvin) fahrenheit() fahrenheit {
	c := k.celsius()
	return c.fahrenheit()
}
func main() {
	const c celsius = 100 // ºC
	fmt.Println(c, "ºC = ", c.fahrenheit(), "ºF ", c.kelvin(), "ºK")

	const f fahrenheit = 100 // ºF
	fmt.Println(f, "ºF = ", f.celsius(), "ºC ", f.kelvin(), "ºK")

	const k kelvin = 100 // ºK
	fmt.Println(k, "ºK = ", k.celsius(), "ºC ", k.fahrenheit(), "ºF")
}
