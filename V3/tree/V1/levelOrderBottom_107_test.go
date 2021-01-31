package V1

import (
	"leecode/V3/tree"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {

}

func levelOrderBottom(root *tree.TreeNode) [][]int {

	res := levelOrder(root)

	resNew := make([][]int, 0, len(res))

	for i := len(res) - 1; i >= 0; i-- {
		resNew = append(resNew, res[i])
	}
	return resNew
}
