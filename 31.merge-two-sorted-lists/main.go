package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func LinkedListValuesToSlice(head *ListNode) []int {
	v := []int{}
	n := head
	for n != nil {
		v = append(v, head.Val)
		n = n.Next
	}
	return v
}

// https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1V := LinkedListValuesToSlice(list1)
	l2V := LinkedListValuesToSlice(list1)

	l = append(l1V[0:])

	return
}
