package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	l1 := &tree.TreeNode{Val: 1}
	l2 := &tree.TreeNode{Val: 2}
	l3 := &tree.TreeNode{Val: 3}

	//l1.Right = l2
	//l2.Left = l1
	//
	l3.Left = l1
	l1.Right = l2
	res := inorderTraversal(l3)

	fmt.Println(res)
}

func inorderTraversal(root *tree.TreeNode) []int {
	res := make([]int, 0)
	inorderTraversalLoop(root, &res)
	return res
}

func inorderTraversalLoop(root *tree.TreeNode, res *[]int) {

	if root == nil {
		return
	}

	inorderTraversalLoop(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalLoop(root.Right, res)
}
