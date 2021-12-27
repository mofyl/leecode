/*
 * @Author: lirui38
 * @Date: 2021-12-25 14:44:15
 * @LastEditTime: 2021-12-25 15:50:01
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\tree\V3\572_isSubtree_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestIsSubtree(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}

	l4.Left = l1
	l4.Right = l2

	l3.Left = l4
	l3.Right = l5

	l41 := &TreeNode{Val: 4}
	l11 := &TreeNode{Val: 1}
	l21 := &TreeNode{Val: 2}

	l41.Left = l11
	l41.Right = l21

	res := isSubtree(l3, l41)

	fmt.Println(res)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	return isSameTreeV2(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func isSameTreeV2(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	return root.Val == subRoot.Val && isSameTreeV2(root.Left, subRoot.Left) && isSameTreeV2(root.Right, subRoot.Right)

}
