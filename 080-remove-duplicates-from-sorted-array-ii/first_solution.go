package main

import "fmt"

func main() {
	input := []int{1, 1, 1, 2, 2, 3}
	result := removeDuplicates(input)

	fmt.Printf("result %d\n", result)

}

func removeDuplicates(input []int) int {
	size := len(input)

	qualified := 2

	for candidate := qualified; candidate < size; candidate++ {
		if input[qualified-2] != input[candidate] {
			input[qualified] = input[candidate]
			qualified++
		}
	}

	return qualified
}
