/*
 * @Author: lirui38
 * @Date: 2021-12-17 15:04:37
 * @LastEditTime: 2021-12-17 15:37:54
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/102_levelOrder_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {

	l3 := &TreeNode{Val: 3}
	l9 := &TreeNode{Val: 9}
	l20 := &TreeNode{Val: 20}
	l15 := &TreeNode{Val: 15}
	l7 := &TreeNode{Val: 7}

	l20.Left = l15
	l20.Right = l7

	l3.Left = l9
	l3.Right = l20

	res := levelOrder(l3)

	fmt.Println(res)

}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	res := make([][]int, 0)

	arr := make([]*TreeNode, 0)
	arr = append(arr, root)

	for len(arr) > 0 {

		length := len(arr)

		tmp := make([]int, 0)

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

		if len(tmp) > 0 {
			res = append(res, tmp)
		}

	}

	return res

}
