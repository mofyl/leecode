package V2

import (
	"fmt"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	l5 := &TreeNode{Val: 5}
	l4 := &TreeNode{Val: 4}
	l8 := &TreeNode{Val: 8}
	l11 := &TreeNode{Val: 11}
	l7 := &TreeNode{Val: 7}
	l2 := &TreeNode{Val: 2}
	l13 := &TreeNode{Val: 13}
	l44 := &TreeNode{Val: 4}
	l1 := &TreeNode{Val: 1}

	l5.Left = l4
	l5.Right = l8

	l4.Left = l11

	l11.Left = l7
	l11.Right = l2

	l8.Left = l13
	l8.Right = l44

	l4.Right = l1

	fmt.Println(hasPathSum(l5, 22))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
