package list

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {

	headA := &tools.ListNode{}
	headB := &tools.ListNode{}

	l1A := &tools.ListNode{Val: 1}
	l2A := &tools.ListNode{Val: 2}
	//l3A := &ListNode{Val: 3}

	headA.Next = l1A
	l1A.Next = l2A
	//l2A.Next = l3A

	l4B := &tools.ListNode{Val: 4}
	l5B := &tools.ListNode{Val: 5}

	headB.Next = l4B
	l4B.Next = l5B

	l6 := &tools.ListNode{Val: 6}
	l7 := &tools.ListNode{Val: 7}

	l6.Next = l7

	//l3A.Next = l6
	l2A.Next = l6
	l5B.Next = l6

	p := getIntersectionNode(headA.Next, headB.Next)

	fmt.Println(p)
}

func getIntersectionNode(headA, headB *tools.ListNode) *tools.ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	p, q := headA, headB

	for p != q {

		if p.Next == nil && q.Next == nil {
			return nil
		}

		p = p.Next
		q = q.Next

		if p == nil {
			p = headB
		}

		if q == nil {
			q = headA
		}

	}

	return q
}
