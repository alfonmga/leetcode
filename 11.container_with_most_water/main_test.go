package container_with_most_water

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	var tests = []struct {
		input []int
		ans   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			actual := MaxArea(tt.input)
			if actual != tt.ans {
				t.Errorf("got %v, want %v", actual, tt.ans)
			}
		})
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	}
}
