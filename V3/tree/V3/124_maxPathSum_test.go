/*
 * @Author: lirui38
 * @Date: 2021-12-20 14:21:50
 * @LastEditTime: 2021-12-20 14:51:34
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/124_maxPathSum_test.go
 */

package V3

import (
	"leecode/tools"
	"testing"
)

func TestMaxPathSum(t *testing.T) {

	// l10 := TreeNode{Val: -10}
	// l9 := TreeNode{Val: 9}
	// l20 := TreeNode{Val: 20}
	// l15 := TreeNode{Val: 15}
	// l7 := TreeNode{Val: 7}

}

func maxPathSum(root *TreeNode) int {

	if root == nil {
		return 0
	}
	maxValue := 0
	res := maxPathSumHelper(root, &maxValue)

	return tools.Max(maxValue, res)
}

// 返回当前 节点 到子树 中 节点和的最大值
/*

	3个节点 这里一共有3种情况
		a
	b		c

	1.最大值为 a+b+c
	2.最大值为 a + b + a的父节点
	3.最大值为 a + c + a的父节点

*/
func maxPathSumHelper(root *TreeNode, maxValue *int) int {

	if root == nil {
		return 0
	}

	left := tools.Max(0, maxPathSumHelper(root.Left, maxValue))
	right := tools.Max(0, maxPathSumHelper(root.Right, maxValue))

	// 这里是对 情况1 的处理
	*maxValue = tools.Max(*maxValue, root.Val+left+right)

	return root.Val + tools.Max(left, right)
}
