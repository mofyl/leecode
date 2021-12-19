package list

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestHasCycle(t *testing.T) {

	//	head := &ListNode{}
	//	AddNode(head, 1)
	//	AddNode(head, 2)
	//	AddNode(head, 3)
	//	AddNode(head, 4)
	//	fmt.Println(hasCycle(head))

	l1 := &tools.ListNode{Val: 1}
	l2 := &tools.ListNode{Val: 2}
	//l3 := &ListNode{Val: 3}
	//l4 := &ListNode{Val: 4}
	//
	l1.Next = l2
	l2.Next = l1
	//l2.Next = l3
	//l3.Next = l4
	//l4.Next = l2

	fmt.Println(hasCycle(l1))

}

func hasCycle(head *tools.ListNode) bool {

	if head == nil {
		return false
	}

	fast := head
	slow := head

	for {

		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next

		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

}
