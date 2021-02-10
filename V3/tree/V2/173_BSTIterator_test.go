package V2

import (
	"fmt"
	"testing"
)

func TestBSTIterator(t *testing.T) {

	l3 := &TreeNode{Val: 3}
	l7 := &TreeNode{Val: 7}
	l9 := &TreeNode{Val: 9}
	l15 := &TreeNode{Val: 15}
	l20 := &TreeNode{Val: 20}

	l7.Left = l3
	l7.Right = l15

	l15.Left = l9
	l15.Right = l20
	iterator := Constructor(l7)
	fmt.Println(iterator.Next())    // 返回 3
	fmt.Println(iterator.Next())    // 返回 7
	fmt.Println(iterator.HasNext()) // 返回 true
	fmt.Println(iterator.Next())    // 返回 9
	fmt.Println(iterator.HasNext()) // 返回 true
	fmt.Println(iterator.Next())    // 返回 15
	fmt.Println(iterator.HasNext()) // 返回 true
	fmt.Println(iterator.Next())    // 返回 20
	fmt.Println(iterator.HasNext()) // 返回 false

}

type BSTIterator struct {
	data []int
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{}
	if root == nil {
		return iter
	}

	res := inorderTraversal(root)
	iter.data = res
	return iter
}

func (this *BSTIterator) Next() int {
	if len(this.data) < 1 {
		return 0
	}

	v := this.data[0]
	this.data = this.data[1:]
	return v
}

func (this *BSTIterator) HasNext() bool {
	return len(this.data) > 0
}
