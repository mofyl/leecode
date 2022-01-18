package V1

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {

	a1 := &ListNode{Val: 1}
	a2 := &ListNode{Val: 2}

	a1.Next = a2

	b1 := &ListNode{Val: 1}
	b2 := &ListNode{Val: 2}
	b3 := &ListNode{Val: 3}

	b1.Next = b2
	b2.Next = b3

	c1 := &ListNode{Val: 4}
	c2 := &ListNode{Val: 5}
	c3 := &ListNode{Val: 6}

	c1.Next = c2
	c2.Next = c3

	a2.Next = c1
	b3.Next = c1

	res := getIntersectionNode(a1, b1)

	fmt.Println(res)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	aLen := getLength(headA)
	bLen := getLength(headB)

	if aLen > bLen {
		// 先让A走，A走 aLen-bLen 这么多步数
		totalStep := aLen - bLen
		curStep := 0
		for headA != nil && curStep < totalStep {
			curStep++
			headA = headA.Next
		}

	} else {
		totalStep := bLen - aLen
		curStep := 0
		for headB != nil && curStep < totalStep {
			curStep++
			headB = headB.Next
		}
	}

	for headA != nil && headB != nil {

		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func getLength(h *ListNode) int {

	if h == nil {
		return 0
	}

	p := h
	res := 0
	for p != nil {
		p = p.Next
		res++
	}

	return res

}
