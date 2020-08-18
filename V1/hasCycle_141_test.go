package V1

import "testing"

func TestHasCycle(t *testing.T) {

}

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	/*
		思想： 快慢指针 快指针走2步， 慢指针走1步，若有环 快指针总会追上慢指针

	*/

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false

}
