package list

import (
	"leecode/tools"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	head := &tools.ListNode{}
	tools.AddNode(head, 1)
	tools.AddNode(head, 2)
	tools.AddNode(head, 3)
	tools.AddNode(head, 4)
	tools.AddNode(head, 5)
	m, n := 1, 2

	res := reverseBetween(head.Next, m, n)

	tools.PrintNode(res)
}

func reverseBetween(head *tools.ListNode, m int, n int) *tools.ListNode {
	if head == nil || m == n {
		return head
	}

	prev := &tools.ListNode{}
	prev.Next = head
	prevP := prev
	cur := head

	count := 0
	for cur != nil {

		count++

		if count == m {
			// 当count == n 的时候 before 指向 反转串的最后一个
			var before *tools.ListNode
			p := cur
			var next *tools.ListNode
			// 这里一定要是等于
			for count <= n {
				next = p.Next

				p.Next = before

				before = p
				p = next
				count++
			}
			prevP.Next = before
			cur.Next = next

			break
		}

		prevP = cur
		cur = cur.Next
	}
	return prev.Next
}
