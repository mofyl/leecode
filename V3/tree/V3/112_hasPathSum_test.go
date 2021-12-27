/*
 * @Author: lirui38
 * @Date: 2021-12-20 10:39:40
 * @LastEditTime: 2021-12-20 10:55:55
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/112_hasPathSum_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestHasPathSum(t *testing.T) {

	// l5 := &TreeNode{Val: 5}
	// l4 := &TreeNode{Val: 4}
	// l8 := &TreeNode{Val: 8}
	// l11 := &TreeNode{Val: 11}
	// l13 := &TreeNode{Val: 13}
	// l7 := &TreeNode{Val: 7}
	// l2 := &TreeNode{Val: 2}
	// l1 := &TreeNode{Val: 1}

	// l5.Left = l4
	// l5.Right = l8

	// l4.Left = l11

	// l8.Left = l13
	// l8.Right = l4

	// l11.Left = l7
	// l11.Right = l2

	// l4.Right = l1

	// fmt.Println(hasPathSum(l5, 22))

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}

	l1.Left = l2
	l1.Right = l3

	fmt.Println(hasPathSum(l1, 5))

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	// 若 root == nil 那么不存在 根节点到 叶子节点
	if root == nil {
		return false
	}
	return hasPathSumHelper(root, targetSum, root.Val)
}

func hasPathSumHelper(root *TreeNode, targetSum int, curSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && targetSum == curSum {
		return true
	}

	if root.Left == nil && root.Right == nil && targetSum != curSum {
		return false
	}

	hasTarget := false
	if root.Left != nil {
		hasTarget = hasPathSumHelper(root.Left, targetSum, curSum+root.Left.Val)
	}

	if hasTarget {
		return hasTarget
	}

	if root.Right != nil {
		hasTarget = hasPathSumHelper(root.Right, targetSum, curSum+root.Right.Val)
	}

	if hasTarget {
		return hasTarget
	}

	return false

}
