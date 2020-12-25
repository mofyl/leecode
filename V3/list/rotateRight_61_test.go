package list

import (
	"testing"
)

func TestRotateRight(t *testing.T) {
	p := &ListNode{}

	AddNode(p, 1)
	AddNode(p, 2)
	AddNode(p, 3)
	//AddNode(p, 4)
	//AddNode(p, 5)

	res := rotateRight(p.Next, 3)

	PrintNode(res)
}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || k < 1 {
		return head
	}

	// 先求出 长度
	listLen := 1

	end := head
	for end.Next != nil {
		listLen++
		end = end.Next
	}

	step := k % listLen

	if step == 0 {
		return head
	}

	newHead := &ListNode{}

	p := head
	num := listLen - step
	for num > 1 {
		p = p.Next
		num--
	}

	newHead.Next = p.Next
	end.Next = head
	p.Next = nil

	return newHead.Next

}
