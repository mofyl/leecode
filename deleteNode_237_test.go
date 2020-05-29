package main

import "testing"

func TestDeleteNode(t *testing.T) {

	h := &ListNode{}

	p := h
	p = tailInsert(p, 4)
	p1 := tailInsert(p, 5)
	p = tailInsert(p1, 1)
	p = tailInsert(p, 9)

	// print(h)

	deleteNode(p1)

	print(h)
}

func deleteNode(node *ListNode) {

	if node.Next != nil {

		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}

}
