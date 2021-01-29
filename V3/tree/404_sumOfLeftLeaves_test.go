package tree

import (
	"fmt"
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {
	l1 := &TreeNode{Val: 3}
	l9 := &TreeNode{Val: 9}
	l20 := &TreeNode{Val: 20}
	l15 := &TreeNode{Val: 15}
	l7 := &TreeNode{Val: 7}

	l1.Left = l9
	l1.Right = l20

	l20.Left = l15
	l20.Right = l7

	sum := sumOfLeftLeaves(l1)
	fmt.Println(sum)
}

func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {
		return 0
	}
	sum := 0
	sumOfLeftLeaves_loop(root, &sum)
	return sum
}

func sumOfLeftLeaves_loop(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*sum += root.Left.Val
	}
	sumOfLeftLeaves_loop(root.Left, sum)
	sumOfLeftLeaves_loop(root.Right, sum)
}
