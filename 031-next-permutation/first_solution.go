package next_permutation

func nextPermutation(nums []int) {
	right := len(nums) - 1
	left := right - 1

	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}

	// Keep the element which needs to swap position later.
	swapIndex := left
	swapVal := nums[swapIndex]

	// Reverse the order from left to right
	left++
	reverse(nums, left, right)

	// Swap the value
	// TODO: This can be done by binary search.
	for ; left < right; left++ {
		if swapVal < nums[left] {
			nums[swapIndex], nums[left] = nums[left], nums[swapIndex]
			break
		}
	}
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
