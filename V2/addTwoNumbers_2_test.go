package V2

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestAddTwoNumber(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	// l11 := &ListNode{Val: 5, Next: nil}
	// l12 := &ListNode{Val: 3}
	// l1.Next = l11
	// l11.Next = l12

	// l2 := &ListNode{Val: 5}
	// l21 := &ListNode{Val: 6}
	// l22 := &ListNode{Val: 4}
	// l2.Next = l21
	// l21.Next = l22
	l2 := &ListNode{Val: 9}
	l21 := &ListNode{Val: 9}
	l2.Next = l21

	res := addTwoNumbers(l1, l2)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	prev := &ListNode{Val: 0}

	cur := prev

	x, y, carray := 0, 0, 0
	for l1 != nil || l2 != nil {

		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		sum := x + y + carray

		carray = sum / 10
		sum = sum % 10

		cur.Next = &ListNode{Val: sum}

		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carray > 0 {
		cur.Next = &ListNode{Val: carray}
	}

	return prev.Next
}
