package list

import "testing"

func TestInsertionSortList(t *testing.T) {
	head := &ListNode{}
	AddNode(head, -1)
	AddNode(head, 5)
	AddNode(head, 3)
	AddNode(head, 4)
	AddNode(head, 0)

	res := insertionSortList(head.Next)

	PrintNode(res)
}

func insertionSortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	/*
		这里我们需要比较相邻的两个节点，p,q 的值。 若 p > q 则需要移动 q。
		移动之前需要将 q先从链表中删除，这样可以防止 一些别的问题 删除后 然后在遍历已排序的序列，找到q的位置
	*/

	newHead := &ListNode{}
	newHead.Next = head

	cur := head

	for cur != nil {
		next := cur.Next
		if next == nil {
			break
		}
		if next.Val >= cur.Val {
			cur = cur.Next
		} else {
			// 将 next 从 链表中删除
			cur.Next = next.Next
			prev := newHead
			p := newHead.Next
			for p.Val < next.Val && p != cur {
				prev = p
				p = p.Next
			}
			prev.Next = next
			next.Next = p
		}

	}
	return newHead.Next
}
