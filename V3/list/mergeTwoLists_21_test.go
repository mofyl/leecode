package list

import (
	"leecode/tools"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	l1 := &tools.ListNode{}
	tools.AddNode(l1, 1)
	tools.AddNode(l1, 2)
	tools.AddNode(l1, 4)
	tools.AddNode(l1, 5)

	l2 := &tools.ListNode{}
	tools.AddNode(l2, 1)
	tools.AddNode(l2, 3)
	tools.AddNode(l2, 4)

	res := mergeTwoLists(l1.Next, l2.Next)

	tools.PrintNode(res)

}

func mergeTwoLists(l1 *tools.ListNode, l2 *tools.ListNode) *tools.ListNode {

	res := &tools.ListNode{}
	p := res
	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else if l1.Val == l2.Val {
			p.Next = l1
			l1 = l1.Next
			p = p.Next

			p.Next = l2
			l2 = l2.Next

		} else if l1.Val > l2.Val {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	for l1 != nil {
		p.Next = l1
		l1 = l1.Next
		p = p.Next
	}
	for l2 != nil {
		p.Next = l2
		l2 = l2.Next
		p = p.Next
	}
	return res.Next
}
