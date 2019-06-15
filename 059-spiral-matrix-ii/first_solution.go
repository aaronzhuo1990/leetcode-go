package main

import "fmt"

func main() {

	matrix := spiralMatrix(4)

	fmt.Printf("Spiral matrix\n")
	for _, row := range matrix {
		fmt.Printf("\t%v\n", row)
	}
}

const (
	right = iota + 1
	down
	left
	up
)

func spiralMatrix(num int) [][]int {
	rows, cols := num, num
	matrix := make([][]int, rows)

	for i := 0; i < rows; i++ {
		row := make([]int, cols)
		matrix[i] = row
	}

	index := 1

	row, col := 0, -1
	status := right
	for rows != 0 && cols != 0 {
		switch status {
		case right:
			for cnt := cols; cnt > 0; cnt-- {
				col++
				matrix[row][col] = index
				index++
			}
			rows--
			status = down
		case down:
			for cnt := rows; cnt > 0; cnt-- {
				row++
				matrix[row][col] = index
				index++
			}
			cols--
			status = left
		case left:
			for cnt := cols; cnt > 0; cnt-- {
				col--
				matrix[row][col] = index
				index++
			}
			rows--
			status = up
		case up:
			for cnt := rows; cnt > 0; cnt-- {
				row--
				matrix[row][col] = index
				index++
			}
			cols--
			status = right
		}
	}

	return matrix
}
