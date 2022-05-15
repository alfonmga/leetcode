package spiral_matrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	testcases := []struct {
		input  [][]int
		answer []int
	}{
		{input: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, answer: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{input: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, answer: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", idx), func(t *testing.T) {
			actual := SpiralOrder(tc.input)
			if !reflect.DeepEqual(actual, tc.answer) {
				t.Errorf("Got %v wanted %v", actual, tc.answer)
			}
		})
	}
}
