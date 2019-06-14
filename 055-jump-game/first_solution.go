package main

import "fmt"

func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	result1 := jumpGame(nums1)
	// Should be true
	fmt.Printf("result1: %t\n", result1)

	nums2 := []int{3, 2, 1, 0, 4}
	result2 := jumpGame(nums2)
	// Should be false
	fmt.Printf("result2: %t\n", result2)
}

func jumpGame(nums []int) bool {

	// start from the second last element
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}

		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > i-j {
				break
			}
		}

		if j == -1 {
			return false
		}

	}

	return true
}
