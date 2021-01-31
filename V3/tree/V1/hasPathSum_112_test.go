package V1

import (
	"leecode/V3/tree"
	"testing"
)

func TestHasPathSum(t *testing.T) {

}

func hasPathSum(root *tree.TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			return true
		}
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
