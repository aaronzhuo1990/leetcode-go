package remove_duplicates_from_sorted_array

// removeDuplicatesFromSortedArray removes the duplicates in-place such that each element appear only once and return the new length.
func removeDuplicatesFromSortedArray(nums []int) int {
	left, right := 0, 1

	for ; right < len(nums); right++ {
		if nums[left] == nums[right] {
			// Keep moving.
			continue
		}

		// Increase left by 1 as we are going to use it to save the next unique element.
		left++
		tmpVal := nums[left]
		nums[left] = nums[right]
		nums[right] = tmpVal

		// The above code can be written as:
		// nums[left], nums[right] = nums[right], nums[left]
		// You can see how confusing it is.

	}

	return left + 1
}
