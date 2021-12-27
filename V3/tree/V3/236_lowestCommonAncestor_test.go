/*
 * @Author: lirui38
 * @Date: 2021-12-20 19:55:20
 * @LastEditTime: 2021-12-20 20:07:15
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/236_lowestCommonAncestor_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor_236(t *testing.T) {

	l3 := &TreeNode{Val: 3}
	l5 := &TreeNode{Val: 5}
	l1 := &TreeNode{Val: 1}
	l6 := &TreeNode{Val: 6}
	l2 := &TreeNode{Val: 2}
	l0 := &TreeNode{Val: 0}
	l8 := &TreeNode{Val: 8}
	l7 := &TreeNode{Val: 7}
	l4 := &TreeNode{Val: 4}

	l2.Left = l7
	l2.Right = l4

	l5.Left = l6
	l5.Right = l2

	l1.Left = l0
	l1.Right = l8

	l3.Left = l5
	l3.Right = l1

	p := l5
	// q := l1

	q := l4

	res := lowestCommonAncestor_236(l3, p, q)

	fmt.Println(res.Val)

}

func lowestCommonAncestor_236(root, p, q *TreeNode) *TreeNode {

	if root == nil || p == nil || q == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor_236(root.Left, p, q)
	right := lowestCommonAncestor_236(root.Right, p, q)

	// 若两边都找到了 那么直接就是 root
	// 否则 返回 不为nil 那一边

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	} else {
		return left
	}

}
