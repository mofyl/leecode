package list

import (
	"leecode/tools"
	"testing"
)

func TestSwapPairs(t *testing.T) {

	p := &tools.ListNode{}
	tools.AddNode(p, 1)
	tools.AddNode(p, 2)
	tools.AddNode(p, 3)
	//AddNode(p, 4)

	//res := swapPairs(p.Next)
	res := swapPairs_V2(p.Next)

	tools.PrintNode(res)
}

func swapPairs(head *tools.ListNode) *tools.ListNode {

	if head == nil {
		return nil
	}

	newHead := &tools.ListNode{
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

func swapPairs_V2(head *tools.ListNode) *tools.ListNode {

	if head == nil {
		return nil
	}

	newHead := &tools.ListNode{
		Next: head,
	}
	// pre 指向 cur的前一个
	pre := newHead
	// cur 指向 待交换的 第一个元素
	first := pre.Next
	var second *tools.ListNode
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
