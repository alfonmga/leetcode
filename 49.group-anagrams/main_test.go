package group_anagrams

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testcases := []struct {
		input  []string
		answer [][]string
	}{
		{input: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, answer: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{input: []string{"a"}, answer: [][]string{{"a"}}},
		{input: []string{""}, answer: [][]string{{""}}},
	}
	for tcIdx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d %v", tcIdx, tc.input), func(t *testing.T) {
			actual := GroupAnagrams(tc.input)
			if len(actual) != len(tc.answer) {
				t.Fatalf("Got %d groups want %d groups", len(actual), len(tc.answer))
			}
			for i := 0; i < len(tc.answer); i++ {
				for j := 0; j < len(tc.answer[i]); j++ {
					wantWord := tc.answer[i][j]
					found := false

					for k := 0; k < len(actual); k++ {
						for l := 0; l < len(actual[k]); l++ {
							if actual[k][l] == wantWord {
								found = true
								break
							}
						}
					}
					if !found {
						t.Fatalf("Want word %v", wantWord)
					}

				}
			}
		})
	}
}
