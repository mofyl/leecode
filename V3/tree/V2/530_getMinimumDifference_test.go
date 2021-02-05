package V2

import (
	"fmt"
	"leecode/tools"
	"math"
	"testing"
)

func TestGetMinimumDifference(t *testing.T) {
	//l1 := &TreeNode{Val: 1}
	//l3 := &TreeNode{Val: 3}
	//l2 := &TreeNode{Val: 2}
	//
	//l1.Right = l3
	//l3.Left = l2

	l236 := &TreeNode{Val: 236}
	l104 := &TreeNode{Val: 104}
	l701 := &TreeNode{Val: 701}
	l227 := &TreeNode{Val: 227}
	l911 := &TreeNode{Val: 911}

	l236.Left = l104
	l236.Right = l701

	l104.Right = l227

	l701.Right = l911

	fmt.Println(getMinimumDifference(l236))
}

func getMinimumDifference(root *TreeNode) int {
	minV := math.MaxInt64
	prev := -1
	getMinimumDifferenceLoop(root, &minV, &prev)
	return minV
}

func getMinimumDifferenceLoop(root *TreeNode, minV *int, prev *int) {

	if root == nil {
		return
	}
	getMinimumDifferenceLoop(root.Left, minV, prev)
	if *prev != -1 {
		curV := tools.Abs(root.Val - *prev)
		*minV = tools.Min(*minV, curV)
	}

	*prev = root.Val
	getMinimumDifferenceLoop(root.Right, minV, prev)
}
