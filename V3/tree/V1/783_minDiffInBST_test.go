package V1

import (
	"fmt"
	"leecode/V3/tree"
	"leecode/tools"
	"math"
	"testing"
)

func TestMinDiffInBST(t *testing.T) {
	//
	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//l4 := &TreeNode{Val: 4}
	//l6 := &TreeNode{Val: 6}
	//
	//l4.Left = l2
	//l4.Right = l6
	//
	//l2.Left = l1
	//l2.Right = l3
	//
	//res := minDiffInBST(l4)
	//fmt.Println(res)

	l1 := &tree.TreeNode{Val: 1}
	l0 := &tree.TreeNode{Val: 0}
	l48 := &tree.TreeNode{Val: 48}
	l12 := &tree.TreeNode{Val: 12}
	l49 := &tree.TreeNode{Val: 49}

	l1.Left = l0
	l1.Right = l48

	l48.Left = l12
	l48.Right = l49

	res := minDiffInBST(l1)
	fmt.Println(res)
}

func minDiffInBST(root *tree.TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	min := math.MaxInt64
	prev := -1
	minDiffInBST_loop(root, &min, &prev)
	return min
}

func minDiffInBST_loop(root *tree.TreeNode, min *int, prev *int) {

	if root == nil {
		return
	}

	minDiffInBST_loop(root.Left, min, prev)

	if *prev != -1 {
		subV := tools.Abs(root.Val - *prev)
		if subV < *min {
			*min = subV
		}
	}

	*prev = root.Val

	minDiffInBST_loop(root.Right, min, prev)

}
