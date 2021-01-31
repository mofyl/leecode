package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestSumRootToLeaf(t *testing.T) {
	////
	//l11 := &TreeNode{Val: 1}
	//l12 := &TreeNode{Val: 1}
	//l13 := &TreeNode{Val: 1}
	//l14 := &TreeNode{Val: 1}
	//l01 := &TreeNode{Val: 0}
	//l02 := &TreeNode{Val: 0}
	//l03 := &TreeNode{Val: 0}
	//
	//l11.Left = l01
	//l11.Right = l12
	//
	//l01.Left = l02
	//l01.Right = l13
	//
	//l12.Left = l03
	//l12.Right = l14
	//
	//res := sumRootToLeaf(l11)

	l11 := &tree.TreeNode{Val: 1}
	l12 := &tree.TreeNode{Val: 1}
	l11.Left = l12

	res := sumRootToLeaf(l11)
	fmt.Println(res)
}

func sumRootToLeaf(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	curNum := 0
	sum := 0
	sumRootToLeaf_loop(root, &sum, curNum)

	return sum
}

func sumRootToLeaf_loop(root *tree.TreeNode, res *int, cur int) {

	if root == nil {
		return
	}

	cur = cur<<1 | root.Val

	if root.Left == nil && root.Right == nil {
		*res += cur
	} else {

		if root.Left != nil {
			sumRootToLeaf_loop(root.Left, res, cur)
		}

		if root.Right != nil {
			sumRootToLeaf_loop(root.Right, res, cur)
		}
	}

}
