package tree

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}

	l1.Left = l2
	l1.Right = l3

	l2.Left = l4
	l2.Right = l5

	res := diameterOfBinaryTree(l1)
	fmt.Println(res)
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	findMaxDeep(root, &max)

	return max
}

func findMaxDeep(root *TreeNode, max *int) int {

	if root == nil {
		return 0
	}

	l := findMaxDeep(root.Left, max)
	r := findMaxDeep(root.Right, max)

	*max = tools.Max(l+r, *max)

	return tools.Max(l, r) + 1
}
