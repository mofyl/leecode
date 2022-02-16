package v4

import "testing"

func TestIsSymmetric(t *testing.T) {

}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(A *TreeNode, B *TreeNode) bool {

	if A == nil && B == nil {
		return true
	}

	if A == nil || B == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}

	return isSymmetricHelper(A.Left, B.Right) && isSymmetricHelper(A.Right, B.Left)

}
