package remove_element

import "testing"

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		nums []int
		val  int
		ans  int
	}{

		{
			[]int{0, 3, 1, 2, 1},
			1,
			3,
		},

		{
			[]int{4, 4, 1, 1, 2, 2, 2, 4, 5, 6},
			4,
			7,
		},
	}

	for i, c := range cases {
		result := removeElement(c.nums, c.val)
		if result != c.ans {
			t.Errorf("case %d failed, expected answer %d, got result %d\n", i, c.ans, result)
		}
	}

}
