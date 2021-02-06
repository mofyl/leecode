package V2

import (
	"fmt"
	"math"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	node := buildTree(preorder, inorder)

	res := inorderTraversal(node)
	fmt.Println(res)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) < 1 {
		return nil
	}
	pre := 0
	in := 0
	return buildTreeHelper(preorder, inorder, &pre, &in, math.MaxInt64)
}

func buildTreeHelper(preorder []int, inorder []int, pre *int, in *int, stop int) *TreeNode {
	if *pre == len(preorder) {
		return nil
	}

	if inorder[*in] == stop {
		*in++
		return nil
	}

	val := preorder[*pre]
	*pre++
	node := &TreeNode{Val: val}

	node.Left = buildTreeHelper(preorder, inorder, pre, in, val)
	node.Right = buildTreeHelper(preorder, inorder, pre, in, stop)

	return node
}
