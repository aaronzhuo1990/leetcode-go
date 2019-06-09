package three_sum

import "testing"
import "github.com/stretchr/testify/assert"

func TestThreeSum(t *testing.T) {
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
