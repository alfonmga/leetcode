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
	}{}
	for tcIdx, tc := range testcases {
		t.Run(fmt.Sprintf("#%d", tcIdx), func(t *testing.T) {
			actual := MergeKLists(tc.input)
			actualVals := linkedListValuesToSliceInt(actual)
			answerVals := linkedListValuesToSliceInt(tc.answer)
			if !reflect.DeepEqual(actualVals, answerVals) {
				t.Errorf("Got %v expected %v", actual, tc.answer)
			}
		})
	}
}

func sliceIntToLinkedList(s []int) *ListNode {
	var linkedListNodes []*ListNode
	for i := 0; i < len(s); i++ {
		linkedListNodes = append(linkedListNodes, &ListNode{Val: s[i]})
	}
	return linkedListNodes[0]
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
