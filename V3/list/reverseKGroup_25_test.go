package list

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	p := &ListNode{}
	AddNode(p, 1)
	AddNode(p, 2)
	AddNode(p, 3)
	AddNode(p, 4)

	//res := reverseKGroup(p.Next, 1)
	res := reverseKGroup_V2(p.Next, 2)

	PrintNode(res)
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil || k <= 1 {
		return head
	}

	p := head
	listL := 1

	for p.Next != nil {
		p = p.Next
		listL++
	}

	if listL < k {
		return head
	}

	totalCount := listL / k

	count := 1

	newHead := &ListNode{}
	newHead.Next = head

	prev := newHead
	first := prev.Next

	for count <= totalCount && first != nil && first.Next != nil {

		idx := 1
		second := first.Next
		// 这里是 先换两个节点
		prev.Next = second
		first.Next = second.Next
		second.Next = first
		idx++

		for idx < k && first.Next != nil {
			// 再去换后面的节点
			third := prev.Next
			second = first.Next

			prev.Next = second
			first.Next = second.Next
			second.Next = third
			idx++
		}

		prev = first
		first = first.Next
		count++

	}
	return newHead.Next
}

func reverseKGroup_V2(head *ListNode, k int) *ListNode {

	if head == nil || k <= 1 {
		return head
	}

	newHead := &ListNode{}
	newHead.Next = head

	prev := newHead
	tail := newHead

	for tail.Next != nil {

		count := 0

		for count < k && tail != nil {
			count++
			tail = tail.Next
		}

		if tail == nil {
			break
		}

		// 这里使用倒转的方式
		start := prev.Next
		next := tail.Next

		tail.Next = nil

		prev.Next = reverse(start)
		start.Next = next

		prev = start
		tail = start

	}
	return newHead.Next
}

func reverse(head *ListNode) *ListNode {

	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next

	}
	return prev
}
