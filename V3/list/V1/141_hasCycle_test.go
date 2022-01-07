/*
 * @Author: lirui
 * @Date: 2022-01-04 13:57:34
 * @LastEditTime: 2022-01-04 14:08:03
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\list\V1\141_hasCycle_test.go
 */

package V1

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {

	l3 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 2}
	l0 := &ListNode{Val: 0}
	l4 := &ListNode{Val: 4}

	l3.Next = l2
	l2.Next = l0
	l0.Next = l4

	l4.Next = l2

	fmt.Println(hasCycle(l3))

}

func hasCycle(head *ListNode) bool {

	// 一次走一步
	p := head
	// 一次走两步
	q := head

	for q != nil && q.Next != nil {

		p = p.Next
		q = q.Next.Next

		if p == q {
			return true
		}
	}

	return false

}
