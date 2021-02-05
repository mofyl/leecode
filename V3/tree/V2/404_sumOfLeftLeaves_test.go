package V2

import (
	"fmt"
	"testing"
)

func TestSumOfLeftLeaves(t *testing.T) {

	l3 := &TreeNode{Val: 3}
	l9 := &TreeNode{Val: 9}
	l20 := &TreeNode{Val: 20}
	l15 := &TreeNode{Val: 15}
	l7 := &TreeNode{Val: 7}

	l3.Left = l9
	l3.Right = l20

	l20.Left = l15
	l20.Right = l7

	res := sumOfLeftLeaves(l3)

	fmt.Println(res)
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0

	sumOfLeftLeavesLoop(root, &res)
	return res
}

func sumOfLeftLeavesLoop(root *TreeNode, res *int) {

	if root == nil {
		return
	}

	// 这里就证明他是一个左叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res += root.Left.Val
	}

	sumOfLeftLeavesLoop(root.Left, res)
	sumOfLeftLeavesLoop(root.Right, res)

}
