package main

import "fmt"

func main() {
	var board [9][9]int

	for column := range board[0] {
		board[0][column] = column
	}

	fmt.Println(board)
}
