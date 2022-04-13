package longest_substring_without_repeating_characters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"aab", 2},
		{"dvdf", 3},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			actual := LengthOfLongestSubstring(tt.s)
			if actual != tt.ans {
				t.Errorf("got %v, want %d", actual, tt.ans)
			}
		})
	}
}
