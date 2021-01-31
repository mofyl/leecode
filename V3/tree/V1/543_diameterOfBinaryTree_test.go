package V1

import (
	"fmt"
	"leecode/V3/tree"
	"leecode/tools"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	l1 := &tree.TreeNode{Val: 1}
	l2 := &tree.TreeNode{Val: 2}
	l3 := &tree.TreeNode{Val: 3}
	l4 := &tree.TreeNode{Val: 4}
	l5 := &tree.TreeNode{Val: 5}

	l1.Left = l2
	l1.Right = l3

	l2.Left = l4
	l2.Right = l5

	res := diameterOfBinaryTree(l1)
	fmt.Println(res)
}

func diameterOfBinaryTree(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	findMaxDeep(root, &max)

	return max
}

func findMaxDeep(root *tree.TreeNode, max *int) int {

	if root == nil {
		return 0
	}

	l := findMaxDeep(root.Left, max)
	r := findMaxDeep(root.Right, max)

	*max = tools.Max(l+r, *max)

	return tools.Max(l, r) + 1
}
