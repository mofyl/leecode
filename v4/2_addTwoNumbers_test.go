package v4

import (
	"leecode/tools"
	"testing"
)

func Test_AddTowNumbers(t *testing.T) {

	l1 := &tools.ListNode{Val: 9}
	l11 := &tools.ListNode{Val: 9}
	l12 := &tools.ListNode{Val: 9}
	l13 := &tools.ListNode{Val: 9}
	l14 := &tools.ListNode{Val: 9}
	l15 := &tools.ListNode{Val: 9}
	l16 := &tools.ListNode{Val: 9}

	l1.Next = l11
	l11.Next = l12
	l12.Next = l13
	l13.Next = l14
	l14.Next = l15
	l15.Next = l16

	l2 := &tools.ListNode{Val: 9}
	l21 := &tools.ListNode{Val: 9}
	l22 := &tools.ListNode{Val: 9}
	l23 := &tools.ListNode{Val: 9}

	l2.Next = l21
	l21.Next = l22
	l22.Next = l23

	res := addTwoNumbers(l1, l2)

	tools.PrintNode(res)

}

func addTwoNumbers(l1 *tools.ListNode, l2 *tools.ListNode) *tools.ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := &tools.ListNode{}
	p := head
	add := 0
	for l1 != nil || l2 != nil {

		m, n := 0, 0

		if l1 != nil {
			m = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n = l2.Val
			l2 = l2.Next
		}

		sum := m + n + add
		sum, add = sum%10, sum/10

		p.Next = &tools.ListNode{Val: sum}
		p = p.Next

	}

	if add != 0 {

		p.Next = &tools.ListNode{Val: add}
	}

	return head.Next

}
