package V2

import (
	"leecode/tools"
	"math"
	"testing"
)

func TestMaxPathSum(t *testing.T) {

}

// 这题 不知道什么意思
func maxPathSum(root *TreeNode) int {

	if root == nil {
		return 0
	}
	maxVal := math.MinInt64
	res := maxPathSumHelper(root, &maxVal)
	return tools.Max(maxVal, res)
}

func maxPathSumHelper(root *TreeNode, maxVal *int) int {

	if root == nil {
		return 0
	}

	left := tools.Max(0, maxPathSumHelper(root.Left, maxVal))
	right := tools.Max(0, maxPathSumHelper(root.Right, maxVal))

	*maxVal = tools.Max(*maxVal, root.Val+left+right)
	// 这里不是很明白 为什么本节点的 产值 是  val + max
	return root.Val + tools.Max(left, right)
}
