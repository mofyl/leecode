package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddToEnd(l *ListNode, val int) {
	tmp := &ListNode{
		Val:  val,
		Next: nil,
	}

	l.Next = tmp
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

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}
