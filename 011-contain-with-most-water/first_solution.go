package container_with_most_water

func findMaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {

		// Calculate area and compare it with the current max area.
		h1, h2 := height[left], height[right]
		h := min(h1, h2)
		area := h * (right - left)

		if maxArea < area {
			maxArea = area
		}

		// Move either left or right to positive position.
		if h1 < h2 {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
