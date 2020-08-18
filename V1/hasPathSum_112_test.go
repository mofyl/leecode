package V1

import (
	"fmt"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	n := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	res := hasPathSum(n, 1)
	fmt.Println(res)
}

func hasPathSum(root *TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	// 这里只有 当前节点的左右儿子都为nil 该节点才是叶子节点
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) ||
		hasPathSum(root.Right, sum-root.Val)
}
