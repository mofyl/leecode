package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestIsSameTree(t *testing.T) {

	root1 := &tree.TreeNode{Val: 1}
	l1 := &tree.TreeNode{Val: 2}
	l2 := &tree.TreeNode{Val: 3}

	root2 := &tree.TreeNode{Val: 1}
	l3 := &tree.TreeNode{Val: 2}
	l4 := &tree.TreeNode{Val: 3}

	root1.Left = l1
	root1.Right = l2

	root2.Left = l4
	root2.Right = l3

	res := isSameTree(root1, root2)
	fmt.Println(res)

}

func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
