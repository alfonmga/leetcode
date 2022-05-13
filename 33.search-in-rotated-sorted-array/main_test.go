package search_in_rotated_sorted_array

import (
	"fmt"
	"testing"
)

type Input struct {
	nums   []int
	target int
}

func TestSearch(t *testing.T) {
	testcases := []struct {
		input  Input
		answer int
	}{
		{input: Input{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, answer: 4},
		{input: Input{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3}, answer: -1},
		{input: Input{nums: []int{1}, target: 0}, answer: -1},
	}

	for tcIdx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", tcIdx), func(t *testing.T) {
			actual := Search(tc.input.nums, tc.input.target)
			if actual != tc.answer {
				t.Errorf("Got %d want %d", actual, tc.answer)
			}
		})
	}
}
