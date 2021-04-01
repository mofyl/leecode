package V3

import "testing"

func TestLowestCommonAncestor(t *testing.T) {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || p == nil || q == nil {
		return nil
	}
	// 这里是一种优化 若p,q不在同一深度，先找到了高的那个，那么就可以返回了
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
