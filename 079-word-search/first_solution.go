package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{1, 2, 3}
	result := subsets(input)

	fmt.Printf("all the subsets\n")
	for _, row := range result {
		fmt.Printf("\t%v\n", row)
	}

}

// TODO: User deep first search to realize the solution.
