package merge_k_sorted_lists

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeKList(t *testing.T) {
	testcases := []struct {
		input  []*ListNode
		answer *ListNode
	}{
		{input: []*ListNode{createListNodeFromSliceInt([]int{1, 4, 5}), createListNodeFromSliceInt([]int{1, 3, 4}), createListNodeFromSliceInt([]int{2, 6})}, answer: createListNodeFromSliceInt([]int{1, 1, 2, 3, 4, 4, 5, 6})},
		{input: []*ListNode{}, answer: &ListNode{}},
	}
	for tcIdx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", tcIdx), func(t *testing.T) {
			actual := MergeKLists(tc.input)
			actualVals := linkedListValuesToSliceInt(actual)
			answerVals := linkedListValuesToSliceInt(tc.answer)
			if !reflect.DeepEqual(actualVals, answerVals) {
				t.Errorf("Got %v expected %v", actualVals, answerVals)
			}
		})
	}
}

func createListNodeFromSliceInt(s []int) *ListNode {
	linkedList := make([]ListNode, len(s))
	for i := 0; i < len(s); i++ {
		var next *ListNode
		if i != len(s)-1 {
			next = &linkedList[i+1]
		}
		linkedList[i].Val = s[i]
		linkedList[i].Next = next
	}
	return &linkedList[0]
}
func linkedListValuesToSliceInt(linkedList *ListNode) []int {
	var vals []int
	next := linkedList
	for next != nil {
		vals = append(vals, next.Val)
		next = next.Next
	}
	return vals
}
