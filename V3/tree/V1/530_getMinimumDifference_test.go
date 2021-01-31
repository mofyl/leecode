package V1

import (
	"fmt"
	"leecode/V3/tree"
	"math"
	"testing"
)

func TestGetMinimumDifference(t *testing.T) {
	//l1 := &TreeNode{Val: 236}
	//l2 := &TreeNode{Val: 104}
	//l3 := &TreeNode{Val: 701}
	//l4 := &TreeNode{Val: 227}
	//l5 := &TreeNode{Val: 911}
	//
	//l1.Left = l2
	//l1.Right = l3
	//
	//l2.Right = l4
	//
	//l3.Right = l5

	l1 := &tree.TreeNode{Val: 1564}
	l2 := &tree.TreeNode{Val: 1434}
	l3 := &tree.TreeNode{Val: 1}
	l4 := &tree.TreeNode{Val: 3048}
	l5 := &tree.TreeNode{Val: 3184}

	l1.Left = l2
	l1.Right = l4

	l2.Left = l3

	l4.Right = l5

	res := getMinimumDifference(l1)

	fmt.Println(res)
}

func getMinimumDifference(root *tree.TreeNode) int {

	if root == nil {
		return 0
	}
	minRes := math.MaxInt64
	prev := -1
	getMinimumDifference_loop(root, &minRes, &prev)

	return minRes
}

func getMinimumDifference_loop(root *tree.TreeNode, minRes *int, prev *int) {

	if root == nil {
		return
	}

	getMinimumDifference_loop(root.Left, minRes, prev)

	if prev != nil {
		if *prev != -1 {
			v := root.Val - *prev
			if *minRes > v {
				*minRes = v
			}
		}

		*prev = root.Val
	}

	getMinimumDifference_loop(root.Right, minRes, prev)
}
