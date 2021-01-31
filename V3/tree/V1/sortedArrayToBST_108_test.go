package V1

import (
	"fmt"
	"leecode/V3/tree"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}

	res := sortedArrayToBST(nums)

	resNum := inorderTraversal(res)
	fmt.Println(resNum)
}

func sortedArrayToBST(nums []int) *tree.TreeNode {

	if len(nums) < 1 {
		return nil
	}

	return sortedArrayToBSTLoop(nums, 0, len(nums)-1)
}

// 这里要求是一颗 二叉搜索树 所以要从 nums 的中间开始
func sortedArrayToBSTLoop(nums []int, start int, end int) *tree.TreeNode {

	if start < 0 || end >= len(nums) {
		return nil
	}

	if start > end {
		return nil
	}

	mid := (start + end) >> 1

	root := &tree.TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBSTLoop(nums, start, mid-1)
	root.Right = sortedArrayToBSTLoop(nums, mid+1, end)

	return root
}
