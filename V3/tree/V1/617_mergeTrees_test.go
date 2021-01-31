package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	l1 := &tree.TreeNode{Val: 1}
	l3 := &tree.TreeNode{Val: 3}
	l2 := &tree.TreeNode{Val: 2}
	l5 := &tree.TreeNode{Val: 5}

	l1.Left = l3
	l1.Right = l2
	l3.Left = l5

	l22 := &tree.TreeNode{Val: 2}
	l11 := &tree.TreeNode{Val: 1}
	l33 := &tree.TreeNode{Val: 3}
	l44 := &tree.TreeNode{Val: 4}
	l7 := &tree.TreeNode{Val: 7}

	l22.Left = l11
	l22.Right = l33
	l11.Right = l44
	l33.Right = l7

	newNode := mergeTrees(l1, l22)

	res := inorderTraversal(newNode)
	fmt.Println(res)
	//inorderTraversal(newNode)
}

func mergeTrees(t1 *tree.TreeNode, t2 *tree.TreeNode) *tree.TreeNode {

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	node := &tree.TreeNode{Val: t1.Val + t2.Val}

	node.Left = mergeTrees(t1.Left, t2.Left)
	node.Right = mergeTrees(t1.Right, t2.Right)

	return node
}
