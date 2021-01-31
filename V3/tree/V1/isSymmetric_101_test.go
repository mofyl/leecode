package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestIsSymmetric_101(t *testing.T) {
	//
	//root := &TreeNode{Val: 1}
	//l1 := &TreeNode{Val: 2}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//l4 := &TreeNode{Val: 3}
	//l5 := &TreeNode{Val: 4}
	//l6 := &TreeNode{Val: 4}
	//
	//root.Left = l1
	//root.Right = l2
	//
	//l1.Left = l3
	//l1.Right = l5
	//
	//l2.Left = l6
	//l2.Right = l4
	//
	//res := isSymmetric(root)
	//
	//fmt.Println(res)

	root := &tree.TreeNode{Val: 1}
	l2 := &tree.TreeNode{Val: 2}
	l3 := &tree.TreeNode{Val: 2}
	l4 := &tree.TreeNode{Val: 3}
	l5 := &tree.TreeNode{Val: 3}

	root.Left = l2
	root.Right = l3

	l2.Left = l4
	l3.Right = l5

	res := isSymmetric_101(root)

	fmt.Println(res)

}

func isSymmetric_101(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	return isSymmetricLoop(root.Left, root.Right)
}

func isSymmetricLoop(p *tree.TreeNode, q *tree.TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !isSymmetricLoop(p.Left, q.Right) {
		return false
	}

	if !isSymmetricLoop(p.Right, q.Left) {
		return false
	}
	return true
}
