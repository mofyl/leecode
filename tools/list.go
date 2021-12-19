package tools

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddNode(l *ListNode, val int) {

	p := l

	for p.Next != nil {
		p = p.Next
	}

	tmp := &ListNode{
		Val: val,
	}

	p.Next = tmp
}

func PrintNode(l *ListNode) {
	if l == nil {
		fmt.Println("nil")
	}
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}
