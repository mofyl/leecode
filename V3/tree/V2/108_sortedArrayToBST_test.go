package V2

import (
	"fmt"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {

	nums := []int{-10, -3, 0, 5, 9}

	root := sortedArrayToBST(nums)

	res := inorderTraversal(root)

	fmt.Println(res)
}

// 平衡二叉树 只能是从中间开始分 中间元素是根节点
func sortedArrayToBST(nums []int) *TreeNode {
	// if len(nums) < 1 {
	// 		return nil
	// 	}

	// 	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := len(nums) / 2

	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

func sortedArrayToBSTHelper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2

	left := sortedArrayToBSTHelper(nums, start, mid-1)
	right := sortedArrayToBSTHelper(nums, mid+1, end)

	root := &TreeNode{Val: nums[mid]}
	root.Left = left
	root.Right = right

	return root
}
