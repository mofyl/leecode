/*
 * @Author: lirui38
 * @Date: 2021-12-20 14:59:18
 * @LastEditTime: 2021-12-20 15:33:33
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/173_BSTIterator_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestBSTIterator(t *testing.T) {

	l7 := &TreeNode{Val: 7}
	l3 := &TreeNode{Val: 3}
	l15 := &TreeNode{Val: 15}
	l9 := &TreeNode{Val: 9}
	l20 := &TreeNode{Val: 20}

	l15.Left = l9
	l15.Right = l20

	l7.Left = l3
	l7.Right = l15

	iter := Constructor(l7)

	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())

	// Constructor{}

}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {

	stack := make([]*TreeNode, 0)
	node := root
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}

	return BSTIterator{
		stack: stack,
	}

}

func (this *BSTIterator) Next() int {

	idx := len(this.stack) - 1
	node := this.stack[idx]
	res := this.stack[idx]

	node = node.Right

	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}

	this.stack = append(this.stack[:idx], this.stack[idx+1:]...)
	return res.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
