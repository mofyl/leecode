package list

import (
	"leecode/tools"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := &tools.ListNode{}

	tools.AddNode(head, 1)
	//AddNode(head, 2)
	//AddNode(head, 3)
	//AddNode(head, 4)
	//AddNode(head, 5)

	n := 1

	l := removeNthFromEnd(head.Next, n)
	tools.PrintNode(l)
}

func removeNthFromEnd(head *tools.ListNode, n int) *tools.ListNode {

	// head 为 nil
	if head == nil {
		return head
	}

	h := &tools.ListNode{
		Next: head,
	}

	des := 0

	p1 := h
	p2 := h
	for p2.Next != nil {
		if des < n {
			p2 = p2.Next
			des++
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}

	// 这里 p1 指向的就是 倒数第n个元素
	p1.Next = p1.Next.Next

	return h.Next
}
