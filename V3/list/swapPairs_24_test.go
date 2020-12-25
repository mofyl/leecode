package list

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {

	p := &ListNode{}
	AddNode(p, 1)
	AddNode(p, 2)
	AddNode(p, 3)
	//AddNode(p, 4)

	//res := swapPairs(p.Next)
	res := swapPairs_V2(p.Next)

	PrintNode(res)
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	newHead := &ListNode{
		Next: head,
	}
	p := newHead
	p1 := p.Next
	tmp := p1.Next
	for tmp != nil {
		p.Next = tmp
		p1.Next = tmp.Next
		tmp.Next = p1

		p = p1
		p1 = p.Next
		if p1 == nil {
			break
		}
		tmp = p1.Next
	}
	return newHead.Next
}

func swapPairs_V2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	newHead := &ListNode{
		Next: head,
	}
	// pre 指向 cur的前一个
	pre := newHead
	// cur 指向 待交换的 第一个元素
	first := pre.Next
	var second *ListNode
	for first != nil && first.Next != nil {

		second = first.Next

		pre.Next = second
		first.Next = second.Next
		second.Next = first

		pre = first
		first = first.Next

	}

	return newHead.Next

}
