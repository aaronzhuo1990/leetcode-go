package remove_element

import "fmt"

// removeElement removes all instances of that value in-place and return the new length.
func removeElement(nums []int, val int) int {
	var left, right int

	// Find the first element which value is `val`, left is going to save the element which value is not `val`.
	for left = 0; left < len(nums); left++ {
		if nums[left] == val {
			break
		}
	}

	for right = left + 1; right < len(nums); right++ {
		if nums[right] == val {
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
	}

	fmt.Printf("nums %v\n", nums)

	return left
}
