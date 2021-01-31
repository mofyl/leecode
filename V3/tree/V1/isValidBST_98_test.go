package V1

import (
	"fmt"
	"leecode/V3/tree"
	"math"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	//
	//root := &TreeNode{
	//	Val: 2,
	//}
	////
	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 4}
	//
	//root.Left = l1
	//root.Right = l2
	//
	//root := &TreeNode{
	//	Val: 5,
	//}
	//
	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 4}
	//l3 := &TreeNode{Val: 3}
	//l4 := &TreeNode{Val: 6}
	//
	//l2.Left = l3
	//l2.Right = l4
	//
	//root.Left = l1
	//root.Right = l2

	//root := &TreeNode{
	//	Val: 2,
	//}
	//l2 := &TreeNode{Val: 1}
	//l3 := &TreeNode{Val: 3}
	//
	//l2.Right = l3
	//root.Left = l2

	root := &tree.TreeNode{Val: 5}

	l1 := &tree.TreeNode{Val: 1}
	l2 := &tree.TreeNode{Val: 4}
	l3 := &tree.TreeNode{Val: 3}
	l4 := &tree.TreeNode{Val: 6}

	l2.Left = l3
	l2.Right = l4

	root.Left = l1
	root.Right = l2

	//fmt.Println(isValidBST(root))
	fmt.Println(isValidBST_V2(root))

}

func isValidBST(root *tree.TreeNode) bool {
	//res := 0
	//inorderTrabversal_98(root, &res, &[]int{})
	//
	//if res == 1 {
	//	return false
	//} else {
	//	return true
	//}

	res := inorderTraversal(root)

	for i := 1; i < len(res); i++ {
		// 注意要小于等于，搜索树里不能有相同元素
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

func isValidBST_V2(root *tree.TreeNode) bool {

	// 空树也是一颗二叉树
	if root == nil {
		return true
	}
	// 只有一个节点 也是一颗二叉树
	if root.Left == nil && root.Right == nil {
		return true
	}
	num := math.MinInt64
	return isVa(root, &num)
}

// prev 的初始值 是 min_int64 . prev 表示 上一个遍历过的节点
// 由于是 中序 是先遍历左子树，那么当前节点的值 就要 > prev
func isVa(root *tree.TreeNode, prev *int) bool {

	// 空树也是一颗二叉树
	if root == nil {
		return true
	}

	if !isVa(root.Left, prev) {
		return false
	}

	if root.Val <= *prev {
		return false
	}

	*prev = root.Val

	if !isVa(root.Right, prev) {
		return false
	}

	return true
}
