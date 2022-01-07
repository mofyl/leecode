/*
 * @Author: lirui
 * @Date: 2022-01-04 14:09:26
 * @LastEditTime: 2022-01-04 14:47:53
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\142_detectCycle_test.go
 */
package V1

import (
	"fmt"
	"testing"
)

func TestDetectCycle(t *testing.T) {

	l3 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 2}
	l0 := &ListNode{Val: 0}
	l4 := &ListNode{Val: 4}

	l3.Next = l2
	l2.Next = l0
	l0.Next = l4

	l4.Next = l2

	node := detectCycle(l3)

	fmt.Println(node.Val)

}

func detectCycle(head *ListNode) *ListNode {

	f := head
	s := head

	for {

		// 没有环
		if f == nil || f.Next == nil {
			return nil
		}

		f = f.Next.Next
		s = s.Next

		if f == s {
			break
		}

	}

	s = head

	for s != f {

		s = s.Next
		f = f.Next

	}

	return s

}
