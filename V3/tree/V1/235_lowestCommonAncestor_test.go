package V1

import (
	"leecode/V3/tree"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {

}

func lowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {

	if root == nil || p == nil || q == nil {
		return nil
	}

	if p.Val > root.Val && q.Val > root.Val {
		// 说明这两个节点在右侧
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		// 说明这两个节点在左侧
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		// 说明 两个节点 一个在左边 一个在右边 那么根节点就是公共祖先
		return root
	}

}
