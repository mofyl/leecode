package V1

import (
	"container/list"
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	//
	//root := &TreeNode{Val: 3}
	//l1 := &TreeNode{Val: 9}
	//l2 := &TreeNode{Val: 20}
	//l3 := &TreeNode{Val: 15}
	//l4 := &TreeNode{Val: 7}
	//
	//root.Left = l1
	//root.Right = l2
	//
	//l2.Left = l3
	//l2.Right = l4
	//
	//res := zigzagLevelOrder(root)
	//
	//fmt.Println(res)

	root := &tree.TreeNode{Val: 1}
	l1 := &tree.TreeNode{Val: 2}
	l2 := &tree.TreeNode{Val: 3}
	l3 := &tree.TreeNode{Val: 4}
	l4 := &tree.TreeNode{Val: 5}

	root.Left = l1
	root.Right = l2

	l1.Left = l3
	l2.Right = l4

	res := zigzagLevelOrder(root)

	fmt.Println(res)

}

/*

flag 状态:
	false表示 从右向左走 先加right 再加left。访问的时候要用Back 加的时候要用 PushFront
	true 表示 从左向右走 先加left 再加right。访问的时候要用Front 加的时候要用 PushBuck

*/
func zigzagLevelOrder(root *tree.TreeNode) [][]int {

	res := make([][]int, 0)

	if root == nil {
		return res
	}

	stack := list.New()

	stack.PushBack(root)
	res = append(res, []int{root.Val})
	flag := false

	for stack.Len() > 0 {

		length := stack.Len()
		tmp := make([]int, 0, length)
		for i := 0; i < length; i++ {
			var e *list.Element

			if !flag {
				e = stack.Back()
			} else {
				e = stack.Front()
			}

			node := e.Value.(*tree.TreeNode)
			stack.Remove(e)

			// 表示要从 右向左加
			if !flag {
				if node.Right != nil {
					stack.PushFront(node.Right)
					tmp = append(tmp, node.Right.Val)
				}
				if node.Left != nil {
					stack.PushFront(node.Left)
					tmp = append(tmp, node.Left.Val)
				}

			} else {
				if node.Left != nil {
					stack.PushBack(node.Left)
					tmp = append(tmp, node.Left.Val)
				}

				if node.Right != nil {
					stack.PushBack(node.Right)
					tmp = append(tmp, node.Right.Val)
				}
			}

		}

		if len(tmp) > 0 {
			res = append(res, tmp)
			flag = !flag
		}

	}

	return res

}
