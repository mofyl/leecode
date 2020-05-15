package main

import (
	"fmt"
	"testing"
)

var ()

func TestReverseList(t *testing.T) {
	head := &ListNode{}
	p := head

	p = tailInsert(p, 1)
	p = tailInsert(p, 2)
	p = tailInsert(p, 3)
	// p = tailInsert(p, 4)
	// p = tailInsert(p, 5)

	print(head.Next)
	fmt.Println("===================")
	h := reverseList(head.Next)

	print(h)
}

func reverseList(head *ListNode) *ListNode {

	// if head == nil || head.Next == nil {
	//   return head
	// }

	// var pre *ListNode = nil
	// cur := head
	// var nextTmp *ListNode
	// for cur != nil {

	//   nextTmp = cur.Next
	//   cur.Next = pre

	//   pre = cur
	//   cur = nextTmp
	// }

	// return pre

	if head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last

}
