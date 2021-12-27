/*
 * @Author: lirui38
 * @Date: 2021-12-17 14:35:53
 * @LastEditTime: 2021-12-17 14:41:13
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/100_isSameTree_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}

	l1.Left = l2
	l1.Right = l3

	l11 := &TreeNode{Val: 1}
	l21 := &TreeNode{Val: 1}
	l31 := &TreeNode{Val: 2}

	l11.Left = l21
	l11.Right = l31

	fmt.Println(isSameTree(l1, l11))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if q == nil && p != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
