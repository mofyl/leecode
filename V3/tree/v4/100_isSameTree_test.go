package v4

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	//
	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//
	//l1.Left = l2
	//l1.Right = l3
	//
	//l11 := &TreeNode{Val: 1}
	//l21 := &TreeNode{Val: 2}
	//l31 := &TreeNode{Val: 3}
	//
	//l11.Left = l21
	//l11.Right = l31
	//
	//res := isSameTree(l1, l11)

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}

	l1.Left = l2

	l11 := &TreeNode{Val: 1}
	l21 := &TreeNode{Val: 2}

	l11.Right = l21

	res := isSameTree(l1, l11)

	fmt.Println(res)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
