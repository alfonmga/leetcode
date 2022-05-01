package three_sum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testcases := []struct {
		input []int
		ans   [][]int
	}{
		// {[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		// {[]int{}, [][]int{}},
		// {[]int{0}, [][]int{}},
		// {[]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}, [][]int{
		// {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}},

		{[]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}, [][]int{{-4, 1, 3}}},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			actual := ThreeSum(tc.input)
			if len(actual) != len(tc.ans) {
				t.Fatalf("Got %v want %v", actual, tc.ans)
			}
			for _, v := range actual {
				eq := false
				for _, ansV := range tc.ans {
					if equalSliceInt(v, ansV) {
						eq = true
					}
				}
				if !eq {
					t.Fatalf("Got %v want %v", actual, tc.ans)
				}
			}
		})
	}
}

func equalSliceInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
