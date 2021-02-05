package V2

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	//
	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l22 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//l33 := &TreeNode{Val: 3}
	//l4 := &TreeNode{Val: 4}
	//l44 := &TreeNode{Val: 4}
	//
	//l1.Left = l2
	//l1.Right = l22
	//
	//l2.Left = l3
	//l2.Right = l4
	//
	//l22.Left = l44
	//l22.Right = l33

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l22 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l33 := &TreeNode{Val: 3}

	l1.Left = l2
	l1.Right = l22

	l2.Right = l3
	l22.Right = l33

	fmt.Println(isSymmetric(l1))

}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isSymmetricLoop(root.Left, root.Right)

}

func isSymmetricLoop(left *TreeNode, right *TreeNode) bool {

	if left == nil && right != nil {
		return false
	}

	if left != nil && right == nil {
		return false
	}

	if left == nil && right == nil {
		return true
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricLoop(left.Left, right.Right) &&
		isSymmetricLoop(left.Right, right.Left)

}
