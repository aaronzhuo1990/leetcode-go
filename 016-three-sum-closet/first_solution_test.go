package remove_duplicates_from_sorted_array

import "testing"
import "github.com/stretchr/testify/assert"

func TestThreeSumCloset(t *testing.T) {
	ast := assert.New(t)

	cases := []struct {
		nums   []int
		target int
		ans    int
	}{

		{
			[]int{-1, 2, 1, -4},
			1,
			2,
		},

		{
			[]int{-1, 2, 1, -4, 5},
			4,
			3,
		},
	}

	for _, c := range cases {
		result := threeSumClosest(c.nums, c.target)
		ast.Equal(c.ans, result, "input %v, target: %d", c.nums, c.target)
	}

}
