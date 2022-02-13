package codeing_interviews

import "testing"

func TestMergeTwoLists(t *testing.T) {

	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l4 := &ListNode{Val: 4}

	l1.Next = l2
	l2.Next = l4

	l11 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 3}
	l41 := &ListNode{Val: 4}

	l11.Next = l3
	l3.Next = l41

	res := mergeTwoLists(l1, l11)

	PrintNode(res)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	p := head
	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
			p = p.Next
		} else if l2.Val < l1.Val {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		} else {
			// 这里表示相等的情况
			p.Next = l1
			p = p.Next
			l1 = l1.Next

			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}

	}

	for l1 != nil {

		p.Next = l1

		p = p.Next
		l1 = l1.Next
	}

	for l2 != nil {
		p.Next = l2

		p = p.Next
		l2 = l2.Next
	}

	return head.Next

}
