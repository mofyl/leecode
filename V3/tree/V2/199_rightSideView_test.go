package V2

import (
	"fmt"
	"testing"
)

func TestRightSideView(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}

	l1.Left = l2
	//l1.Right = l3

	l2.Right = l5

	l3.Right = l4

	res := rightSideView(l1)
	fmt.Println(res)
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	// dfs
	//rightSideViewHelper(root, 0, &res)

	// bfs
	que := make([]*TreeNode, 0)
	que = append(que, root)

	for len(que) > 0 {

		qLen := len(que)

		for i := 0; i < qLen; i++ {

			v := que[0]
			que = que[1:]

			if v.Left != nil {
				que = append(que, v.Left)
			}

			if v.Right != nil {
				que = append(que, v.Right)
			}
			// 这里表示当前层的最后一个节点
			if i == qLen-1 {
				res = append(res, v.Val)
			}

		}

	}

	return res
}

func rightSideViewHelper(root *TreeNode, depth int, res *[]int) {

	if root == nil {
		return
	}

	if depth == len(*res) {
		*res = append(*res, root.Val)
	}

	rightSideViewHelper(root.Right, depth+1, res)
	rightSideViewHelper(root.Left, depth+1, res)

}
