package V2

import (
	"leecode/tools"
	"testing"
)

func TestMaxDepth(t *testing.T) {

}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return tools.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}
