package V1

import (
	"fmt"
	"leecode/V3/tree"
	"leecode/tools"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := &tree.TreeNode{Val: 3}
	l1 := &tree.TreeNode{Val: 9}
	l2 := &tree.TreeNode{Val: 20}
	l3 := &tree.TreeNode{Val: 15}
	l4 := &tree.TreeNode{Val: 7}

	root.Left = l1
	root.Right = l2

	l2.Left = l3
	l2.Right = l4

	res := maxDepth(root)

	fmt.Println(res)
}

func maxDepth(root *tree.TreeNode) int {

	res := levelOrder(root)
	return len(res)
}

// 计算 左边的深度 在去 计算右边的深度，最后取最大值就好了
// 这里是 递归 从下向上求的。 先是找到 没有节点的那一层 然后返回0，上面那一层 返回的就是 0+1
func maxDepth_v2(root *tree.TreeNode) int {

	if root == nil {
		return 0
	}

	return tools.Max(maxDepth_v2(root.Left), maxDepth_v2(root.Right)) + 1
}
