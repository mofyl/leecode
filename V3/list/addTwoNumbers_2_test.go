package list

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	//l1 := tools.NewList()
	//l2 := tools.NewList()
	//
	//l1.AddNode(2)
	//l1.AddNode(4)
	//l1.AddNode(3)
	//
	//l2.AddNode(5)
	//l2.AddNode(6)
	//l2.AddNode(4)

	//res := addTwoNumbers(l1.Next, l2.Next)

	//res.PrintList()

}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{}
	sum := 0
	p := res
	n := 0
	for l1 != nil || l2 != nil || n != 0 {
		firstV := 0
		secondV := 0

		if l1 != nil {
			firstV = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			secondV = l2.Val
			l2 = l2.Next
		}

		sum, n = getNodeSum(firstV, secondV, n)
		tmp := &ListNode{Val: sum}

		p.Next = tmp
		p = p.Next

	}

	return res.Next

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{}
	p := res
	cur := 0
	n := 0
	for l1 != nil && l2 != nil {

		cur, n = getNodeSum(l1.Val, l2.Val, n)

		tmp := &ListNode{
			Val:  cur,
			Next: nil,
		}

		p.Next = tmp

		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {

		cur, n = getNodeSum(l1.Val, n, 0)
		tmp := &ListNode{
			Val:  cur,
			Next: nil,
		}
		p.Next = tmp
		p = p.Next
		l1 = l1.Next
	}

	for l2 != nil {
		cur, n = getNodeSum(l2.Val, n, 0)

		tmp := &ListNode{
			Val:  cur,
			Next: nil,
		}
		p.Next = tmp
		p = p.Next

		l2 = l2.Next
	}

	for n > 0 {
		cur, n = getNodeSum(0, n, 0)

		tmp := &ListNode{
			Val:  cur,
			Next: nil,
		}
		p.Next = tmp
		p = p.Next
	}

	return res.Next
}

// 返回值的第一个为 当前位的值。 第二个为 进位的值
func getNodeSum(a, b, c int) (int, int) {

	sum := a + b + c

	if sum >= 10 {
		n := sum % 10
		sum = sum / 10
		return n, sum
	}
	return sum, 0
}
