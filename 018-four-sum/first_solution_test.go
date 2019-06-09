package four_sum

import "testing"
import "github.com/stretchr/testify/assert"

func TestThreeSum(t *testing.T) {
	ast := assert.New(t)

	cases := []struct {
		nums   []int
		target int
		answer [][]int
	}{

		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			answer: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
	}

	for _, c := range cases {
		result := fourSum(c.nums, c.target)
		ast.Equal(c.answer, result, "input %v, target %d", c.nums, c.target)
	}

}
