package v4

import "testing"

func TestInvertTree(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}

	l2.Left = l1
	l2.Right = l3

	res := invertTree(l2)

	print(res)
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	tmp := root.Left

	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmp)

	return root
}
