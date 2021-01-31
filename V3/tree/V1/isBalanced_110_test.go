package V1

import (
	"leecode/V3/tree"
	"leecode/tools"
	"testing"
)

func TestIsBalanced_110(t *testing.T) {

}

func isBalanced_110(root *tree.TreeNode) bool {

	if root == nil {
		return true
	}

	return isBalancedLoop(root) != -1

}

func isBalancedLoop(root *tree.TreeNode) int {

	if root == nil {
		return 0
	}

	leftDeep := isBalancedLoop(root.Left)
	if leftDeep == -1 {
		return -1
	}
	rightDeep := isBalancedLoop(root.Right)
	if rightDeep == -1 {
		return -1
	}

	abs := tools.Abs(leftDeep - rightDeep)
	if abs > 1 {
		return -1
	} else {
		return tools.Max(leftDeep, rightDeep) + 1
	}

}
