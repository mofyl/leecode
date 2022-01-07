/*
 * @Author: lirui
 * @Date: 2022-01-04 16:23:44
 * @LastEditTime: 2022-01-04 16:50:50
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\21_mergeTwoLists_test.go
 */

package V1

import "testing"

func TestMergeTwoLists(t *testing.T) {

	l11 := &ListNode{Val: 1}
	l12 := &ListNode{Val: 2}
	l14 := &ListNode{Val: 4}

	l11.Next = l12
	l12.Next = l14

	l21 := &ListNode{Val: 1}
	l22 := &ListNode{Val: 3}
	l24 := &ListNode{Val: 4}

	l21.Next = l22
	l22.Next = l24

	res := mergeTwoLists(l11, l21)

	PrintNode(res)

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	p := list1
	q := list2

	tmp := &ListNode{}
	res := tmp

	for p != nil && q != nil {

		if p.Val <= q.Val {
			res.Next = p
			p = p.Next
		} else {
			res.Next = q
			q = q.Next
		}

		res = res.Next
	}

	for p != nil {
		res.Next = p

		res = res.Next
		p = p.Next
	}

	for q != nil {
		res.Next = q

		res = res.Next
		q = q.Next
	}

	return tmp.Next
}
