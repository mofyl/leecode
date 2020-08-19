package V2

import (
	"fmt"
	"testing"
)

/*
	妖魔化的双指针！
	这里我们知道的条件就是要删除倒数第几个 n，还有就是当前节点是否为空 或 当前节点的下一个是否为空
	那么我们就可以使用两个指针，保持这两个指针之间的距离 = n，当前一个指针的next为空时
	perv.next 就指向了要删除的节点

*/

func TestRemoveNthFromEnd(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	head := removeNthFromEnd(l1, 2)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	prev := &ListNode{Next: head}

	l1, l2 := prev, prev
	num1 := 0
	num2 := 0

	for l2.Next != nil {

		if num2-num1 >= n {
			num1++
			l1 = l1.Next
		}

		num2++
		l2 = l2.Next
	}

	l1.Next = l1.Next.Next

	return prev.Next

}
