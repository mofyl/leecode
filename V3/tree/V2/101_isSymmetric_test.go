package V2

import "testing"

func TestIsSymmetric_101(t *testing.T) {

}

func isSymmetric_101(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return isSymmetric_101Helper(root.Left, root.Right)
}

func isSymmetric_101Helper(left *TreeNode, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetric_101Helper(left.Left, right.Right) && isSymmetric_101Helper(left.Right, right.Left)
}
