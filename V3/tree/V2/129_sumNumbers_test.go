package V2

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//
	//l1.Left = l2
	//l1.Right = l3

	l0 := &TreeNode{Val: 0}
	l1 := &TreeNode{Val: 1}
	l5 := &TreeNode{Val: 5}
	l9 := &TreeNode{Val: 9}
	l4 := &TreeNode{Val: 4}

	l4.Left = l9
	l4.Right = l0

	l9.Left = l5
	l9.Right = l1

	sumNumbers(l4)
}

func sumNumbers(root *TreeNode) int {

	if root == nil {
		return 0
	}
	sum := 0
	sumNumbersHelper(root, 0, &sum)

	return sum
}

func sumNumbersHelper(root *TreeNode, cur int, sum *int) {
	if root == nil {
		return
	}

	cur = cur*10 + root.Val

	if root.Left == nil && root.Right == nil {
		*sum += cur
	}

	sumNumbersHelper(root.Left, cur, sum)
	sumNumbersHelper(root.Right, cur, sum)
}
