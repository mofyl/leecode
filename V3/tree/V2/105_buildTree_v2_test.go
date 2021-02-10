package V2

import (
	"math"
	"testing"
)

func TestBuildTreeV2(t *testing.T) {

}
func buildTreeV2(preorder []int, inorder []int) *TreeNode {
	pre := 0
	in := 0

	return buildTreeV2Helper(preorder, inorder, &pre, &in, math.MaxInt64)
}

func buildTreeV2Helper(preorder []int, inorder []int, pre *int, in *int, stop int) *TreeNode {

	if *pre >= len(preorder) {
		return nil
	}

	if *in == stop {
		*in++
		return nil
	}

	val := preorder[*pre]
	*pre++
	node := &TreeNode{Val: val}

	node.Left = buildTreeV2Helper(preorder, inorder, pre, in, val)
	node.Right = buildTreeV2Helper(preorder, inorder, pre, in, stop)

	return node
}
