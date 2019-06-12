package main

import (
	"fmt"
	"sort"
)

func main() {
	cans := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	result := combinationUniqueSum(cans, target)

	fmt.Printf("result %#v\n", result)
}

func combinationUniqueSum(candidates []int, target int) *[][]int {

	sort.Ints(candidates)

	result := make([][]int, 0)
	solution := make([]int, 0)

	calcCombinationSum(candidates, solution, target, &result)

	return &result
}

func calcCombinationSum(candidates []int, solution []int, target int, result *[][]int) {

	// It means the sum of the elements in `solution` equals to given target.
	if target == 0 {
		*result = append(*result, solution)
	}

	if len(candidates) == 0 || target < candidates[0] {
		// It means the given solution is not a qualified solution, return.
		return
	}

	// Put the first element of candidates into the new solution and keep calculating
	newSolution := append(solution, candidates[0])
	calcCombinationSum(candidates[1:], newSolution, target-candidates[0], result)

	// Keeping finding the solution from the rest.
	calcCombinationSum(next(candidates), solution, target, result)
}

func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}
