package list

import (
	"fmt"
	"testing"
)

func TestMyLinkedList(t *testing.T) {

	head := Constructor()
	head.AddAtHead(1)
	head.AddAtTail(3)
	head.AddAtIndex(1, 2)

	head.PrintAllNode()
	fmt.Println("=================")
	val := head.Get(1)
	fmt.Println(val)
	fmt.Println("=================")
	head.DeleteAtIndex(1)
	//
	val = head.Get(1)
	fmt.Println(val)
}

type MyLinkedList struct {
	Val    int
	Next   *MyLinkedList
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Val:    0,
		Next:   nil,
		length: 0,
	}
}

// TODO  这里index 从0 开始！！！！
/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.length {
		return -1
	}

	count := 0
	p := this.Next

	for p != nil {
		if count == index {
			return p.Val
		}
		p = p.Next
		count++
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	tmp := &MyLinkedList{
		Val: val,
	}

	next := this.Next
	tmp.Next = next
	this.Next = tmp
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {

	p := this

	for p.Next != nil {
		p = p.Next
	}

	tmp := &MyLinkedList{
		Val:  val,
		Next: nil,
	}

	p.Next = tmp
	this.length++

}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length {
		return
	}

	if index < 0 {
		this.AddAtHead(val)
		return
	}

	p := this
	count := 0

	for p != nil {

		if count == index {
			tmp := &MyLinkedList{
				Val: val,
			}
			next := p.Next
			p.Next = tmp
			tmp.Next = next
			this.length++
			return
		}
		count++
		p = p.Next
	}

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.length {
		return
	}

	p := this
	count := 0

	for p.Next != nil {
		if count == index {
			p.Next = p.Next.Next
			this.length--
			return
		}
		count++
		p = p.Next
	}
}

func (this *MyLinkedList) PrintAllNode() {
	if this.length < 1 {
		return
	}

	fmt.Println("length is ", this.length)

	p := this.Next

	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
