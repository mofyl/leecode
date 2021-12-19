package list

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestDetectCycle(t *testing.T) {

	l1 := &tools.ListNode{Val: 1}
	l2 := &tools.ListNode{Val: 2}
	l3 := &tools.ListNode{Val: 3}
	l4 := &tools.ListNode{Val: 4}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	l4.Next = l2

	res := detectCycle(l1)

	fmt.Println(res.Val)

}

func detectCycle(head *tools.ListNode) *tools.ListNode {

	if head == nil {
		return head
	}

	slow := head
	fast := head

	for {
		// 这里表示没有环
		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}

	}
	// 此时slow 在p点 我们新建一个p指向head
	// p和slow 以相同的步伐同时走，
	// 当p和slow 相遇的时候 就到了入环点
	p := head

	for p != slow {
		p = p.Next
		slow = slow.Next
	}

	return p

}
