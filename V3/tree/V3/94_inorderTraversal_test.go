/*
 * @Author: lirui38
 * @Date: 2021-12-16 17:13:39
 * @LastEditTime: 2021-12-16 17:24:17
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/94_inorderTraversal_test.go
 */

package V3

import (
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	t1 := &TreeNode{
		Val: 1,
	}
	t2 := &TreeNode{
		Val: 2,
	}
	t3 := &TreeNode{
		Val: 3,
	}

	t1.Right = t2
	t2.Left = t3

	res := inorderTraversal(t1)

	t.Logf("%+v", res)
}

func inorderTraversal(root *TreeNode) []int {

	var res []int

	inorderTraversalLoop(root, &res)
	return res
}

func inorderTraversalLoop(root *TreeNode, res *[]int) {

	if root == nil {
		return
	}

	inorderTraversalLoop(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalLoop(root.Right, res)
}
