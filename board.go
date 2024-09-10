package main

import (
	"fmt"
)

func main() {
	board := make([][]bool, 6)
	for row := range board {
		board[row] = make([]bool, 5)
		for col := range board[row] {
			fmt.Println(board[row][col])
		}
	}
	fmt.Println(board)
}

