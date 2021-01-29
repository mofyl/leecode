package tree

import (
	"container/list"
	"fmt"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {

	l3 := &TreeNode{Val: 3}
	l9 := &TreeNode{Val: 9}
	l20 := &TreeNode{Val: 20}
	l15 := &TreeNode{Val: 15}
	l7 := &TreeNode{Val: 7}

	l3.Left = l9
	l3.Right = l20

	l20.Left = l15
	l20.Right = l7

	res := averageOfLevels(l3)

	fmt.Println(res)
}

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}

	stack := list.New()

	stack.PushBack(root)

	stackLen := 0
	sum := 0
	for stack.Len() > 0 {
		sum = 0
		stackLen = stack.Len()
		for i := 0; i < stackLen; i++ {
			e := stack.Front()
			stack.Remove(e)
			node := e.Value.(*TreeNode)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			sum += node.Val
		}

		res = append(res, float64(sum)/float64(stackLen))
	}

	return res
}
