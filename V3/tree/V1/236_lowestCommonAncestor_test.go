package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestLowestCommonAncestor_236(t *testing.T) {

	l3 := &tree.TreeNode{Val: 3}
	l5 := &tree.TreeNode{Val: 5}
	l1 := &tree.TreeNode{Val: 1}
	l6 := &tree.TreeNode{Val: 6}
	l2 := &tree.TreeNode{Val: 2}
	l7 := &tree.TreeNode{Val: 7}
	l4 := &tree.TreeNode{Val: 4}
	l0 := &tree.TreeNode{Val: 0}
	l8 := &tree.TreeNode{Val: 8}

	l3.Left = l5
	l3.Right = l1

	l5.Left = l6
	l5.Right = l2

	l1.Left = l0
	l1.Right = l8

	l2.Left = l7
	l2.Right = l4

	p := &tree.TreeNode{Val: 5}
	q := &tree.TreeNode{Val: 1}

	node := lowestCommonAncestor_236(l3, p, q)

	if node != nil {

		fmt.Println(node.Val)
	} else {
		fmt.Println("not find")
	}
}

func lowestCommonAncestor_236(root, p, q *tree.TreeNode) *tree.TreeNode {

	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor_236(root.Left, p, q)
	right := lowestCommonAncestor_236(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}
