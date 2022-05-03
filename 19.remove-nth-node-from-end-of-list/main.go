package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return &ListNode{Val: 0, Next: nil}
}
