package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestIncreasingBST(t *testing.T) {
	l1 := &tree.TreeNode{Val: 1}
	l2 := &tree.TreeNode{Val: 2}
	l3 := &tree.TreeNode{Val: 3}
	l4 := &tree.TreeNode{Val: 4}
	l5 := &tree.TreeNode{Val: 5}
	l6 := &tree.TreeNode{Val: 6}
	l7 := &tree.TreeNode{Val: 7}
	l8 := &tree.TreeNode{Val: 8}
	l9 := &tree.TreeNode{Val: 9}

	l5.Left = l3
	l5.Right = l6

	l3.Left = l2
	l3.Right = l4

	l2.Left = l1

	l6.Right = l8

	l8.Left = l7
	l8.Right = l9

	res1 := increasingBST(l5)

	d := inorderTraversal(res1)
	fmt.Println(d)
}

var (
	cur *tree.TreeNode
)

func increasingBST(root *tree.TreeNode) *tree.TreeNode {

	if root == nil {
		return nil
	}

	ans := &tree.TreeNode{}

	cur = ans

	increasingBST_loop(root)
	return ans.Right
}

func increasingBST_loop(root *tree.TreeNode) {

	if root == nil {
		return
	}

	increasingBST_loop(root.Left)
	// 这里如果没有的话 就会出现回环,
	// 因为是 先遍历左边，我们这里就会把左边加入到ans上，返回后 就到了当前子树的根节点，
	// 然后我们又将 该跟节点 链接上 那么就出现了环
	root.Left = nil
	cur.Right = root
	cur = root

	increasingBST_loop(root.Right)
}

//
//var (
//	res *TreeNode
//	p   *TreeNode
//)
//
//func increasingBST(root *TreeNode) *TreeNode {
//
//	if root == nil {
//		return nil
//	}
//
//	res = &TreeNode{}
//	p = res
//
//	increasingBST_loop(root)
//	return res.Right
//}
//
//
//func increasingBST_loop(root *TreeNode) {
//
//	if root == nil {
//		return
//	}
//
//	increasingBST_loop(root.Left)
//
//	tmp := &TreeNode{
//		Val: root.Val,
//	}
//
//	p.Right = tmp
//	p = p.Right
//
//	increasingBST_loop(root.Right)
//}
