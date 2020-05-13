package main

import "testing"

func TestMergToList(t *testing.T) {
	l1 := &ListNode{}
	l11 := l1

	l11 = tailInsert(l11, 1)
	l11 = tailInsert(l11, 2)
	l11 = tailInsert(l11, 4)

	l2 := &ListNode{}
	l22 := l2

	l22 = tailInsert(l22, 1)
	l22 = tailInsert(l22, 3)
	l22 = tailInsert(l22, 4)

	n := mergeTwoLists(l1.Next, l2.Next)

	print(n)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	res := &ListNode{}
	res1 := res

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			res1.Next = l1
			l1 = l1.Next
		} else {
			res1.Next = l2
			l2 = l2.Next
		}
		res1 = res1.Next
	}

	if l1 != nil {
		res1.Next = l1
	}
	if l2 != nil {
		res1.Next = l2
	}

	return res.Next

}

func tailInsert(l *ListNode, ele int) *ListNode {
	n := &ListNode{
		Next: nil,
		Val:  ele,
	}

	l.Next = n

	return n

}
