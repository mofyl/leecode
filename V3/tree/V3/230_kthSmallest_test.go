/*
 * @Author: lirui38
 * @Date: 2021-12-20 16:41:51
 * @LastEditTime: 2021-12-20 19:20:58
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/230_kthSmallest_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {

	// l3 := &TreeNode{Val: 3}
	// l1 := &TreeNode{Val: 1}
	// l4 := &TreeNode{Val: 4}
	// l2 := &TreeNode{Val: 2}

	// l1.Right = l2

	// l3.Left = l1
	// l3.Right = l4

	// res := kthSmallest(l3, 3)

	// fmt.Println(res)

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l6 := &TreeNode{Val: 6}

	l2.Left = l1

	l3.Left = l2
	l3.Right = l4

	l5.Left = l3
	l5.Right = l6

	// res := kthSmallest(l5, 3)
	res := kthSmallestV2(l5, 3)

	fmt.Println(res)
}

func kthSmallest(root *TreeNode, k int) int {
	num := inorderTraversal(root)

	return num[k-1]
}

func kthSmallestV2(root *TreeNode, k int) int {

	prev := 0
	kthSmallestV2Helper(root, &k, &prev)

	return prev
}

func kthSmallestV2Helper(root *TreeNode, k *int, prev *int) {

	if root == nil {
		return
	}

	if *k == 0 {
		return
	}

	kthSmallestV2Helper(root.Left, k, prev)

	if *k == 0 {
		return
	}

	*prev = root.Val
	*k--

	kthSmallestV2Helper(root.Right, k, prev)
	if *k == 0 {
		return
	}
}
