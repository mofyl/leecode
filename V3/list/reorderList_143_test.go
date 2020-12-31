package list

import (
	"testing"
)

func TestReorderList(t *testing.T) {
	head := &ListNode{}
	AddNode(head, 1)
	AddNode(head, 2)
	AddNode(head, 3)
	AddNode(head, 4)

	reorderList(head.Next)

	PrintNode(head.Next)
}

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	// 先找中点  若是偶数 这里找的是 左中点
	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 这里slow 指向的就是左中点
	start := slow.Next
	slow.Next = nil
	newHead := reverse_143(start)

	p := head
	// 这里 slow 指向的为nil，但是head 还是可以找到slow
	for newHead != nil && p != nil {

		next1 := p.Next
		next2 := newHead.Next

		p.Next = newHead
		newHead.Next = next1

		p = next1
		newHead = next2
	}
}

func reverse_143(head *ListNode) *ListNode {

	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next

		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev
}
