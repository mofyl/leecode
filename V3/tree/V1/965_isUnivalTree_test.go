package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestIsUnivalTree(t *testing.T) {
	l1 := &tree.TreeNode{Val: 1}
	l2 := &tree.TreeNode{Val: 1}
	l3 := &tree.TreeNode{Val: 1}
	l4 := &tree.TreeNode{Val: 2}
	l5 := &tree.TreeNode{Val: 1}

	l1.Left = l2
	l1.Right = l3

	l2.Left = l4
	l2.Right = l5

	res := isUnivalTree(l1)
	fmt.Println(res)
}

func isUnivalTree(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return isUnivalTree_loop(root, root.Val)
}

func isUnivalTree_loop(root *tree.TreeNode, val int) bool {
	if root == nil {
		return true
	}

	left := isUnivalTree_loop(root.Left, val)

	if !left {
		return false
	}

	if root.Val != val {
		return false
	}

	right := isUnivalTree_loop(root.Right, val)

	if !right {
		return false
	}
	return true
}
