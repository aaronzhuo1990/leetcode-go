package main

import "fmt"

func main() {

	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(matrix)

	// Expected output:
	// [0,0,0,0],
	// [0,4,5,0],
	// [0,3,1,0]

	fmt.Printf("result\n")
	for _, row := range matrix {
		fmt.Printf("\t%v\n", row)
	}
}

func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	first_col := -1

	for row := 0; row < rows; row++ {
		if first_col != 0 && matrix[row][0] == 0 {
			first_col = 0
		}
		for col := 1; col < cols; col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 1; col-- {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
		if first_col == 0 {
			matrix[row][0] = 0
		}
	}
}
