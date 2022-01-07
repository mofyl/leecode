/*
 * @Author: lirui
 * @Date: 2022-01-03 16:25:29
 * @LastEditTime: 2022-01-04 13:52:06
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\2_addTwoNumbers_test.go
 */

package V1

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	// l2 := &ListNode{Val: 2}
	// l4 := &ListNode{Val: 4}
	// l3 := &ListNode{Val: 3}

	// l2.Next = l4
	// l4.Next = l3

	// l5 := &ListNode{Val: 5}
	// l6 := &ListNode{Val: 6}
	// l41 := &ListNode{Val: 4}

	// l5.Next = l6
	// l6.Next = l41

	// node := addTwoNumbers(l2, l5)

	// l0 := &ListNode{Val: 0}
	// l01 := &ListNode{Val: 0}

	// node := addTwoNumbers(l0, l01)

	l91 := &ListNode{Val: 9}
	l92 := &ListNode{Val: 9}
	l93 := &ListNode{Val: 9}
	l94 := &ListNode{Val: 9}
	l95 := &ListNode{Val: 9}
	l96 := &ListNode{Val: 9}
	l97 := &ListNode{Val: 9}

	l91.Next = l92
	l92.Next = l93
	l93.Next = l94
	l94.Next = l95
	l95.Next = l96
	l96.Next = l97

	l291 := &ListNode{Val: 9}
	l292 := &ListNode{Val: 9}
	l293 := &ListNode{Val: 9}
	l294 := &ListNode{Val: 9}

	l291.Next = l292
	l292.Next = l293
	l293.Next = l294

	node := addTwoNumbers(l91, l291)

	PrintNode(node)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{}

	p := res
	carry := 0
	sum := 0
	for l1 != nil && l2 != nil {
		sum, carry = getSum(l1.Val, l2.Val, carry)

		tmp := &ListNode{Val: sum}
		p.Next = tmp

		l1 = l1.Next
		l2 = l2.Next
		p = p.Next
	}

	for l1 != nil {

		sum, carry = getSum(l1.Val, 0, carry)

		tmp := &ListNode{Val: sum}
		p.Next = tmp

		l1 = l1.Next
		p = p.Next
	}

	for l2 != nil {
		sum, carry = getSum(0, l2.Val, carry)

		tmp := &ListNode{Val: sum}
		p.Next = tmp

		l2 = l2.Next
		p = p.Next
	}

	if carry != 0 {
		tmp := &ListNode{Val: carry}
		p.Next = tmp
	}

	return res.Next

}

// 第一个返回值 表示 和，
// 第二个返回值 表示 进位
func getSum(a, b, c int) (int, int) {

	sum := a + b + c

	if sum >= 10 {
		carry := sum / 10
		sum = sum % 10

		return sum, carry
	}

	return sum, 0
}
