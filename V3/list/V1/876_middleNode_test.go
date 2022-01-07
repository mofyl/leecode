/*
 * @Author: lirui
 * @Date: 2022-01-04 18:06:30
 * @LastEditTime: 2022-01-04 18:11:12
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\876_middleNode_test.go
 */

package V1

import (
	"fmt"
	"testing"
)

func TestMiddleNode(t *testing.T) {

	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l6 := &ListNode{Val: 6}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6

	res := middleNode(l1)

	fmt.Println(res.Val)

}

func middleNode(head *ListNode) *ListNode {

	// 返回 右中点
	f := head
	s := head

	for f != nil && f.Next != nil {

		f = f.Next.Next
		s = s.Next
	}

	return s

}
