package V1

import (
	"fmt"
	"leecode/V3/tree"
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

	l1 := &tree.TreeNode{Val: 1}
	l2 := &tree.TreeNode{Val: 2}
	l22 := &tree.TreeNode{Val: 2}
	l3 := &tree.TreeNode{Val: 3}
	l33 := &tree.TreeNode{Val: 3}
	l4 := &tree.TreeNode{Val: 4}
	l44 := &tree.TreeNode{Val: 4}

	l1.Left = l2
	l1.Right = l22

	l2.Left = l3
	l2.Right = l33

	l3.Left = l4
	l3.Right = l44

	res := isBalanced(l1)
	fmt.Println(res)

}

func isBalanced(root *tree.TreeNode) bool {

	if root == nil {
		return true
	}

	if tools.Abs(getDeep(root.Left)-getDeep(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDeep(root *tree.TreeNode) int {

	if root == nil {
		return 0
	}

	return tools.Max(getDeep(root.Left), getDeep(root.Right)) + 1
}
