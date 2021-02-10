package V2

import (
	"fmt"
	"math"
	"testing"
)

func TestBuildTree_106(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	node := buildTree_106(inorder, postorder)

	res := inorderTraversal(node)

	fmt.Println(res)
}

func buildTree_106(inorder []int, postorder []int) *TreeNode {

	post := len(postorder) - 1
	in := len(inorder) - 1

	return buildTree_106_Helper(inorder, postorder, &in, &post, math.MaxInt64)
}

func buildTree_106_Helper(inorder, postorder []int, in, post *int, stop int) *TreeNode {

	if *post < 0 {
		return nil
	}

	if inorder[*in] == stop {
		*in--
		return nil
	}

	val := postorder[*post]
	*post--

	node := &TreeNode{Val: val}

	node.Right = buildTree_106_Helper(inorder, postorder, in, post, val)
	node.Left = buildTree_106_Helper(inorder, postorder, in, post, stop)

	return node
}
