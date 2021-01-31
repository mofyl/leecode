package V1

import (
	"leecode/V3/tree"
	"leecode/tools"
	"testing"
)

func TestOfferMaxDepth(t *testing.T) {

}

func offerMaxDepth(root *tree.TreeNode) int {

	if root == nil {
		return 0
	}

	left := offerMaxDepth(root.Left) + 1
	right := offerMaxDepth(root.Right) + 1

	return tools.Max(left, right)
}
