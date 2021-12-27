/*
 * @Author: lirui38
 * @Date: 2021-12-17 14:46:39
 * @LastEditTime: 2021-12-17 14:55:26
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/101_isSymmetric_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {

	// l1 := &TreeNode{Val: 1}
	// l2 := &TreeNode{Val: 2}
	// l21 := &TreeNode{Val: 2}
	// l3 := &TreeNode{Val: 3}
	// l31 := &TreeNode{Val: 3}
	// l4 := &TreeNode{Val: 4}
	// l41 := &TreeNode{Val: 4}

	// l1.Left = l2
	// l1.Right = l21

	// l2.Left = l3
	// l2.Right = l4

	// l21.Left = l41
	// l21.Right = l31

	// fmt.Println(isSymmetric(l1))

	// l1 := &TreeNode{Val: 1}
	// l2 := &TreeNode{Val: 2}
	// l21 := &TreeNode{Val: 2}
	// l3 := &TreeNode{Val: 3}
	// l31 := &TreeNode{Val: 3}

	// l1.Left = l2
	// l1.Right = l21

	// l2.Right = l3
	// l21.Right = l31

	// fmt.Println(isSymmetric(l1))

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}

	l1.Left = l2
	l1.Right = l3

	fmt.Println(isSymmetric(l1))
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if left == nil && right != nil {
		return false
	}

	if right == nil && left != nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}
