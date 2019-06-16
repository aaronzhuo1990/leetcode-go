package main

import "fmt"

func main() {

	input := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	result := minimumPathSum(input)
	fmt.Printf("result %d\n", result) // Should be 7
}

func minimumPathSum(input [][]int) int {

	if len(input) == 0 {
		return 0
	}

	rows, cols := len(input), len(input[0])

	for i := 1; i < cols; i++ {
		input[0][i] += input[0][i-1]
	}

	for j := 1; j < rows; j++ {
		input[j][0] += input[0][j-1]
	}

	for i := 1; i < cols; i++ {
		for j := 1; j < rows; j++ {
			input[j][i] += min(input[j-1][i], input[j][i-1]) // left + top
		}
	}

	return input[rows-1][cols-1]
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}
