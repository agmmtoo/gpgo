package main

import "fmt"

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((9 / 5 * c) + 32)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

const (
	line = "======================="
	rowFormat = "|%8s|%8s|\n"
	numberFormat = "%.1f"

	start = -4
	end = 100
	step = 5
)

type getRowFn func(row int) (string, string)

func ctofRow(row int) (string, string) {
	c := celsius(row)
	f := c.fahrenheit()

	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)

	return cell1, cell2
}

func ftocRow(row int) (string, string) {
	f := fahrenheit(row)
	c := f.celsius()

	return fmt.Sprintf(numberFormat, f), fmt.Sprintf(numberFormat, c)
}

func drawTable(hdr1, hdr2 string, start, end, step int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)

	for row := start; row <= end; row += step {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func main() {
	drawTable("ºC", "ºF", start, end, step, ctofRow)
	drawTable("ºF", "ºC", start, end, step, ftocRow)
}
