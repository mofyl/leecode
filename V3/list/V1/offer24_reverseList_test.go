/*
 * @Author: lirui
 * @Date: 2022-01-04 15:15:07
 * @LastEditTime: 2022-01-04 15:26:12
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\offer24_reverseList_test.go
 */

package V1

import "testing"

func TestReverseList(t *testing.T) {

	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	reverseList(l1)

	PrintNode(l1)
}

func reverseList(head *ListNode) {

	// 找到左中点
	f := head.Next
	s := head

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	// 将 前半段 和  后半段断开
	p := s.Next

	s.Next = nil

	// 倒转后半段
	p = reverseListHelper(p)
	q := head

	for p != nil {

		pNext := p.Next
		qNext := q.Next
		q.Next = p
		p.Next = qNext

		p = pNext
		q = qNext
	}

}

func reverseListHelper(head *ListNode) *ListNode {

	var (
		prev *ListNode
		cur  *ListNode
	)
	cur = head

	for cur != nil {

		next := cur.Next

		cur.Next = prev
		prev = cur
		cur = next

	}
	return prev
}
