package main

import "fmt"

func main() {

	result := uniquePaths(7, 3)
	fmt.Printf("result %d\n", result) // should be 28
}

func uniquePaths(cols int, rows int) int {
	path := make([][]int, rows)

	for i := 0; i < rows; i++ {
		tmp := make([]int, cols)
		path[i] = tmp
	}

	for i := 0; i < cols; i++ {
		path[0][i] = 1
	}

	for j := 0; j < rows; j++ {
		path[j][0] = 1
	}

	for i := 1; i < cols; i++ {
		for j := 1; j < rows; j++ {
			path[j][i] = path[j-1][i] + path[j][i-1] // top + left
		}
	}

	return path[rows-1][cols-1]
}
