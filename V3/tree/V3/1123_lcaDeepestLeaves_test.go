package V3

import (
	"fmt"
	"testing"
)

func TestLcaDeepetLeaves(t *testing.T) {

	l0 := &TreeNode{Val: 0}
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l6 := &TreeNode{Val: 6}
	l7 := &TreeNode{Val: 7}
	l8 := &TreeNode{Val: 8}

	l3.Left = l5
	l3.Right = l1

	l5.Left = l6
	l5.Right = l2

	l2.Left = l7
	l2.Right = l4

	l1.Left = l0
	l1.Right = l8

	v := lcaDeepestLeaves(l3)
	fmt.Println(v)
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	_, v := getDepth(root, nil)
	return v
}

func getDepth(root *TreeNode, parent *TreeNode) (int, *TreeNode) {

	if root == nil {
		return 0, parent
	}

	left, lP := getDepth(root.Left, root)
	right, rP := getDepth(root.Right, root)

	if left > right {
		return left + 1, lP
	} else if left < right {
		return right + 1, rP
	} else {
		return left + 1, root
	}

}
