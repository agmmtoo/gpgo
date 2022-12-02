package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

const (
	rows    = 9
	columns = 9
)

type Grid [rows][columns]int8

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError

	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit

	return nil
}

func inBounds(row, column int) bool {
	return row >= 0 && row < rows && column >= 0 && column < columns
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func main() {
	var g Grid
	err := g.Set(1, 9, 19)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("-%v\n", e)
			}
		}
		os.Exit(1)
	}
}
