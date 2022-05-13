package combination_sum

import (
	"fmt"
	"testing"
)

type Input struct {
	candidates []int
	target     int
}

func TestCombinationSum(t *testing.T) {
	testcases := []struct {
		input  Input
		answer [][]int
	}{
		{input: Input{candidates: []int{2, 3, 6, 7}, target: 7}, answer: [][]int{{2, 2, 3}, {7}}},
		{input: Input{candidates: []int{2, 3, 5}, target: 8}, answer: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{input: Input{candidates: []int{2}, target: 1}, answer: [][]int{}},
	}

	for tcIdx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", tcIdx), func(t *testing.T) {
			actual := CombinationSum(tc.input.candidates, tc.input.target)
			if len(actual) != len(tc.answer) {
				t.Fatalf("Got #%d combinations want %d", len(actual), len(tc.answer))
			}
			for i := 0; i < len(actual); i++ {
				s := actual[i]
				sum := SumSliceInts(s)
				if sum != tc.input.target {
					t.Fatalf("#%d combination doesn't add up to the target %d", i, tc.input.target)
				}
			}
		})
	}
}
