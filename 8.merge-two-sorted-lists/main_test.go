package merge_two_sorted_lists

import (
	"fmt"
	"reflect"
	"testing"
)

type inputData struct {
	l1 *ListNode
	l2 *ListNode
}

func TestMergeTwoLists(t *testing.T) {
	testcases := []struct {
		input  inputData
		answer []int
	}{
		{input: inputData{
			l1: createLinkedListFromSlice([]int{1, 2, 4}),
			l2: createLinkedListFromSlice([]int{1, 3, 4}),
		}, answer: []int{1, 1, 2, 3, 4, 4}},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", idx), func(t *testing.T) {
			actual := MergeTwoLists(tc.input.l1, tc.input.l2)
			if !reflect.DeepEqual(LinkedListValuesToSlice(actual), tc.answer) {
				t.Errorf("Got %v want %v", LinkedListValuesToSlice(actual), tc.answer)
			}
		})
	}
}

func createLinkedListFromSlice(s []int) *ListNode {
	list := make([]ListNode, len(s))
	for i := 0; i < len(list); i++ {
		var nextNode *ListNode
		if i != len(list)-1 {
			nextNode = &list[i+1]
		}
		list[i].Val = s[i]
		list[i].Next = nextNode
	}
	return &list[0]
}
