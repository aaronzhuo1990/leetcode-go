package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Printf("result: %t\n", wordSearch(board, "ABCCED")) // Should be true
	fmt.Printf("result: %t\n", wordSearch(board, "ABCB"))   // Should be false

}

//
func wordSearch(board [][]byte, word string) bool {
	rows := len(board)
	if rows == 0 {
		return false
	}

	cols := len(board[0])
	if cols == 0 {
		return false
	}

	if len(word) == 0 {
		return false
	}

	var dfs func(int, int, int) bool
	dfs = func(row, col, wordIndex int) bool {
		if len(word) == wordIndex {
			return true // The word is found in the last round.
		}

		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[wordIndex] {
			return false
		}

		temp := board[row][col]
		board[row][col] = '0'            // To avoid that the same character is used twice.
		wordIndex++                      // Move to the next character
		if dfs(row-1, col, wordIndex) || // Move left
			dfs(row, col-1, wordIndex) || // Up
			dfs(row+1, col, wordIndex) || // Right
			dfs(row, col+1, wordIndex) { // Down
			return true
		}

		board[row][col] = temp
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
