package main

import (
	"encoding/json"
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) Decimal() float64 {
	sign := 1.0

	switch c.h {
	case 'S', 's', 'W', 'w':
		sign = -1.0
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%v\" %c", c.d, c.m, c.s, c.h)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	type jsonStruct struct {
		Decimal, Degrees, Minutes, Seconds float64
		Dms, Hemisphere                    string
	}

	return json.Marshal(jsonStruct{
		Decimal:    c.Decimal(),
		Dms:        c.String(),
		Degrees:    c.d,
		Minutes:    c.m,
		Seconds:    c.s,
		Hemisphere: string(c.h),
	})
}

type location struct {
	Lat, Long coordinate
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.Lat, l.Long)
}

func main() {
	elysium := location{
		coordinate{4.0, 30.0, 0.0, 'N'},
		coordinate{135, 54, 0, 'E'},
	}
	fmt.Println("Elysium Planitia is at", elysium)

	r, _ := json.Marshal(elysium)
	fmt.Println(string(r))
}
