/*
 * @Author: lirui
 * @Date: 2022-01-04 17:40:59
 * @LastEditTime: 2022-01-04 18:04:26
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\82_deleteDuplicates_test.go
 */

package V1

import "testing"

func TestDeleteDuplicates_82(t *testing.T) {

	// l1 := &ListNode{Val: 1}
	// l2 := &ListNode{Val: 2}
	// l3 := &ListNode{Val: 3}
	// l31 := &ListNode{Val: 3}
	// l4 := &ListNode{Val: 4}
	// l41 := &ListNode{Val: 4}
	// l5 := &ListNode{Val: 5}

	// l1.Next = l2
	// l2.Next = l3
	// l3.Next = l31
	// l31.Next = l4
	// l4.Next = l41
	// l41.Next = l5

	// l1 := &ListNode{Val: 1}
	// l11 := &ListNode{Val: 1}
	// l12 := &ListNode{Val: 1}
	// l2 := &ListNode{Val: 2}
	// l3 := &ListNode{Val: 3}

	// l1.Next = l11
	// l11.Next = l12
	// l12.Next = l2
	// l2.Next = l3

	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l21 := &ListNode{Val: 2}

	l1.Next = l2
	l2.Next = l21

	res := deleteDuplicates_82(l1)

	PrintNode(res)

}

func deleteDuplicates_82(head *ListNode) *ListNode {

	tmp := &ListNode{}

	res := tmp

	p := head

	for p != nil && p.Next != nil {

		if p.Val == p.Next.Val {

			for p != nil && p.Next != nil && p.Val == p.Next.Val {
				p.Next = p.Next.Next
			}
			p = p.Next

		} else {

			res.Next = &ListNode{Val: p.Val}
			// res.Next = p

			p = p.Next
			res = res.Next
		}

	}

	if p != nil {
		res.Next = &ListNode{Val: p.Val}
	}

	return tmp.Next

}
