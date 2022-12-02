package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	rows    = 9
	columns = 9
)

type SudokuError []error

var (
	NotInBoundErr = errors.New("Index out of range.")
	NotEmptyErr   = errors.New("Cell is not empty.")
	RowErr        = errors.New("Incorrect row.")
	ColErr        = errors.New("Incorrect column.")
	RegionErr     = errors.New("Incorrect region.")
)

func (se SudokuError) Error() string {
	var s []string
	for idx, e := range se {
		s = append(s, fmt.Sprintf("%d - %v", idx, e))
	}
	return strings.Join(s, "\n")
}

type Grid [rows][columns]int8

func (g *Grid) Set(row, column int, digit int8) error {
	var err SudokuError

	if !g.IsInBound(row, column) {
		err = append(err, NotInBoundErr)
	}

	if !g.IsCellEmpty(row, column) {
		err = append(err, NotEmptyErr)
	}

	if !g.IsRowCorrect(row, digit) {
		err = append(err, RowErr)
	}

	if !g.IsColCorrect(column, digit) {
		err = append(err, ColErr)
	}

	if !g.IsRegionCorrect(row, column, digit) {
		err = append(err, RegionErr)
	}

	if len(err) > 0 {
		return err
	}

	g[row][column] = digit

	fmt.Printf("Set %d at (%d, %d).\n", digit, row, column)

	return nil
}

func (g Grid) IsInBound(row, column int) bool {
	return row >= 0 && row < rows && column >= 0 && column < columns	
}

func (g Grid) IsCellEmpty(row, column int) bool {
	return g[row][column] == 0
}

func (g Grid) IsRowCorrect(row int, digit int8) bool {
	for _, cell := range g[row] {
		if cell == digit {
			return false
		}
	}
	return true
}

func (g Grid) IsColCorrect(col int, digit int8) bool {
	for _, row := range g {
		if row[col] == digit {
			return false
		}
	}
	return true
}

func (g Grid) IsRegionCorrect(row, col int, digit int8) bool {
	h, v := rows/3, columns/3
	r, c := (row+h)%h, (col+v)%v
	var (
		irows, icolumns []int
	)

	switch r {
	case 0:
		irows = append(irows, row+1, row+2)
	case 1:
		irows = append(irows, row-1, row+1)
	case 2:
		irows = append(irows, row-2, row-1)
	}

	switch c {
	case 0:
		icolumns = append(icolumns, col+1, col+2)
	case 1:
		icolumns = append(icolumns, col-1, col+1)
	case 2:
		icolumns = append(icolumns, col-1, col-2)
	}

	for _, r := range irows {
		for _, c := range icolumns {
			if g[r][c] == digit {
				return false
			}
		}
	}
	return true
}

func (g *Grid) clear (row, column int) bool {
	g.Set(row, column, 0)
}

func (g Grid) String() string {
	var s string
	for _, r := range g {
		for _, c := range r {
			if c == 0 {
				s += fmt.Sprintf("%3s", "*")
			} else {
				s += fmt.Sprintf("%3d", c)
			}
		}
		s += "\n"
	}
	return s
}

func NewSudoku(s [rows][columns]int8) *Grid {
	g := Grid(s)
	return &g
}

func main() {
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	fmt.Println(*s)

	if e := s.Set(0, 8, 3); e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	fmt.Println(*s)
}
