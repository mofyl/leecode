package V1

import (
	"fmt"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	h := &ListNode{}
	h1 := h
	h1 = tailInsert(h1, 1)
	h1 = tailInsert(h1, 2)
	h1 = tailInsert(h1, 3)
	h1 = tailInsert(h1, 4)
	h1 = tailInsert(h1, 5)
	print(h)

	fmt.Println("==================")

	// h = reverseBetween(h.Next, 2, 4)
	h = reverseEndM(h.Next, 4)

	print(h)
}

func reverseEndM(head *ListNode, m int) *ListNode {

	if head == nil {
		return head
	}

	var pre *ListNode
	var cur = head

	var next *ListNode
	// var nextNext *ListNode
	step := 1
	for cur != nil && step <= m {
		next = cur.Next
		cur.Next = pre

		pre = cur
		cur = next
		step++
	}

	head.Next = cur

	return pre

}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if head == nil {
		return head
	}

	guard := &ListNode{}
	guard.Next = head

	step := 0
	pre := guard
	cur := head

	for step < m-1 {
		step++
		pre = cur
		cur = pre.Next
	}

	var next *ListNode
	for i := 0; i < n-m; i++ {
		next = cur.Next

		cur.Next = cur.Next.Next
		next.Next = pre.Next
		pre.Next = next

	}

	return guard.Next

}
