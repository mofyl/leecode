package tree

import "testing"

func TestInvertTree(t *testing.T) {

}

// 该题只要 前序递归 交换左右子节点就好
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
