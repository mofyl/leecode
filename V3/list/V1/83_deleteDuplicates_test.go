/*
 * @Author: lirui
 * @Date: 2022-01-04 17:34:19
 * @LastEditTime: 2022-01-04 17:39:34
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\83_deleteDuplicates_test.go
 */

package V1

import "testing"

func TestDeleteDuplicates(t *testing.T) {

	// l1 := &ListNode{Val: 1}
	// l11 := &ListNode{Val: 1}
	// l2 := &ListNode{Val: 2}

	// l1.Next = l11
	// l11.Next = l2

	l1 := &ListNode{Val: 1}
	l11 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l31 := &ListNode{Val: 3}

	l1.Next = l11
	l11.Next = l2
	l2.Next = l3
	l3.Next = l31

	res := deleteDuplicates(l1)

	PrintNode(res)
}

func deleteDuplicates(head *ListNode) *ListNode {

	p := head

	for p != nil && p.Next != nil {

		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}

	}

	return head
}
