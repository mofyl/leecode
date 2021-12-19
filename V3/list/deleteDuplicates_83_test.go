package list

import (
	"leecode/tools"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	p := &tools.ListNode{}
	//AddNode(p, 1)
	//AddNode(p, 1)
	//AddNode(p, 2)

	tools.AddNode(p, 1)
	tools.AddNode(p, 1)
	tools.AddNode(p, 2)
	tools.AddNode(p, 3)
	//AddNode(p, 3)

	res := deleteDuplicates(p.Next)

	tools.PrintNode(res)
}

func deleteDuplicates(head *tools.ListNode) *tools.ListNode {

	if head == nil {
		return head
	}

	cur := head

	for cur != nil && cur.Next != nil {

		if cur.Val == cur.Next.Val {
			// 删除当前元素
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return head
}
