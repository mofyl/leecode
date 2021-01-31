package V1

import (
	"container/list"
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestLevelOrder(t *testing.T) {

	root := &tree.TreeNode{Val: 3}
	l1 := &tree.TreeNode{Val: 9}
	l2 := &tree.TreeNode{Val: 20}
	l3 := &tree.TreeNode{Val: 15}
	l4 := &tree.TreeNode{Val: 7}

	root.Left = l1
	root.Right = l2

	l2.Left = l3
	l2.Right = l4

	res := levelOrder(root)

	fmt.Println(res)
	fmt.Println(len(res))

}

func levelOrder(root *tree.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	stack := list.New()

	stack.PushBack(root)
	tmp := make([]int, 1)
	tmp[0] = root.Val
	res = append(res, tmp)
	for stack.Len() > 0 {
		length := stack.Len()

		tmp := make([]int, 0)
		for i := 0; i < length; i++ {

			e := stack.Back()
			stack.Remove(e)

			node := e.Value.(*tree.TreeNode)

			if node.Left != nil {
				stack.PushFront(node.Left)
				tmp = append(tmp, node.Left.Val)
			}
			if node.Right != nil {
				stack.PushFront(node.Right)
				tmp = append(tmp, node.Right.Val)
			}

		}

		if len(tmp) > 0 {
			res = append(res, tmp)
		}

	}

	return res

}
