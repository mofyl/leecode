package main

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {

	c := make([]int, 0, 3)
	c = append(c, 2)
	c = append(c, 3)
	c = append(c, 4)

	v := make([]int, len(c))
	copy(v, c)

	fmt.Println(v)
}

func pathSum_113(root *TreeNode, sum int) [][]int {

	res := make([][]int, 0)

	pathSum1(root, sum, make([]int, 0), &res)

	return res
}

func pathSum1(root *TreeNode, sum int, ary []int, res *[][]int) {

	if root == nil {
		return
	}

	ary = append(ary, root.Val)

	if root.Left == nil && root.Right == nil && sum == root.Val {
		tmp := make([]int, len(ary))
		copy(tmp, ary)
		*res = append(*res, tmp)
		return
	}

	pathSum1(root.Left, sum-root.Val, ary, res)
	pathSum1(root.Right, sum-root.Val, ary, res)
}
