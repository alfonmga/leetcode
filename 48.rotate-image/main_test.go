package rotate_image

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	testcases := []struct {
		input  [][]int
		answer [][]int
	}{
		{input: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, answer: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{input: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}, answer: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
	}
	for tcIdx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d %v", tcIdx, tc.input), func(t *testing.T) {
			_input := tc.input
			Rotate(&_input)
			for i := 0; i < len(_input); i++ {
				for j := 0; j < len(_input[i]); j++ {
					if _input[i][j] != tc.answer[i][j] {
						t.Fatalf("_input[%d] Got %v want %v", i, _input[i], tc.answer[i])
					}
				}
			}
		})
	}
}
