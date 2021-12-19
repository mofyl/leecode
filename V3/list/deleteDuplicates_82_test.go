package list

import (
	"leecode/tools"
	"testing"
)

func TestDeleteDuplicates_V2(t *testing.T) {
	p := &tools.ListNode{}
	tools.AddNode(p, 1)
	tools.AddNode(p, 1)
	tools.AddNode(p, 1)
	tools.AddNode(p, 2)
	tools.AddNode(p, 3)
	//AddNode(p, 3)

	res := deleteDuplicates_V2(p.Next)

	tools.PrintNode(res)
}

func deleteDuplicates_V2(head *tools.ListNode) *tools.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	newHead := &tools.ListNode{
		Next: head,
	}

	prev := newHead

	for prev != nil && prev.Next != nil {

		cur := prev.Next

		// 这里表示 若cur 已经到了最后一个 || cur 的值和 下一个节点的值不同 那么就移动prev
		if cur.Next == nil || cur.Val != cur.Next.Val {
			prev = cur
		} else {
			// 这里表示 cur 不是最后一个 && cur 的值和 下一个节点的值相同
			// 这时 就去 跳过这些相同的节点
			// 循环结束时 cur 停在 最后一个重复节点上
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next

		}

	}

	return newHead.Next
}
