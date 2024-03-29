package list

import (
	"leecode/tools"
	"testing"
)

func TestReorderList(t *testing.T) {
	head := &tools.ListNode{}
	tools.AddNode(head, 1)
	tools.AddNode(head, 2)
	tools.AddNode(head, 3)
	tools.AddNode(head, 4)

	reorderList(head.Next)

	tools.PrintNode(head.Next)
}

func reorderList(head *tools.ListNode) {

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

func reverse_143(head *tools.ListNode) *tools.ListNode {

	var prev *tools.ListNode
	cur := head

	for cur != nil {
		next := cur.Next

		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev
}
