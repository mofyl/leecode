/*
 * @Author: lirui38
 * @Date: 2021-12-19 11:14:13
 * @LastEditTime: 2021-12-19 11:21:10
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\tree\V3\108_sortedArrayToBST_test.go
 */
package V3

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

func sortedArrayToBST(nums []int) *TreeNode {

	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)

}

func sortedArrayToBSTHelper(nums []int, start, end int) *TreeNode {

	if start > end {
		return nil
	}

	sum := start + end
	idx := sum / 2
	if sum&1 == 1 {
		idx++
	}

	root := &TreeNode{Val: nums[idx]}
	root.Left = sortedArrayToBSTHelper(nums, start, idx-1)
	root.Right = sortedArrayToBSTHelper(nums, idx+1, end)

	return root

}
