package V2

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l4 := &TreeNode{Val: 4}
	l41 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l51 := &TreeNode{Val: 5}
	l7 := &TreeNode{Val: 7}
	l8 := &TreeNode{Val: 8}
	l11 := &TreeNode{Val: 11}
	l13 := &TreeNode{Val: 13}

	l5.Left = l4
	l5.Right = l8

	l4.Left = l11

	l11.Left = l7
	l11.Right = l2

	l8.Left = l13
	l8.Right = l41

	l41.Left = l51
	l41.Right = l1

	res := pathSum(l5, 22)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {

	if root == nil {
		return nil
	}

	res := make([][]int, 0)

	pathSumHelper(root, targetSum, []int{}, &res)
	return res
}

func pathSumHelper(root *TreeNode, targetSum int, cur []int, res *[][]int) {

	if root == nil {
		return
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		cur = append(cur, root.Val)
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	cur = append(cur, root.Val)
	pathSumHelper(root.Left, targetSum-root.Val, cur, res)
	cur = cur[:len(cur)-1]

	cur = append(cur, root.Val)
	pathSumHelper(root.Right, targetSum-root.Val, cur, res)
	cur = cur[:len(cur)-1]

}
