/*
 * @Author: lirui
 * @Date: 2022-01-04 16:51:54
 * @LastEditTime: 2022-01-04 17:13:15
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\23_mergeKLists_test.go
 */

package V1

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {

	l11 := &ListNode{Val: 1}
	l14 := &ListNode{Val: 4}
	l15 := &ListNode{Val: 5}

	l11.Next = l14
	l14.Next = l15

	l21 := &ListNode{Val: 1}
	l23 := &ListNode{Val: 3}
	l24 := &ListNode{Val: 4}

	l21.Next = l23
	l23.Next = l24

	l32 := &ListNode{Val: 2}
	l36 := &ListNode{Val: 6}

	l32.Next = l36

	lists := make([]*ListNode, 0)
	lists = append(lists, l11)
	lists = append(lists, l21)
	lists = append(lists, l32)

	res := mergeKLists(lists)

	PrintNode(res)
}

func mergeKLists(lists []*ListNode) *ListNode {

	return mergeKListsHelper(lists, 0, len(lists)-1)
}

func mergeKListsHelper(lists []*ListNode, l, r int) *ListNode {

	if l == r {
		return lists[l]
	}

	if l > r {
		return nil
	}

	mid := (l + r) >> 1

	return mergeTwoLists(mergeKListsHelper(lists, l, mid), mergeKListsHelper(lists, mid+1, r))

}
