package V1

import (
	"container/list"
	"fmt"
	"testing"
)

type NNode struct {
	Val      int
	Children []*NNode
}

func TestPreorder(t *testing.T) {
	l1 := &NNode{Val: 1}
	l2 := &NNode{Val: 2}
	l3 := &NNode{Val: 3}
	l4 := &NNode{Val: 4}
	l5 := &NNode{Val: 5}
	l6 := &NNode{Val: 6}

	l3.Children = append(l3.Children, l5)
	l3.Children = append(l3.Children, l6)

	l1.Children = append(l1.Children, l3)
	l1.Children = append(l1.Children, l2)
	l1.Children = append(l1.Children, l4)

	res := preorder_v2(l1)

	fmt.Println(res)
}

func preorder_v2(root *NNode) []int {

	if root == nil {
		return []int{}
	}

	res := make([]int, 0)

	stack := list.New()

	stack.PushBack(root)

	for stack.Len() > 0 {

		ele := stack.Front()
		e := ele.Value.(*NNode)
		res = append(res, e.Val)

		nEle := ele.Next()
		for i := 0; i < len(e.Children); i++ {
			if nEle != nil {
				stack.InsertBefore(e.Children[i], nEle)
			} else {
				stack.PushBack(e.Children[i])
			}
		}
		stack.Remove(ele)
	}
	return res
}

func preorder(root *NNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)

	preorder_loop(root, &res)

	return res
}

func preorder_loop(root *NNode, res *[]int) {

	if root == nil {
		return
	}

	*res = append(*res, root.Val)

	for i := 0; i < len(root.Children); i++ {

		if root.Children[i] != nil {
			preorder_loop(root.Children[i], res)
		}

	}

}
