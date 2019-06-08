package twosum

// twoSum returns indices of the two numbers such that they add up to a specific target.
func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for i, num := range nums {
		// num is b in this case
		if j, ok := index[target-num]; ok {
			// if a = target - b
			// a found as long as ok is true, return the indexes of a and b.
			return []int{j, i}
		}

		index[num] = i
	}

	return nil
}
