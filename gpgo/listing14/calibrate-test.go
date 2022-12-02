package main

import "fmt"
import "math/rand"

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	sensor := calibrate(fakeSensor, 5)
	fmt.Println(sensor(), sensor(), sensor(), sensor())
}
