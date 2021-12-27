/*
 * @Author: lirui38
 * @Date: 2021-12-20 20:13:19
 * @LastEditTime: 2021-12-20 20:45:25
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/450_deleteNode_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {

	l5 := &TreeNode{Val: 5}
	l3 := &TreeNode{Val: 3}
	l6 := &TreeNode{Val: 6}
	l2 := &TreeNode{Val: 2}
	l4 := &TreeNode{Val: 4}
	l7 := &TreeNode{Val: 7}

	l3.Left = l2
	l3.Right = l4

	l6.Right = l7

	l5.Left = l3
	l5.Right = l6

	// root := deleteNode(l5, 3)

	// res := inorderTraversal(root)

	// fmt.Println(res)

	root := deleteNode(l5, 3)

	res := inorderTraversal(root)

	fmt.Println(res)

}

func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {

		if root.Left == nil && root.Right == nil {
			// 如果该节点 没有 子节点 直接删除
			return nil
		}

		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		// 两个子树都不为空，则到 左边找一个最大的，或者 到右边找一个最小的 节点返回即可

		// 到左边找一个最大的
		maxLeftNode := findMaxNode(root.Left)
		// 交换两个节点的值
		root.Val = maxLeftNode.Val
		// 到左边把这个节点删除掉
		root.Left = deleteNode(root.Left, maxLeftNode.Val)

		return root
	}

	return root

}

// 到左子树中 找一个最大的 就是一直往 右走
func findMaxNode(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Right == nil {
		return root
	}

	return findMaxNode(root.Right)

}
