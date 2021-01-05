package list

import "testing"

func TestMiddleNode(t *testing.T) {

	head := &ListNode{}
	AddNode(head, 1)
	AddNode(head, 2)
	AddNode(head, 3)
	AddNode(head, 4)
	//AddNode(head, 5)

	res := middleNode(head.Next)

	PrintNode(res)
}

func middleNode(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	// 这里使用双指针法
	/*
		若是基数长度  那么slow 指向中点 fast 指向 最后一个节点
		若是偶数长度  那么slow 指向左中点 fast指向 nil
	*/
	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
