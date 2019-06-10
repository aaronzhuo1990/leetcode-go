package next_permutation

import "fmt"

func nextPermutation(nums []int) {
	right := len(nums) - 1
	left := right - 1

	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}

	// Reverse the order from left to right
	reverse(nums, left+1, right)

	if left == -1 {
		return
	}

	// Swap the value
	swapIndex := search(nums, left+1, nums[left])
	nums[left], nums[swapIndex] = nums[swapIndex], nums[left]
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func search(nums []int, left, target int) int {
	right := len(nums) - 1

	for left+1 < right {
		fmt.Printf("left %d, right %d\n", left, right)
		mid := (left + right) / 2

		if target < nums[mid] {
			right = mid
		} else {
			left = mid
		}

	}

	// The last `right` points to the first element which vale is greater than target.
	return right
}
