package V1

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	h := &ListNode{}
	h1 := h
	h1 = tailInsert(h1, 1)
	h1 = tailInsert(h1, 2)
	h1 = tailInsert(h1, 3)
	h1 = tailInsert(h1, 4)

	print(h)
	fmt.Println("==============")
	hN := swapPairs(h.Next)

	print(hN)
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	p := &ListNode{}
	p.Next = head

	pre := p
	cur := head
	for cur != nil && cur.Next != nil {

		first := cur
		secode := cur.Next

		pre.Next = secode
		first.Next = secode.Next
		secode.Next = first

		pre = first
		cur = first.Next

	}

	return p.Next

}
