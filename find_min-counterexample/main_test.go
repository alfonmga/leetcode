package find_min_counterxample

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMinCounterexample(t *testing.T) {
	testcases := []struct {
		input  int
		answer []int
	}{
		{4, []int{4, 5, 6, 7}},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%d", tc.input), func(t *testing.T) {
			actual := FindMinCounterexample(tc.input)
			if !reflect.DeepEqual(actual, tc.answer) {
				t.Errorf("Got %v want %v", actual, tc.answer)
			}
		})

	}
}
