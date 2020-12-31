package list

import "testing"

func TestPartition(t *testing.T) {
	head := &ListNode{}

	AddNode(head, 1)
	//AddNode(head, 4)
	//AddNode(head, 3)
	//AddNode(head, 2)
	//AddNode(head, 5)
	//AddNode(head, 2)

	//k := 3
	k := 0

	res := partition(head.Next, k)

	PrintNode(res)
}

func partition(head *ListNode, x int) *ListNode {

	if head == nil {
		return head
	}

	lessHead := &ListNode{}
	lessP := lessHead
	hugeHead := &ListNode{}
	hugeP := hugeHead

	p := head

	for p != nil {

		if p.Val < x {
			lessP.Next = p
			lessP = lessP.Next
		} else {
			hugeP.Next = p
			hugeP = hugeP.Next
		}
		p = p.Next
	}

	newHead := &ListNode{}
	lessP.Next = hugeHead.Next
	newHead.Next = lessHead.Next
	hugeP.Next = nil
	return newHead.Next
}
