/*
 * @Author: lirui38
 * @Date: 2021-12-23 10:19:40
 * @LastEditTime: 2021-12-23 10:23:09
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/700_searchBST_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestSearchBST(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l7 := &TreeNode{Val: 7}

	l2.Left = l1
	l2.Right = l3

	l4.Left = l2
	l4.Right = l7

	root := searchBST(l4, 5)
	res := inorderTraversal(root)

	fmt.Println(res)
}

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}
