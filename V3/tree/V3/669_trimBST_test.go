/*
 * @Author: lirui38
 * @Date: 2021-12-23 10:07:03
 * @LastEditTime: 2021-12-23 10:18:09
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/669_trimBST_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestTrimBST(t *testing.T) {

	l0 := &TreeNode{Val: 0}
	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}

	l2.Left = l1

	l0.Right = l2

	l3.Left = l0
	l3.Right = l4

	root := trimBST(l3, 1, 3)

	res := inorderTraversal(root)

	fmt.Println(res)
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {

	if root == nil {
		return root
	}

	if root.Val < low {
		// 由于是二叉搜索树，若 Val < low 则表示要将 以 root 为根节点的左子树都删除
		// 这里的删除策略是，跳过 左子树,直接去遍历右子树，然后返回
		return trimBST(root.Right, low, high)
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
