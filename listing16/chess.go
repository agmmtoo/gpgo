package main

import "fmt"

func main() {
	// chess board
	var board [8][8]rune

	// Kings
	board[0][4] = 'k'
	board[7][4] = 'K'

	// Queens
	board[0][3] = 'q'
	board[7][3] = 'Q'

	// Rooks
	board[0][0], board[0][7] = 'r', 'r'
	board[7][0], board[7][7] = 'R', 'R'

	// Bishops
	board[0][2], board[0][5] = 'b', 'b'
	board[7][2], board[7][5] = 'B', 'B'

	// Knights
	board[0][1], board[0][6] = 'k', 'k'
	board[7][1], board[7][6] = 'K', 'K'

	// Pawns
	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}

	for _, row := range board {
		for _, p := range row {
			fmt.Printf("%4c", p)
		}
		fmt.Println("")
	}
}
