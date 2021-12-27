package V3

import (
	"fmt"
	"math"
	"testing"
)

func TestIsValidBST(t *testing.T) {

	// l1 := &TreeNode{Val: 1}
	// l2 := &TreeNode{Val: 2}
	// l3 := &TreeNode{Val: 3}

	// l2.Left = l1
	// l2.Right = l3

	// l1 := &TreeNode{Val: 1}
	// l4 := &TreeNode{Val: 4}
	// l3 := &TreeNode{Val: 3}
	// l5 := &TreeNode{Val: 5}
	// l6 := &TreeNode{Val: 6}

	// l4.Left = l3
	// l4.Right = l6

	// l5.Left = l1
	// l5.Right = l4

	// l5 := &TreeNode{Val: 5}
	// l4 := &TreeNode{Val: 4}
	// l6 := &TreeNode{Val: 6}
	// l3 := &TreeNode{Val: 3}
	// l7 := &TreeNode{Val: 7}

	// l6.Left = l3
	// l6.Right = l7

	// l5.Left = l4
	// l5.Right = l6

	l1 := &TreeNode{Val: -1}
	l11 := &TreeNode{Val: -1}

	l1.Left = l11

	fmt.Println(isValidBST(l1))
}

func isValidBST(root *TreeNode) bool {
	prev := math.MinInt64
	return isValidBSTHelper(root, &prev)
}

func isValidBSTHelper(root *TreeNode, prev *int) bool {

	if root == nil {
		return true
	}

	if !isValidBSTHelper(root.Left, prev) {
		return false
	}

	if *prev != math.MinInt64 && root.Val <= *prev {
		return false
	}
	*prev = root.Val

	return isValidBSTHelper(root.Right, prev)
}
