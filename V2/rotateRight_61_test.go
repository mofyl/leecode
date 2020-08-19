package V2

import (
	"fmt"
	"testing"
)

func TestRotateList(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	head := rotateRight(l1, 2)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

/*

	这里先去求出链表的长度size
	若k是size的整数倍 那么旋转过后还是原来的链表
	真正需要移动的次数为 k%size
	旋转：就是找一个位置pos，将链表从pos位置断开。然后将pos后面的链表接到head前面，
	移动k个位置代表将最后k个数字移动到head前面，那么开始断开的位置pos=size-k%size





*/

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	size := 1
	p := head

	for p.Next != nil {
		p = p.Next
		size++
	}
	move := k % size

	if move == 0 {
		return head
	}

	cur := head
	// 这里表示先将cur移动到需要断开的位置
	// 这里由于i=0 所以需要-1
	for i := 0; i < size-move-1; i++ {
		cur = cur.Next
	}

	res := cur.Next
	cur.Next = nil
	p.Next = head
	return res
}
