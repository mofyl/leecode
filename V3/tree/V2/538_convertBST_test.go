package V2

import (
	"fmt"
	"testing"
)

func TestConvertBST(t *testing.T) {

	l0 := &TreeNode{Val: 0}
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l6 := &TreeNode{Val: 6}
	l7 := &TreeNode{Val: 7}
	l8 := &TreeNode{Val: 8}

	l4.Left = l1
	l4.Right = l6

	l1.Left = l0
	l1.Right = l2

	l2.Right = l3

	l6.Left = l5
	l6.Right = l7

	l7.Right = l8

	node := convertBST(l4)

	res := inorderTraversal(node)
	fmt.Println(res)
}

func convertBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	prev := 0
	convertBSTHelper(root, &prev)
	return root
}

func convertBSTHelper(root *TreeNode, prev *int) {

	if root == nil {
		return
	}

	convertBSTHelper(root.Right, prev)

	root.Val = root.Val + *prev

	*prev = root.Val

	convertBSTHelper(root.Left, prev)

}
