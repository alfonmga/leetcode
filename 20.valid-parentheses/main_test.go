package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	testcases := []struct {
		input  string
		answer bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"{{}[][[[]]]}", true},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			actual := IsValid(tc.input)
			if actual != tc.answer {
				t.Errorf("Got %v want %v", actual, tc.answer)
			}
		})
	}
}
