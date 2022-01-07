/*
 * @Author: lirui
 * @Date: 2022-01-04 18:12:11
 * @LastEditTime: 2022-01-04 18:22:24
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\92_reverseBetween_test.go
 */

package V1

import "testing"

func TestReverseBetween(t *testing.T) {
	l1 := &ListNode{Val: 1}
	// l2 := &ListNode{Val: 2}
	// l3 := &ListNode{Val: 3}
	// l4 := &ListNode{Val: 4}
	// l5 := &ListNode{Val: 5}

	// l1.Next = l2
	// l2.Next = l3
	// l3.Next = l4
	// l4.Next = l5

	res := reverseBetween(l1, 1, 1)

	PrintNode(res)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	n := 1

	tmp := &ListNode{}

	tmp.Next = head

	p := head

	prev := tmp

	for p != nil && n < left {
		prev = p
		p = p.Next
		n++
	}

	start := p
	prev.Next = nil

	for p != nil && n < right {
		p = p.Next
		n++
	}
	next := p.Next

	end := p
	end.Next = nil

	reverseListHelper(start)

	prev.Next = end
	start.Next = next

	return tmp.Next

}
