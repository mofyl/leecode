package V2

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	//l2 := &TreeNode{Val: 2}
	//l1 := &TreeNode{Val: 1}
	//l3 := &TreeNode{Val: 3}
	//
	//l2.Left = l1
	//l2.Right = l3
	//
	//res := isValidBST(l2)
	//fmt.Println(res)

	l1 := &TreeNode{Val: 1}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l6 := &TreeNode{Val: 6}

	l5.Left = l1
	l5.Right = l4

	l4.Left = l3
	l4.Right = l6

	res := isValidBST(l5)
	fmt.Println(res)

}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	prev := -1
	return isValidBSTHelper(root, &prev)
}

func isValidBSTHelper(root *TreeNode, prev *int) bool {

	if root == nil {
		return true
	}

	if !isValidBSTHelper(root.Left, prev) {
		return false
	}

	if *prev != -1 {
		if root.Val <= *prev {
			return false
		}
	}

	*prev = root.Val

	if !isValidBSTHelper(root.Right, prev) {
		return false
	}
	return true
}
