package list

import "testing"

func TestOddEvenList(t *testing.T) {
	head := &ListNode{}

	AddNode(head, 2)
	AddNode(head, 1)
	AddNode(head, 3)
	AddNode(head, 5)
	AddNode(head, 6)
	AddNode(head, 4)
	AddNode(head, 7)

	//res := oddEvenList(head.Next)
	res := oddEvenList_v2(head.Next)

	PrintNode(res)
}

func oddEvenList(head *ListNode) *ListNode {
	// 只有 head为空 或 只有一个节点  或 只有2个节点
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	// 保存 奇数节点
	h1 := &ListNode{}
	// 保存偶数节点
	h2 := &ListNode{}

	count := 1

	p := head
	p1 := h1
	p2 := h2

	for p != nil {
		if count%2 != 0 {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}

		p = p.Next
		count++
	}

	p1.Next = h2.Next
	p2.Next = nil

	return h1.Next

}

func oddEvenList_v2(head *ListNode) *ListNode {
	// 只有 head为空 或 只有一个节点  或 只有2个节点
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	odd := head

	event := odd.Next
	eventHead := event

	for event != nil && event.Next != nil {

		odd.Next = event.Next
		odd = odd.Next

		event.Next = odd.Next
		event = event.Next
	}

	odd.Next = eventHead

	return head

}
