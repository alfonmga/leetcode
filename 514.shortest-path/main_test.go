package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	testcases := []struct {
		input []int
		want  int
	}{
		{[]int{2, 1, 1, 3, 2, 1, 1, 3}, 3},
		{[]int{7, 5, 2, 7, 2, 7, 4, 7}, 6},
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			actual := Solution(tt.input)
			if actual != tt.want {
				t.Fatalf("got %v, want %v", actual, tt.want)
			}
		})
	}
}
