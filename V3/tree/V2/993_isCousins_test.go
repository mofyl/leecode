package V2

import (
	"fmt"
	"testing"
)

func TestIsCousins(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}

	//l1.Left = l2
	//l1.Right = l3
	//
	//l2.Left = l4
	//
	//res := isCousins(l1, 4, 3)
	//fmt.Println(res)

	l5 := &TreeNode{Val: 5}

	l1.Left = l2
	l1.Right = l3

	l2.Right = l4
	l3.Right = l5

	res := isCousins(l1, 5, 4)
	fmt.Println(res)
}

var (
	xDeep = 0
	yDeep = 0

	xParent *TreeNode
	yParent *TreeNode
)

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil || x == y {
		return false
	}

	isCousinsLoop(root, nil, x, y, 0)
	return xDeep == yDeep && xParent != yParent
}

func isCousinsLoop(root, parent *TreeNode, x, y int, deep int) {

	if root == nil {
		return
	}

	if root.Val == x {
		xDeep = deep
		xParent = parent
	} else if root.Val == y {
		yDeep = deep
		yParent = parent
	}

	isCousinsLoop(root.Left, root, x, y, deep+1)
	isCousinsLoop(root.Right, root, x, y, deep+1)

}
