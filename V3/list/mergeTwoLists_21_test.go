package list

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	l1 := &ListNode{}
	AddNode(l1, 1)
	AddNode(l1, 2)
	AddNode(l1, 4)
	AddNode(l1, 5)

	l2 := &ListNode{}
	AddNode(l2, 1)
	AddNode(l2, 3)
	AddNode(l2, 4)

	res := mergeTwoLists(l1.Next, l2.Next)

	PrintNode(res)

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{}
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
