package V2

import (
	"testing"
)

func TestFlatten(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l6 := &TreeNode{Val: 6}

	l1.Left = l2
	l1.Right = l5

	l2.Left = l3
	l2.Right = l4

	l5.Right = l6

	flatten(l1)

}

//
////  这种方式超时
//var (
//	cur_114 *TreeNode
//)
//
//func flatten(root *TreeNode) {
//	//ans := &TreeNode{}
//	//cur_114 = ans
//	cur_114 = root
//	flattenHelper(root)
//	res := inorderTraversal(root)
//
//	fmt.Println(res)
//	//root = ans.Right
//}
//
//func flattenHelper(root *TreeNode) {
//	if root == nil {
//		return
//	}
//
//	right := root.Right
//	left := root.Left
//	cur_114.Right = root
//	root.Left = nil
//	cur_114 = cur_114.Right
//
//	flattenHelper(left)
//	flattenHelper(right)
//}

func flatten(root *TreeNode) {

	if root == nil {
		return
	}

	flatten(root.Left)

	right := root.Right

	root.Right = root.Left
	root.Left = nil

	for root.Right != nil {
		root = root.Right
	}

	flatten(right)

	root.Right = right

}
