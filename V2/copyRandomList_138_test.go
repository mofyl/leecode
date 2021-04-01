package V2

import (
	"fmt"
	"testing"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func TestCopyRandomList(t *testing.T) {

	l1 := &Node{Val: 7}
	l2 := &Node{Val: 13}
	l3 := &Node{Val: 11}
	l4 := &Node{Val: 10}
	l5 := &Node{Val: 1}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	l1.Random = nil
	l2.Random = l1
	l3.Random = l5
	l4.Random = l3
	l5.Random = l1

	h1 := copyRandomList(l1)

	for l1 != nil {
		fmt.Println("Val ", l1.Val)
		if l1.Random != nil {

			fmt.Println("Random ", l1.Random.Val)
		}
		l1 = l1.Next
	}

	fmt.Println("===============")

	for h1 != nil {
		fmt.Println("Val", h1.Val)
		if h1.Random != nil {
			fmt.Println("Random ", h1.Random.Val)
		}
		h1 = h1.Next
	}
}

func copyRandomList_V1(head *Node) *Node {

	if head == nil {
		return nil
	}

	m := make(map[*Node]*Node)

	p := head

	for p != nil {
		newNode := &Node{Val: p.Val}
		m[p] = newNode
		p = p.Next
	}

	p = head
	for p != nil {
		v := m[p]
		if v != nil {
			v.Next = m[p.Next]
			if p.Random != nil {
				v.Random = m[p.Random]
			}
		}
		p = p.Next
	}
	return m[head]
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		return head
	}

	p := head
	// 在head的每个节点后面增加一个新的节点
	for p != nil {

		tmp := p.Next
		newNode := &Node{
			Val: p.Val,
		}

		newNode.Next = p.Next
		p.Next = newNode

		p = tmp
	}

	p = head
	// 将random 节点 复制给 新增的节点
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	// 将两个链表分开
	p = head
	q := p.Next
	res := q
	for q != nil && q.Next != nil {
		p.Next = q.Next
		q.Next = q.Next.Next
		q = q.Next
		p = p.Next
	}

	p.Next = nil

	return res

}
