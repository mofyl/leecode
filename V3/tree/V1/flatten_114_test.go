package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestFlatten(t *testing.T) {
	l1 := &tree.TreeNode{Val: 1}
	l2 := &tree.TreeNode{Val: 2}
	l5 := &tree.TreeNode{Val: 5}
	l3 := &tree.TreeNode{Val: 3}
	l4 := &tree.TreeNode{Val: 4}
	l6 := &tree.TreeNode{Val: 6}

	l1.Left = l2
	l1.Right = l5

	l2.Left = l3
	l2.Right = l4

	l5.Right = l6

	flatten(l1)

	res := inorderTraversal(l1)

	fmt.Println(res)
}

// 这里采用 中序遍历
func flatten(root *tree.TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)

	right := root.Right

	root.Right = root.Left
	root.Left = nil

	for root.Right != nil {
		// 这里一直到 当前最后一个 右子树节点
		root = root.Right
	}
	// 对 当前右子树进行递归。
	flatten(right)
	// 递归完成后 将 结果保存到最后
	root.Right = right
}
