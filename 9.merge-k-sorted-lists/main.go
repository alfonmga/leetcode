package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	return &ListNode{Val: 0, Next: nil}
}
