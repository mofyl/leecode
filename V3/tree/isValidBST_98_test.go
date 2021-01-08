package tree

import (
	"fmt"
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

	root := &TreeNode{Val: 5}

	l1 := &TreeNode{Val: 4}
	l2 := &TreeNode{Val: 6}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 7}

	l2.Left = l3
	l2.Right = l4

	root.Left = l1
	root.Right = l2

	fmt.Println(isValidBST(root))

}

func isValidBST(root *TreeNode) bool {
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

// true 表示继续 false 表示不用继续
func inorderTrabversal_98(root *TreeNode, res *int, roots *[]int) bool {

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if !inorderTrabversal_98(root.Left, res, roots) {
		return false
	}

	*roots = append(*roots, root.Val)

	if root.Left != nil && root.Right != nil {
		if root.Left.Val >= root.Right.Val {
			*res = 1
			return false
		}

		for i := 0; i < len(*roots); i++ {
			if (*roots)[i] <= root.Left.Val || (*roots)[i] >= root.Right.Val {
				*res = 1
				return false
			}
		}

	}

	if root.Left == nil {
		for i := 0; i < len(*roots); i++ {
			if (*roots)[i] >= root.Right.Val {
				*res = 1

				return false
			}
		}
	}

	if root.Right == nil {
		for i := 0; i < len(*roots); i++ {
			if (*roots)[i] <= root.Left.Val {
				*res = 1
				return false
			}

		}

	}

	if !inorderTrabversal_98(root.Right, res, roots) {
		return false
	}
	return true
}
