package V2

import (
	"testing"
)

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
	right := lowestCommonAncestor_236(root.Right, p, q)

	if left == nil {
		// 这里表示 左子树没有找到，那么就返回上一次层 继续找
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}
