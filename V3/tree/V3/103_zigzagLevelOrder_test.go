/*
 * @Author: lirui38
 * @Date: 2021-12-17 15:39:27
 * @LastEditTime: 2021-12-17 16:32:04
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/103_zigzagLevelOrder_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {

	// l3 := &TreeNode{Val: 3}
	// l9 := &TreeNode{Val: 9}
	// l20 := &TreeNode{Val: 20}
	// l15 := &TreeNode{Val: 15}
	// l7 := &TreeNode{Val: 7}

	// l20.Left = l15
	// l20.Right = l7

	// l3.Left = l9
	// l3.Right = l20

	// res := zigzagLevelOrder(l3)

	// fmt.Println(res)

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}

	l2.Left = l4
	l3.Right = l5

	l1.Left = l2
	l1.Right = l3

	res := zigzagLevelOrder(l1)

	fmt.Println(res)

}

func zigzagLevelOrder(root *TreeNode) [][]int {

	// 本质上和 层序一样  只不过 在偶数层的时候将 该层的结果翻转一次

	if root == nil {
		return nil
	}

	res := make([][]int, 0)

	arr := make([]*TreeNode, 0)
	arr = append(arr, root)
	level := 0
	for len(arr) > 0 {

		level++
		length := len(arr)
		tmp := make([]int, 0, length)
		for i := 0; i < length; i++ {

			node := arr[0]
			arr = arr[1:]

			tmp = append(tmp, node.Val)

			if node.Left != nil {
				arr = append(arr, node.Left)
			}

			if node.Right != nil {
				arr = append(arr, node.Right)
			}
		}

		if level&1 == 0 {
			// 表示是偶数
			n := len(tmp)
			count := n / 2
			for i := 0; i < count; i++ {
				tmp[i], tmp[n-i-1] = tmp[n-i-1], tmp[i]
			}
		}

		if len(tmp) > 0 {
			res = append(res, tmp)
		}

	}

	return res

}
