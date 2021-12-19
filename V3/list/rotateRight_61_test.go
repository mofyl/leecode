package list

import (
	"leecode/tools"
	"testing"
)

func TestRotateRight(t *testing.T) {
	p := &tools.ListNode{}

	tools.AddNode(p, 1)
	tools.AddNode(p, 2)
	tools.AddNode(p, 3)
	//AddNode(p, 4)
	//AddNode(p, 5)

	res := rotateRight(p.Next, 3)

	tools.PrintNode(res)
}

func rotateRight(head *tools.ListNode, k int) *tools.ListNode {

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

	newHead := &tools.ListNode{}

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
