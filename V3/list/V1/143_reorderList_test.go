/*
 * @Author: lirui
 * @Date: 2022-01-04 14:51:16
 * @LastEditTime: 2022-01-04 15:36:35
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\143_reorderList_test.go
 */

package V1

import "testing"

func TestreorderList(t *testing.T) {

}

func reorderList(head *ListNode) {

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

func reverseListHelper_143(head *ListNode) *ListNode {

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
