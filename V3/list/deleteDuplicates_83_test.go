package list

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	p := &ListNode{}
	//AddNode(p, 1)
	//AddNode(p, 1)
	//AddNode(p, 2)

	AddNode(p, 1)
	AddNode(p, 1)
	AddNode(p, 2)
	AddNode(p, 3)
	//AddNode(p, 3)

	res := deleteDuplicates(p.Next)

	PrintNode(res)
}

func deleteDuplicates(head *ListNode) *ListNode {

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
