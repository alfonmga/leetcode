package merge_k_sorted_lists

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return &ListNode{}
	}

	var valuesFromLists []int
	for i := 0; i < len(lists); i++ {
		n := lists[i]
		for n != nil {
			valuesFromLists = append(valuesFromLists, n.Val)
			n = n.Next
		}
	}
	sort.Ints(valuesFromLists)

	var linkedList = make([]ListNode, len(valuesFromLists))
	for i := 0; i < len(valuesFromLists); i++ {
		var nextN *ListNode
		if i != len(valuesFromLists)-1 {
			nextN = &linkedList[i+1]
		}
		linkedList[i].Val = valuesFromLists[i]
		linkedList[i].Next = nextN
	}

	return &linkedList[0]
}
