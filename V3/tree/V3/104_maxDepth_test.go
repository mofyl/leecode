/*
 * @Author: lirui38
 * @Date: 2021-12-25 14:38:20
 * @LastEditTime: 2021-12-25 14:40:57
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\tree\V3\104_maxDepth_test.go
 */

package V3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxDepth(t *testing.T) {

	l3 := &TreeNode{Val: 3}
	l9 := &TreeNode{Val: 9}
	l20 := &TreeNode{Val: 20}
	l15 := &TreeNode{Val: 15}
	l7 := &TreeNode{Val: 7}

	l20.Left = l15
	l20.Right = l7

	l3.Left = l9
	l3.Right = l20

	res := maxDepth(l3)

	fmt.Println(res)
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		return left + right + 1
	}

	return tools.Max(left, right) + 1

}
