package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	head = insert(head, 2)
	head1 := removeNthFromEnd(head, 2)
	print(head1)
}

func insert(head *ListNode, ele int) *ListNode {
	newNode := &ListNode{
		Val:  ele,
		Next: nil,
	}
	temp := head
	for ; temp.Next != nil; temp = temp.Next {

	}

	temp.Next = newNode
	return head
}

func print(head *ListNode) {
	temp := head

	for ; temp != nil; temp = temp.Next {
		fmt.Println(temp.Val)
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	/*
		一次遍历
		快慢指针
		分析：
			若我们直接有一个指针fast指向最后的nil的的地方，那么我们只需要删除fast 向前数n的那个节点
			那么我们这里再引入一个指针slow ，我们保证 fast和slow 之间的距离就是n，也就是说fast slow之间隔了n个节点,那么slow->next 就是需要删除的节点
		快指针先走
		总结：
			首先先引入一个虚拟节点,加在头结点的前面(因为若要删除头节点，按照原来的话p已经指向头节点了，没有指向头结点的前一个节点)
			我们先让fast指向该虚拟节点,那么我们移动fast的时候就要移动n+1个，因为我们手动引入了一个新的节点。
			移动完成后，判断 fast时候为nil，若不为nil 则同时移动 fast slow，直到fast为nil,那么slow->next 就是我们要删除的节点

	*/

	dummyHead := new(ListNode)
	dummyHead.Next = head

	fast, slow := dummyHead, dummyHead

	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummyHead.Next
}

func removeNthFromEnd_v1(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}
	lenList := 0
	temp := head
	for ; temp != nil; temp = temp.Next {
		lenList++
	}
	num := lenList - n
	cur := 0
	temp1 := head
	if cur == num {
		head = temp1.Next
		return head
	}
	for ; temp1.Next != nil; temp1 = temp1.Next {
		if cur+1 == num {
			temp2 := temp1.Next
			temp1.Next = temp2.Next
			return head
		}
		cur++
	}
	return nil

}
