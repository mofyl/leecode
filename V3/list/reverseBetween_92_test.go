package list

import "testing"

func TestReverseBetween(t *testing.T) {
	head := &ListNode{}
	AddNode(head, 1)
	AddNode(head, 2)
	AddNode(head, 3)
	AddNode(head, 4)
	AddNode(head, 5)
	m, n := 1, 2

	res := reverseBetween(head.Next, m, n)

	PrintNode(res)
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n {
		return head
	}

	prev := &ListNode{}
	prev.Next = head
	prevP := prev
	cur := head

	count := 0
	for cur != nil {

		count++

		if count == m {
			// 当count == n 的时候 before 指向 反转串的最后一个
			var before *ListNode
			p := cur
			var next *ListNode
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
