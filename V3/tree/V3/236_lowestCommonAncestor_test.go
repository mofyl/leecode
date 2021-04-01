package V3

import "testing"

func TestLowestCommonAncestor_236(t *testing.T) {

}

func lowestCommonAncestor_236(root, p, q *TreeNode) *TreeNode {

	if root == nil || p == nil || q == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor_236(root.Left, p, q)
	right := lowestCommonAncestor_236(root.Left, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}

}
