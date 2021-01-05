package list

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {

	headA := &ListNode{}
	headB := &ListNode{}

	l1A := &ListNode{Val: 1}
	l2A := &ListNode{Val: 2}
	//l3A := &ListNode{Val: 3}

	headA.Next = l1A
	l1A.Next = l2A
	//l2A.Next = l3A

	l4B := &ListNode{Val: 4}
	l5B := &ListNode{Val: 5}

	headB.Next = l4B
	l4B.Next = l5B

	l6 := &ListNode{Val: 6}
	l7 := &ListNode{Val: 7}

	l6.Next = l7

	//l3A.Next = l6
	l2A.Next = l6
	l5B.Next = l6

	p := getIntersectionNode(headA.Next, headB.Next)

	fmt.Println(p)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

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
