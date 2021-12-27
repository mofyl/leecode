/*
 * @Author: lirui38
 * @Date: 2021-12-20 19:23:53
 * @LastEditTime: 2021-12-20 19:51:02
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/235_lowestCommonAncestor_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {

	l6 := &TreeNode{Val: 6}
	l2 := &TreeNode{Val: 2}
	l8 := &TreeNode{Val: 8}
	l0 := &TreeNode{Val: 0}
	l4 := &TreeNode{Val: 4}
	l7 := &TreeNode{Val: 7}
	l9 := &TreeNode{Val: 9}
	l3 := &TreeNode{Val: 3}
	l5 := &TreeNode{Val: 5}

	l4.Left = l3
	l4.Right = l5

	l2.Left = l0
	l2.Right = l4

	l8.Left = l7
	l8.Right = l9

	l6.Left = l2
	l6.Right = l8

	p := l2
	q := l8

	res := lowestCommonAncestor(l6, p, q)

	fmt.Println(res)

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || p == nil || q == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}

}
