package V1

import (
	"leecode/V3/tree"
	"testing"
)

func TestIsCousins(t *testing.T) {

}

var (
	p1 *tree.TreeNode // 表示 第一个节点的父亲节点
	p2 *tree.TreeNode // 表示 第二个节点的父节点

	d1 int // 表示第一个节点的深度
	d2 int // 表示第二个节点的深度
)

func isCousins(root *tree.TreeNode, x int, y int) bool {

	if root == nil || x == y {
		return false
	}

	isCousins_loop(root, nil, 0, x, y)
	return p1 != p2 && d1 == d2

}

func isCousins_loop(root *tree.TreeNode, parent *tree.TreeNode, k int, x int, y int) {

	if root == nil {
		return
	}

	if root.Val == x {
		p1 = parent
		d1 = k
	} else if root.Val == y {
		p2 = parent
		d2 = k
	}

	isCousins_loop(root.Left, root, k+1, x, y)
	isCousins_loop(root.Right, root, k+1, x, y)
}
