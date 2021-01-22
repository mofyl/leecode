package tree

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	// true
	//l3 := &TreeNode{Val: 3}
	//l20 := &TreeNode{Val: 20}
	//l9 := &TreeNode{Val: 9}
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

	if tools.Abs(getDeep(root.Left)-getDeep(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDeep(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return tools.Max(getDeep(root.Left), getDeep(root.Right)) + 1
}
