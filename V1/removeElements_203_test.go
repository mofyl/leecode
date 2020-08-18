package V1

import (
	"testing"
)

func TestRemoveEle(t *testing.T) {
	h := &ListNode{}

	p := h

	p = tailInsert(p, 1)
	// p = tailInsert(p, 1)
	p = tailInsert(p, 2)
	p = tailInsert(p, 6)
	p = tailInsert(p, 3)
	p = tailInsert(p, 4)
	p = tailInsert(p, 5)
	p = tailInsert(p, 6)

	// print(h)
	p1 := removeElements(h.Next, 6)

	print(p1)
}

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return head
	}

	tH := &ListNode{}

	tH.Next = head

	p := tH

	for p != nil && p.Next != nil {

		if p.Next.Val == val {
			if p.Next != nil {
				p.Next = p.Next.Next
			} else {
				p.Next = nil
			}

			continue
		}
		p = p.Next
	}

	return tH.Next
}
