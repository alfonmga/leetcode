package maximum_subarray

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	testcases := []struct {
		input  []int
		answer int
	}{
		{input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, answer: 6},
		{input: []int{5, 4, -1, 7, 8}, answer: 23},
		{input: []int{1}, answer: 1},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d %v", idx, tc.input), func(t *testing.T) {
			actual := MaxSubArray(tc.input)
			if actual != tc.answer {
				t.Errorf("Got %d want %d", actual, tc.answer)
			}
		})
	}
}
