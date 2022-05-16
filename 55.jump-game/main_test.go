package jump_game

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	testcases := []struct {
		input  []int
		answer bool
	}{
		// {input: []int{2, 3, 1, 1, 4}, answer: true},
		{input: []int{3, 2, 1, 0, 4}, answer: false},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d %v", idx, tc.input), func(t *testing.T) {
			actual := CanJump(tc.input)
			if actual != tc.answer {
				t.Errorf("Got %v want %v", actual, tc.answer)
			}
		})
	}
}
