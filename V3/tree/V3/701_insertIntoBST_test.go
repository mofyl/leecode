/*
 * @Author: lirui38
 * @Date: 2021-12-23 10:25:04
 * @LastEditTime: 2021-12-23 14:26:51
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/701_insertIntoBST_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestInsertIntoBST(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l7 := &TreeNode{Val: 7}

	l2.Left = l1
	l2.Right = l3

	l4.Left = l2
	l4.Right = l7

	root := insertIntoBSTV2(l4, 5)

	res := inorderTraversal(root)

	fmt.Println(res)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root

}

func insertIntoBSTV2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	node := root

	for node != nil {

		if node.Val > val {

			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				break
			} else {
				node = node.Left
			}

		} else {

			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				break
			} else {
				node = node.Right
			}

		}

	}

	return root

}
