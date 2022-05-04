package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var linkedList []*ListNode
	lastN := head
	for {
		linkedList = append(linkedList, lastN)
		if lastN.Next != nil {
			lastN = lastN.Next
		} else {
			break
		}
	}

	prevTargetN := linkedList[len(linkedList)-n-1]
	nextTargetN := linkedList[len(linkedList)-n+1]
	prevTargetN.Next = nextTargetN

	linkedList = append(linkedList[:len(linkedList)-n], linkedList[len(linkedList)-n+1:]...)

	return linkedList[0]
}
