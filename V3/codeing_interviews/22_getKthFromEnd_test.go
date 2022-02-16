package codeing_interviews

import (
	"testing"
)

func TestGetKthFromEnd(t *testing.T) {

	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	k := 2

	res := getKthFromEnd(l1, k)

	PrintNode(res)
}

func getKthFromEnd(head *ListNode, k int) *ListNode {

	p := head
	q := head
	for i := 0; i < k; i++ {
		q = q.Next
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}

	return p

}
