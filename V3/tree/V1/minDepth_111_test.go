package V1

import (
	"leecode/V3/tree"
	"leecode/tools"
	"testing"
)

func TestMinDepth(t *testing.T) {

}

func minDepth(root *tree.TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDeep := minDepth(root.Left)

	rightDeep := minDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		return leftDeep + rightDeep + 1
	}

	return tools.Min(leftDeep, rightDeep) + 1
}
