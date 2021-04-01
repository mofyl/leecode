package V1

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2}
	l2 := &ListNode{Val: 4}
	l3 := &ListNode{Val: 3}
	l1.Next = l2
	l2.Next = l3

	l4 := &ListNode{Val: 5}
	l5 := &ListNode{Val: 6}
	l6 := &ListNode{Val: 4}
	l4.Next = l5
	l5.Next = l6

	res := addTwoNumbers(l1, l4)
	print(res)
}

func print(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}

	n := 0 // 该数字表示 是否有进位 若有则有值
	res := &ListNode{}
	head := res
	for l1 != nil && l2 != nil {
		var node *ListNode
		n, node = getListNode(l1.Val, l2.Val, n)
		head.Next = node

		head = head.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		for l1 != nil {
			var node *ListNode
			n, node = getListNode(l1.Val, n, 0)
			head.Next = node

			head = head.Next
			l1 = l1.Next
		}
	}

	if l2 != nil {
		for l2 != nil {
			var node *ListNode
			n, node = getListNode(l2.Val, n, 0)
			head.Next = node

			head = head.Next
			l2 = l2.Next
		}

	}

	if n != 0 {
		tmp := &ListNode{
			Val:  n,
			Next: nil,
		}
		head.Next = tmp
	}

	return res.Next
}

func getListNode(n1, n2, n3 int) (n int, node *ListNode) {
	sum := n1 + n2 + n3

	if sum >= 10 {
		n = sum / 10
		sum %= 10
	} else {
		n = 0
	}

	node = &ListNode{
		Val:  sum,
		Next: nil,
	}
	return
}
