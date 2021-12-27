/*
 * @Author: lirui38
 * @Date: 2021-12-25 14:19:56
 * @LastEditTime: 2021-12-25 14:38:04
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\tree\V3\111_minDepth_test.go
 */

package V3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMinDepth(t *testing.T) {

	// l3 := &TreeNode{Val: 3}
	// l9 := &TreeNode{Val: 9}
	// l20 := &TreeNode{Val: 20}
	// l15 := &TreeNode{Val: 15}
	// l7 := &TreeNode{Val: 7}

	// l20.Left = l15
	// l20.Right = l7

	// l3.Left = l9
	// l3.Right = l20

	// depth := minDepth(l3)

	// fmt.Println(depth)

	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l6 := &TreeNode{Val: 6}

	l5.Right = l6

	l4.Right = l5

	l3.Right = l4

	l2.Right = l3

	depth := minDepth(l2)

	fmt.Println(depth)
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	// 叶子节点表示没有子节点的节点
	// 如若一个节点有 左子树 || 右子树 那么该节点不能称之为叶子节点
	// 此时我们应该返回子树中 不为空的那一边的长度
	if root.Left == nil || root.Right == nil {
		return left + right + 1
	}

	return tools.Min(left, right) + 1

}
