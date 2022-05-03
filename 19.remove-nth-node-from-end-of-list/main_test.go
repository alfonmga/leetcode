package remove_nth_node_from_end_of_list

import (
	"fmt"
	"testing"
)

type Input struct {
	head *ListNode
	n    int
}

func TestRemoveNthFromEnd(t *testing.T) {
	var tests = []struct {
		input Input
		ans   *ListNode
	}{
		{input: Input{head: createLinkedListAndReturnHeadListNode([]int{1, 2, 3, 4, 5}), n: 2}, ans: createLinkedListAndReturnHeadListNode([]int{1, 2, 3, 5})},
	}

	for _, tt := range tests {
		inputNodesVals := listNodesValueToArrayInt(tt.input.head)
		t.Run(fmt.Sprintf("%v %d", inputNodesVals, tt.input.n), func(t *testing.T) {
			actual := RemoveNthFromEnd(tt.input.head, tt.input.n)
			actualV := listNodesValueToArrayInt(actual)
			ansV := listNodesValueToArrayInt(tt.ans)
			if len(actualV) != len(ansV) {
				t.Fatalf("got %v, want %v", actualV, ansV)
			}
			for i := 0; i < len(ansV); i++ {
				if actualV[i] != ansV[i] {
					t.Fatalf("got %v, want %v", actualV, ansV)
				}
			}
		})
	}
}

func createLinkedListAndReturnHeadListNode(nodesV []int) *ListNode {
	list := make([]ListNode, len(nodesV))
	for i := 0; i < len(list); i++ {
		var nextNode *ListNode
		if i != len(list)-1 {
			nextNode = &list[i+1]
		}
		list[i].Val = nodesV[i]
		list[i].Next = nextNode
	}
	return &list[0]
}
func listNodesValueToArrayInt(head *ListNode) []int {
	var listV []int
	n := head
	for n != nil {
		listV = append(listV, n.Val)
		if n.Next != nil {
			n = n.Next
		} else {
			n = nil
		}
	}
	return listV
}
