package V2

import "testing"

func TestIsUnivalTree(t *testing.T) {

}

func isUnivalTree(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isUnivalTreeLoop(root, root.Val)
}

func isUnivalTreeLoop(root *TreeNode, v int) bool {

	if root == nil {
		return true
	}

	if root.Val != v {
		return false
	}

	return isUnivalTreeLoop(root.Left, v) && isUnivalTreeLoop(root.Right, v)

}
