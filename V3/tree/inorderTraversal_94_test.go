package tree

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}

	//l1.Right = l2
	//l2.Left = l1
	//
	l3.Left = l1
	l1.Right = l2
	res := inorderTraversal(l3)

	fmt.Println(res)
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inorderTraversalLoop(root, &res)
	return res
}

func inorderTraversalLoop(root *TreeNode, res *[]int) {

	if root == nil {
		return
	}

	inorderTraversalLoop(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalLoop(root.Right, res)
}
