package V2

import (
	"fmt"
	"testing"
)

func TestCountNodes(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l6 := &TreeNode{Val: 6}

	l1.Left = l2
	l1.Right = l3

	l2.Left = l4
	l2.Right = l5

	l3.Left = l6

	res := countNodes(l1)
	fmt.Println(res)
}

//
//func countNodes(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	que := make([]*TreeNode, 0)
//	res := 0
//
//	que = append(que, root)
//
//	for len(que) > 0 {
//
//		qLen := len(que)
//		res += qLen
//		for i := 0; i < qLen; i++ {
//
//			v := que[0]
//			que = que[1:]
//
//			if v.Left != nil {
//				que = append(que, v.Left)
//				//res++
//			}
//			if v.Right != nil {
//				que = append(que, v.Right)
//				//res++
//			}
//
//		}
//
//	}
//	return res
//}

func countNodes(root *TreeNode) int {

	h := getHeight(root)

	if h == 0 {
		return h
	}

	// 这里表示右子树的高度 是 h-1 那么说明 左子树一定满节点
	// h是当前树的高度， h-1则是子树的最大高度
	if getHeight(root.Right) == h-1 {
		// 这里 计算公式的时候 本来是要-1的，但是我们加上了根节点
		return (1 << (h - 1)) + countNodes(root.Right)
	} else {
		// 这里不等于 说明 右子树 一定满节点
		return (1 << (h - 2)) + countNodes(root.Left)
	}
}

func getHeight(root *TreeNode) int {

	h := 0
	p := root

	for p != nil {
		p = p.Left
		h++
	}
	return h
}
