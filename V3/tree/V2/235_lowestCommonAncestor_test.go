package V2

import "testing"

func TestLowestCommonAncestor(t *testing.T) {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || p == nil || q == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}

}
