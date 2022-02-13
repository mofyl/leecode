package codeing_interviews

import (
	"testing"
)

func TestBuildTree(t *testing.T) {

	//preorder := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}

	preorder := []int{-1}
	inorder := []int{-1}

	root := buildTree(preorder, inorder)

	print(root)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	preIdx := 0
	inIdx := 0

	return buildTreeHelper(preorder, inorder, &preIdx, &inIdx, nil)
}

func buildTreeHelper(preorder []int, inorder []int, preIdx *int, inIdx *int, pre *int) *TreeNode {

	if *preIdx >= len(preorder) ||
		*inIdx >= len(inorder) {
		return nil
	}

	if pre != nil && *pre == inorder[*inIdx] {
		*inIdx++

		return nil
	}

	node := &TreeNode{Val: preorder[*preIdx]}

	*preIdx++

	node.Left = buildTreeHelper(preorder, inorder, preIdx, inIdx, &node.Val)
	node.Right = buildTreeHelper(preorder, inorder, preIdx, inIdx, pre)

	return node
}
