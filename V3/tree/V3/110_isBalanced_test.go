/*
 * @Author: lirui38
 * @Date: 2021-12-17 17:22:57
 * @LastEditTime: 2021-12-20 10:36:34
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/110_isBalanced_test.go
 */

package V3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestIsBalanced(t *testing.T) {

	// l3 := &TreeNode{Val: 3}
	// l9 := &TreeNode{Val: 9}
	// l20 := &TreeNode{Val: 20}
	// l15 := &TreeNode{Val: 15}
	// l7 := &TreeNode{Val: 7}

	// l20.Left = l15
	// l20.Right = l7

	// l3.Left = l9
	// l3.Right = l20

	// fmt.Println(isBalanced(l3))

	l1 := &TreeNode{Val: 1}
	l21 := &TreeNode{Val: 2}
	l22 := &TreeNode{Val: 2}
	l31 := &TreeNode{Val: 3}
	l32 := &TreeNode{Val: 3}
	l41 := &TreeNode{Val: 4}
	l42 := &TreeNode{Val: 4}

	l1.Left = l21
	l1.Right = l22

	l21.Left = l31
	l21.Right = l32

	l31.Left = l41
	l31.Right = l42

	fmt.Println(isBalanced(l1))
}

func isBalanced(root *TreeNode) bool {

	res := 0

	isBalancedHelper(root, &res)

	return res == 0
}

// *res == 1 表示该树 已经不是 平衡二叉树了 ，就该直接return了
func isBalancedHelper(root *TreeNode, res *int) int {

	if root == nil {
		return 1
	}

	if *res == 1 {
		return 0
	}

	left := isBalancedHelper(root.Left, res)
	right := isBalancedHelper(root.Right, res)

	if *res == 1 {
		return 0
	}

	if tools.Abs(int32(left-right)) > 1 {
		*res = 1
		return 0
	}

	return tools.Max(left, right) + 1
}
