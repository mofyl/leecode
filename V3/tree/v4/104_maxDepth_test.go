package v4

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxDepth(t *testing.T) {

	l3 := &TreeNode{Val: 3}
	l7 := &TreeNode{Val: 7}
	l9 := &TreeNode{Val: 9}
	l15 := &TreeNode{Val: 15}
	l20 := &TreeNode{Val: 20}

	l3.Left = l9
	l3.Right = l20

	l20.Left = l15
	l20.Right = l7

	h := maxDepth(l3)

	fmt.Println(h)
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1

	return tools.Max(l, r)

}
