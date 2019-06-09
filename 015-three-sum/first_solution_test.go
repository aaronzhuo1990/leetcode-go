package remove_duplicates_from_sorted_array

import "testing"
import "github.com/stretchr/testify/assert"

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	ast := assert.New(t)

	cases := []struct {
		nums []int
		ans  [][]int
	}{

		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}

	for _, c := range cases {
		result := threeSum(c.nums)
		ast.Equal(c.ans, result, "input %v", c.nums)
	}

}
