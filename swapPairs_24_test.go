package main

import "testing"

func TestSwapPairs(t *testing.T) {
	h := &ListNode{}
	h1 := h
	h1 = tailInsert(h1, 1)
	h1 = tailInsert(h1, 2)
	h1 = tailInsert(h1, 3)
	h1 = tailInsert(h1, 4)

	hN := swapPairs(h)

	print(hN)
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var p1, p2, p3 *ListNode

	p1 = head
	p2 = head.Next
	p3 = p2.Next
	p2.Next = p1
	p1.Next = p3

	for p3.Next != nil {

		p1 = p3
		p3 = p3.Next.Next

		p1.Next.Next = p1
		p2.Next = p1.Next
		p1.Next = p3

		p2 = p1

	}

	return head

}
