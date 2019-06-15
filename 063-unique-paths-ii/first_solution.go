package main

import "fmt"

func main() {

	input := [][]int{
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0},
	}

	result := uniquePathsII(input)
	fmt.Printf("result %d\n", result) // Should be 6
}

func uniquePathsII(input [][]int) int {

	if len(input) == 0 {
		return 0
	}

	rows, cols := len(input), len(input[0])

	path := make([][]int, rows)
	for i := 0; i < rows; i++ {
		tmp := make([]int, cols)
		path[i] = tmp
	}

	for i := 0; i < cols; i++ {
		if input[0][i] == 0 {
			// It is not an obstacle
			path[0][i] = 1
		} else {
			// Not more way to go
			break
		}
	}

	for j := 0; j < rows; j++ {
		if input[j][0] == 0 {
			// It is not an obstacle
			path[j][0] = 1
		} else {
			// Not more way to go
			break
		}
	}

	for i := 1; i < cols; i++ {
		for j := 1; j < rows; j++ {
			if input[j][i] == 1 {
				// It is an obstacle
				path[j][i] = 0
			} else {
				path[j][i] = path[j-1][i] + path[j][i-1] // top + left
			}
		}
	}

	return path[rows-1][cols-1]
}
