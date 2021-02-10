package V2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestRob(t *testing.T) {

	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//l31 := &TreeNode{Val: 3}
	//l32 := &TreeNode{Val: 3}
	//
	//l3.Left = l2
	//l3.Right = l31
	//
	//l2.Right = l32
	//
	//l31.Right = l1
	//
	//res := rob(l3)

	l1 := &TreeNode{Val: 1}
	l11 := &TreeNode{Val: 1}
	l3 := &TreeNode{Val: 3}
	l31 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}

	l3.Left = l4
	l3.Right = l5

	l4.Left = l1
	l4.Right = l31

	l5.Right = l11

	res := rob(l3)

	fmt.Println(res)
}

func rob(root *TreeNode) int {

	if root == nil {
		return 0
	}

	que := make([]*TreeNode, 0)
	odd := 0
	even := 0

	isOdd := true

	que = append(que, root)

	for len(que) > 0 {

		qLen := len(que)

		for i := 0; i < qLen; i++ {

			v := que[0]
			que = que[1:]

			if isOdd {
				odd += v.Val
			} else {
				even += v.Val
			}

			if v.Left != nil {
				que = append(que, v.Left)
			}
			if v.Right != nil {
				que = append(que, v.Right)
			}

		}
		isOdd = !isOdd
	}
	return tools.Max(odd, even)
}
