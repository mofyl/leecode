package tree

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l21 := &TreeNode{Val: 2}
	l22 := &TreeNode{Val: 2}
	l31 := &TreeNode{Val: 3}
	l32 := &TreeNode{Val: 3}
	l41 := &TreeNode{Val: 4}
	l42 := &TreeNode{Val: 4}

	l1.Left = l21
	l1.Right = l22

	l21.Left = l31
	l21.Right = l41

	l22.Left = l42
	l22.Left = l32

	//l1 := &TreeNode{Val: 1}
	//l21 := &TreeNode{Val: 2}
	//l22 := &TreeNode{Val: 2}
	//l31 := &TreeNode{Val: 3}
	//l32 := &TreeNode{Val: 3}
	//
	//l1.Left = l21
	//l1.Right = l22
	//
	//l21.Right = l31
	//l22.Right = l32
	fmt.Println(isSymmetric(l1))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetric_loop(root.Left, root.Right)
}

func isSymmetric_loop(left *TreeNode, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if (left == nil && right != nil) || (right == nil && left != nil) ||
		(left.Val != right.Val) {
		return false
	}

	return isSymmetric_loop(left.Left, right.Right) && isSymmetric_loop(left.Right, right.Left)
}
