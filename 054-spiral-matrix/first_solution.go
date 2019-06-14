package main

import "fmt"

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	result := spiralMatrix(matrix)

	fmt.Printf("Spiral matrix %v\n", result)
}

const (
	right = iota + 1
	down
	left
	up
)

func spiralMatrix(matrix [][]int) []int {

	if len(matrix) == 0 {
		return nil
	}

	rows, cols := len(matrix), len(matrix[0])
	result := make([]int, rows*cols)
	index := 0

	row, col := 0, -1
	status := right
	for rows != 0 && cols != 0 {
		switch status {
		case right:
			for cnt := cols; cnt > 0; cnt-- {
				col++
				result[index] = matrix[row][col]
				index++
			}
			rows--
			status = down
		case down:
			for cnt := rows; cnt > 0; cnt-- {
				row++
				result[index] = matrix[row][col]
				index++
			}
			cols--
			status = left
		case left:
			for cnt := cols; cnt > 0; cnt-- {
				col--
				result[index] = matrix[row][col]
				index++
			}
			rows--
			status = up
		case up:
			for cnt := rows; cnt > 0; cnt-- {
				row--
				result[index] = matrix[row][col]
				index++
			}
			cols--
			status = right
		}
	}

	return result
}
