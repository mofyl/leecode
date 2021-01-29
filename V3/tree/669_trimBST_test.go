package tree

import (
	"fmt"
	"testing"
)

func TestTrimBST(t *testing.T) {
	//
	//l1 := &TreeNode{Val: 1}
	//l0 := &TreeNode{Val: 0}
	//l2 := &TreeNode{Val: 2}
	//
	//l1.Left = l0
	//l1.Right = l2
	//
	//node := trimBST(l1, 1, 2)
	//
	//l0 := &TreeNode{Val: 0}
	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//l4 := &TreeNode{Val: 4}
	//
	//l3.Left = l0
	//l3.Right = l4
	//
	//l0.Right = l2
	//
	//l2.Left = l1
	//
	//node := trimBST(l3, 1, 3)

	//l1 := &TreeNode{Val: 1}
	//
	//node := trimBST(l1, 1, 2)
	//
	//
	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//
	//l1.Right = l2
	//
	//node := trimBST(l1, 2, 4)

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}

	l3.Left = l1
	l3.Right = l4

	l1.Right = l2

	node := trimBST(l3, 1, 2)

	res := inorderTraversal(node)
	fmt.Println(res)
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {

	if root == nil {
		return root
	}

	if root.Val < low {

		return trimBST(root.Right, low, high)
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
