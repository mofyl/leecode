package V2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	//l3 := &TreeNode{Val: 3}
	//l9 := &TreeNode{Val: 9}
	//l20 := &TreeNode{Val: 20}
	//l15 := &TreeNode{Val: 15}
	//l7 := &TreeNode{Val: 7}
	//
	//l3.Left = l9
	//l3.Right = l20
	//
	//l20.Left = l15
	//l20.Right = l7
	//
	//res := isBalanced(l3)
	//
	//fmt.Println(res)

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l22 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l33 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l44 := &TreeNode{Val: 4}

	l1.Left = l2
	l1.Right = l22

	l2.Left = l3
	l2.Right = l33

	l3.Left = l4
	l3.Right = l44

	res := isBalanced(l1)

	fmt.Println(res)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalancedLoop(root) != -1
}

func isBalancedLoop(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := isBalancedLoop(root.Left)
	if left == -1 {
		return -1
	}
	right := isBalancedLoop(root.Right)
	if right == -1 {
		return right
	}

	if tools.Abs(left-right) > 1 {
		return -1
	}

	return tools.Max(left, right) + 1

}
