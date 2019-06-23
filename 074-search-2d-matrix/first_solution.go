package main

import "fmt"

func main() {

	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}

	fmt.Printf("result for searching 11 %t\n", searchMatrix(matrix, 11)) // should be true
	fmt.Printf("result for searching 55 %t\n", searchMatrix(matrix, 55)) // should be false
}

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}

	cols := len(matrix[0])
	if cols == 0 {
		return false
	}

	if target < matrix[0][0] || matrix[rows-1][cols-1] < target {
		return false
	}

	// Find the right row
	row := 1
	for ; row < rows; row++ {
		if matrix[row-1][0] <= target && target < matrix[row][0] {
			row--
			break
		}
	}

	i, j := 0, cols-1
	for i <= j {
		mid := (i + j) / 2
		switch {
		case matrix[row][mid] < target:
			j = mid - 1
		case matrix[row][mid] > target:
			i = mid + 1
		default:
			return true
		}
	}

	return false
}
