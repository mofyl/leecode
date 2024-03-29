package list

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	head := &tools.ListNode{}
	tools.AddNode(head, 1)
	tools.AddNode(head, 3)
	tools.AddNode(head, 1)
	//AddNode(head, 2)
	//AddNode(head, 1)

	fmt.Println(isPalindrome(head.Next))

}

func isPalindrome(head *tools.ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	// 先找中点 这里找的是 左中点
	fast, slow := head.Next, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	h2 := slow.Next
	slow.Next = nil

	// 倒转 next
	var prev *tools.ListNode
	cur := h2

	for cur != nil {
		next := cur.Next

		cur.Next = prev

		prev = cur
		cur = next
	}
	p := head
	for p != nil && prev != nil {

		if p.Val != prev.Val {
			return false
		}
		p = p.Next
		prev = prev.Next
	}

	return true
}
