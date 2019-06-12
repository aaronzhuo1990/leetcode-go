package main

import "fmt"

func main() {
	cans := []int{2, 3, 6, 7}
	target := 7

	result := combinationSum(cans, target)

	fmt.Printf("result %#v\n", result)
}

func combinationSum(candidates []int, target int) *[][]int {
	result := make([][]int, 0)
	solution := make([]int, 0)

	calcCombinationSum(candidates, solution, target, &result)

	return &result
}

func calcCombinationSum(candidates []int, solution []int, target int, result *[][]int) {
	// It means the sum of the elements in `solution` equals to given target.
	if target == 0 {
		*result = append(*result, solution)
		fmt.Printf("result %#v\n", result)
	}

	if len(candidates) == 0 || target < candidates[0] {
		// It means the given solution is not a qualified solution, return.
		return
	}

	newSolution := append(solution, candidates[0])

	// Put the first element of candidates into the new solution and keep calculating
	calcCombinationSum(candidates, newSolution, target-candidates[0], result)

	// Keeping finding the solution from the rest.
	calcCombinationSum(candidates[1:], solution, target, result)
}
