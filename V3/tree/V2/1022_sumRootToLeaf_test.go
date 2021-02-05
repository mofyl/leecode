package V2

import (
	"fmt"
	"testing"
)

func TestSumRootToLeaf(t *testing.T) {
	l1 := &TreeNode{Val: 1}
	l11 := &TreeNode{Val: 1}
	l12 := &TreeNode{Val: 1}
	l13 := &TreeNode{Val: 1}
	l0 := &TreeNode{Val: 0}
	l01 := &TreeNode{Val: 0}
	l02 := &TreeNode{Val: 0}

	l1.Left = l0
	l1.Right = l11

	l0.Left = l01
	l0.Right = l12

	l11.Left = l02
	l11.Right = l13

	fmt.Println(sumRootToLeaf(l1))
}

func sumRootToLeaf(root *TreeNode) int {

	if root == nil {
		return 0
	}
	sum := 0
	sumRootToLeafLoop(root, 0, &sum)
	return sum
}

func sumRootToLeafLoop(root *TreeNode, cur int, sum *int) {
	if root == nil {
		return
	}

	cur = cur<<1 | root.Val

	if root.Left == nil && root.Right == nil {
		*sum += cur
		return
	}

	sumRootToLeafLoop(root.Left, cur, sum)
	sumRootToLeafLoop(root.Right, cur, sum)

}
