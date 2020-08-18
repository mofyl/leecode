package V1

import "testing"

func TestDetectCycle(t *testing.T) {

}

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return fast
		}
	}

	return nil

}
