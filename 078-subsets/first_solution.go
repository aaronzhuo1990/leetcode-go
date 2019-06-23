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

func subsets(nums []int) [][]int {
	res := make([][]int, 0, resSize(nums))
	findSubSets(nums, []int{}, &res)
	return res
}

func resSize(nums []int) int {
	return int(math.Pow(2, float64(len(nums))))
}

func findSubSets(input, tmp []int, result *[][]int) {
	num := len(input)

	if num == 0 {
		*result = append(*result, tmp)
		return
	}

	// take the last element
	last, input := input[num-1], input[:num-1]

	// Find all the subsets with tmp and without last
	findSubSets(input, tmp, result)

	withLast := make([]int, 0, len(tmp)+1)
	withLast = append(withLast, tmp...)
	withLast = append(withLast, last)
	findSubSets(input, withLast, result)
}
