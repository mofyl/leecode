package V2

import (
	"fmt"
	"math"
	"testing"
)

func TestLargestValues(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l5 := &TreeNode{Val: 5}
	l31 := &TreeNode{Val: 3}
	l9 := &TreeNode{Val: 9}

	l1.Left = l3
	l1.Right = l2

	l3.Left = l5
	l3.Right = l31

	l2.Right = l9

	res := largestValues(l1)
	fmt.Println(res)
}

// 层序遍历
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	que := make([]*TreeNode, 0)
	res := make([]int, 0)
	maxLevel := math.MinInt64

	que = append(que, root)

	for len(que) > 0 {

		qLen := len(que)
		maxLevel = math.MinInt64
		for i := 0; i < qLen; i++ {

			n := que[0]
			que = que[1:]
			if n.Val > maxLevel {
				maxLevel = n.Val
			}
			if n.Left != nil {
				que = append(que, n.Left)
			}
			if n.Right != nil {
				que = append(que, n.Right)
			}
		}

		res = append(res, maxLevel)

	}
	return res
}
