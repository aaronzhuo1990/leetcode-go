package search_insert_position

import "testing"

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		ans    int
	}{

		{
			[]int{1, 3, 5, 6},
			2,
			1,
		},

		{
			[]int{1, 3, 5, 6},
			7,
			4,
		},
	}

	for i, c := range cases {
		result := searchInsert(c.nums, c.target)
		if result != c.ans {
			t.Errorf("case %d failed, expected answer %d, got result %d\n", i, c.ans, result)
		}
	}

}
